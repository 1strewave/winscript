package runtime

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
	"unsafe"

	"github.com/1strewave/winscript/internal/models"
	"github.com/atotto/clipboard"
	"github.com/lxn/win"
	"github.com/micmonay/keybd_event"
)

type CommandHandler func(cmd models.Command) error

type CommandRegistry map[string]CommandHandler

func NewCommandRegistry() CommandRegistry {
	return CommandRegistry{
		"open":       handleOpen,
		"type":       handleType,
		"wait":       handleWait,
		"log":        handleLog,
		"focus":      handleFocus,
		"press":      handlePress,
		"hotkey":     handleHotkey,
		"move_mouse": handleMoveMouse,
		"click":      handleClick,
	}
}

func Execute(commands []models.Command) error {
	registry := NewCommandRegistry()
	ctx := context.Background()

	for _, cmd := range commands {
		if handler, exists := registry[cmd.Name]; exists {
			if err := handler(cmd); err != nil {
				return fmt.Errorf("error executing command '%s' (line %d): %w",
					cmd.OrigText, cmd.LineNum, err)
			}
		} else {
			return fmt.Errorf("unknown command '%s' at line %d", cmd.Name, cmd.LineNum)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	return nil
}

func handleOpen(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'open' requires at least one argument")
	}

	args := []string{"/C", "start"}
	args = append(args, cmd.Args...)

	command := exec.Command("cmd", args...)
	return command.Start()
}

func handleType(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'type' requires one argument")
	}

	return typeViaClipboard(cmd.Args[0])
}

func handleWait(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'wait' requires one argument")
	}

	d, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid duration '%s': %w", cmd.Args[0], err)
	}

	time.Sleep(d)
	return nil
}

func handleLog(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'log' requires one argument")
	}

	fmt.Println(cmd.Args[0])
	return nil
}

func handleFocus(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'focus' requires one argument")
	}

	return focusWindow(cmd.Args[0])
}

func handlePress(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'press' requires one argument (key)")
	}

	key := cmd.Args[0]
	return pressKey(key)
}

func handleHotkey(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'hotkey' requires one argument (combo)")
	}

	combo := cmd.Args[0]
	return pressHotkey(combo)
}

func handleMoveMouse(cmd models.Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("'move_mouse' requires two arguments (X Y coordinates)")
	}

	x, err := strconv.Atoi(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("invalid X coordinate '%s': %w", cmd.Args[0], err)
	}

	y, err := strconv.Atoi(cmd.Args[1])
	if err != nil {
		return fmt.Errorf("invalid Y coordinate '%s': %w", cmd.Args[1], err)
	}

	return moveMouse(x, y)
}

func handleClick(cmd models.Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("'click' requires one argument (button)")
	}

	button := strings.ToLower(cmd.Args[0])
	return mouseClick(button)
}

func UTF16PtrFromString(s string) *uint16 {
	encoded := utf16.Encode([]rune(s))
	encoded = append(encoded, 0)
	return &encoded[0]
}

func focusWindow(title string) error {
	hwnd := win.FindWindow(nil, UTF16PtrFromString(title))
	if hwnd == 0 {
		return fmt.Errorf("window with title '%s' not found", title)
	}

	win.ShowWindow(hwnd, win.SW_RESTORE)
	win.SetForegroundWindow(hwnd)

	time.Sleep(200 * time.Millisecond)
	return nil
}

func typeViaClipboard(text string) error {
	oldClipboard, backupErr := clipboard.ReadAll()
	if backupErr != nil {
		fmt.Println("Warning: Could not backup clipboard content:", backupErr)
	}

	if err := clipboard.WriteAll(text); err != nil {
		return fmt.Errorf("failed to write to clipboard: %w", err)
	}

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return fmt.Errorf("failed to create keyboard event: %w", err)
	}

	kb.SetKeys(keybd_event.VK_V)
	kb.HasCTRL(true)

	time.Sleep(50 * time.Millisecond)

	if err := kb.Press(); err != nil {
		return fmt.Errorf("failed to press keys: %w", err)
	}
	time.Sleep(50 * time.Millisecond)
	if err := kb.Release(); err != nil {
		return fmt.Errorf("failed to release keys: %w", err)
	}

	if backupErr == nil && oldClipboard != "" {
		time.Sleep(50 * time.Millisecond)
		_ = clipboard.WriteAll(oldClipboard)
	}

	return nil
}

