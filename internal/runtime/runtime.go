package runtime

import (
	"context"
	"fmt"
	"os/exec"
	"time"
	"unicode/utf16"

	"github.com/1strewave/winscript/internal/models"
	"github.com/atotto/clipboard"
	"github.com/lxn/win"
	"github.com/micmonay/keybd_event"
)

type CommandHandler func(cmd models.Command) error

type CommandRegistry map[string]CommandHandler

func NewCommandRegistry() CommandRegistry {
	registry := CommandRegistry{
		"open":  handleOpen,
		"type":  handleType,
		"wait":  handleWait,
		"log":   handleLog,
		"focus": handleFocus,
	}
	return registry
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
