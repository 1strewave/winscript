package models

import (
	"github.com/micmonay/keybd_event"
)

type Command struct {
	Name     string
	Args     []string
	LineNum  int
	OrigText string
}

const (
	VK_PERIOD       = 190
	VK_BACKQUOTE    = 192
	VK_LEFTBRACKET  = 219
	VK_BACKSLASH    = 220
	VK_RIGHTBRACKET = 221
	VK_QUOTE        = 222
)

var KeyMap = map[string]int{
	"a": keybd_event.VK_A,
	"b": keybd_event.VK_B,
	"c": keybd_event.VK_C,
	"d": keybd_event.VK_D,
	"e": keybd_event.VK_E,
	"f": keybd_event.VK_F,
	"g": keybd_event.VK_G,
	"h": keybd_event.VK_H,
	"i": keybd_event.VK_I,
	"j": keybd_event.VK_J,
	"k": keybd_event.VK_K,
	"l": keybd_event.VK_L,
	"m": keybd_event.VK_M,
	"n": keybd_event.VK_N,
	"o": keybd_event.VK_O,
	"p": keybd_event.VK_P,
	"q": keybd_event.VK_Q,
	"r": keybd_event.VK_R,
	"s": keybd_event.VK_S,
	"t": keybd_event.VK_T,
	"u": keybd_event.VK_U,
	"v": keybd_event.VK_V,
	"w": keybd_event.VK_W,
	"x": keybd_event.VK_X,
	"y": keybd_event.VK_Y,
	"z": keybd_event.VK_Z,

	"0": keybd_event.VK_0,
	"1": keybd_event.VK_1,
	"2": keybd_event.VK_2,
	"3": keybd_event.VK_3,
	"4": keybd_event.VK_4,
	"5": keybd_event.VK_5,
	"6": keybd_event.VK_6,
	"7": keybd_event.VK_7,
	"8": keybd_event.VK_8,
	"9": keybd_event.VK_9,

	"f1":  keybd_event.VK_F1,
	"f2":  keybd_event.VK_F2,
	"f3":  keybd_event.VK_F3,
	"f4":  keybd_event.VK_F4,
	"f5":  keybd_event.VK_F5,
	"f6":  keybd_event.VK_F6,
	"f7":  keybd_event.VK_F7,
	"f8":  keybd_event.VK_F8,
	"f9":  keybd_event.VK_F9,
	"f10": keybd_event.VK_F10,
	"f11": keybd_event.VK_F11,
	"f12": keybd_event.VK_F12,

	"enter":     keybd_event.VK_ENTER,
	"return":    keybd_event.VK_ENTER,
	"tab":       keybd_event.VK_TAB,
	"space":     keybd_event.VK_SPACE,
	"backspace": keybd_event.VK_BACKSPACE,
	"delete":    keybd_event.VK_DELETE,
	"esc":       keybd_event.VK_ESC,
	"escape":    keybd_event.VK_ESC,
	"home":      keybd_event.VK_HOME,
	"end":       keybd_event.VK_END,
	"pageup":    keybd_event.VK_PAGEUP,
	"pagedown":  keybd_event.VK_PAGEDOWN,
	"insert":    keybd_event.VK_INSERT,
	"up":        keybd_event.VK_UP,
	"down":      keybd_event.VK_DOWN,
	"left":      keybd_event.VK_LEFT,
	"right":     keybd_event.VK_RIGHT,

	";":  keybd_event.VK_SEMICOLON,
	"=":  keybd_event.VK_EQUAL,
	",":  keybd_event.VK_COMMA,
	"-":  keybd_event.VK_MINUS,
	".":  VK_PERIOD,
	"/":  keybd_event.VK_SLASH,
	"`":  VK_BACKQUOTE,
	"[":  VK_LEFTBRACKET,
	"\\": VK_BACKSLASH,
	"]":  VK_RIGHTBRACKET,
	"'":  VK_QUOTE,
}