func pressKey(key string) error {
	keyCode, err := getKeyCode(key)
	if err != nil {
		return err
	}

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return fmt.Errorf("failed to create keyboard event: %w", err)
	}

	kb.SetKeys(keyCode)

	func() {
		defer func() {
			releaseErr := kb.Release()
			if releaseErr != nil {
				fmt.Printf("Warning: failed to release key: %v\n", releaseErr)
			}
			time.Sleep(100 * time.Millisecond)
		}()

		time.Sleep(50 * time.Millisecond)
		if err := kb.Press(); err != nil {
			fmt.Printf("Warning: failed to press key: %v\n", err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}()

	time.Sleep(100 * time.Millisecond)

	return nil
}

func pressHotkey(combo string) error {
	parts := strings.Split(combo, "+")
	if len(parts) < 2 {
		return fmt.Errorf("hotkey must include at least one modifier and a key, separated by '+' (e.g., 'ctrl+s')")
	}

	key := parts[len(parts)-1]
	keyCode, err := getKeyCode(key)
	if err != nil {
		return err
	}

	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return fmt.Errorf("failed to create keyboard event: %w", err)
	}

	kb.SetKeys(keyCode)

	for i := 0; i < len(parts)-1; i++ {
		modifier := strings.ToLower(parts[i])
		switch modifier {
		case "ctrl":
			kb.HasCTRL(true)
		case "alt":
			kb.HasALT(true)
		case "shift":
			kb.HasSHIFT(true)
		case "win":
			kb.HasSuper(true)
		default:
			return fmt.Errorf("unsupported modifier: %s. Supported modifiers are: ctrl, alt, shift, win", modifier)
		}
	}

	func() {
		defer func() {
			releaseErr := kb.Release()
			if releaseErr != nil {
				fmt.Printf("Warning: failed to release hotkey: %v\n", releaseErr)
			}
			time.Sleep(100 * time.Millisecond)
		}()

		time.Sleep(50 * time.Millisecond)
		if err := kb.Press(); err != nil {
			fmt.Printf("Warning: failed to press hotkey: %v\n", err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}()

	time.Sleep(100 * time.Millisecond)

	return nil
}

func getKeyCode(key string) (int, error) {
	key = strings.TrimSpace(key)
	key = strings.Trim(key, "\"'")
	key = strings.ToLower(key)

	if code, ok := models.KeyMap[key]; ok {
		return code, nil
	}

	return 0, fmt.Errorf("unsupported key: %s", key)
}

func moveMouse(x, y int) error {
	screenWidth := int(win.GetSystemMetrics(win.SM_CXSCREEN))
	screenHeight := int(win.GetSystemMetrics(win.SM_CYSCREEN))

	if x < 0 || x > screenWidth || y < 0 || y > screenHeight {
		return fmt.Errorf("coordinates (%d, %d) are outside screen bounds (0--%d, 0--%d)",
			x, y, screenWidth, screenHeight)
	}

	win.SetCursorPos(int32(x), int32(y))

	time.Sleep(50 * time.Millisecond)

	return nil
}

func mouseClick(button string) error {
	var down, up uint32

	switch button {
	case "left":
		down = win.MOUSEEVENTF_LEFTDOWN
		up = win.MOUSEEVENTF_LEFTUP
	case "right":
		down = win.MOUSEEVENTF_RIGHTDOWN
		up = win.MOUSEEVENTF_RIGHTUP
	case "middle":
		down = win.MOUSEEVENTF_MIDDLEDOWN
		up = win.MOUSEEVENTF_MIDDLEUP
	default:
		return fmt.Errorf("unsupported mouse button: %s. Supported buttons are: left, right, middle", button)
	}

	var input win.MOUSE_INPUT
	input.Type = win.INPUT_MOUSE

	input.Mi.DwFlags = down
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
	time.Sleep(50 * time.Millisecond)

	input.Mi.DwFlags = up
	win.SendInput(1, unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
	time.Sleep(50 * time.Millisecond)

	return nil
}
