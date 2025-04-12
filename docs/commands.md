# WinScript Command Reference

This document lists all currently supported commands in WinScript along with usage examples.

---

## 🖥 Application Control

### `open "program.exe"`

Launches an application.

**Example:**

```winscript
open "notepad.exe"
```

---

## ⌨️ Keyboard Input

### `type "text"`

Types the given text as keyboard input.

**Example:**

```winscript
type "Hello, World!"
```

---

### `press "key"`

Simulates pressing a single key.

**Example:**

```winscript
press "enter"
```

---

### `hotkey "combo"`

Simulates a key combination (modifier + key).

**Example:**

```winscript
hotkey "ctrl+s"
```

Supported modifiers: `ctrl`, `alt`, `shift`, `win`

---

## 🖱 Mouse Control

### `move_mouse X Y`

Moves the mouse cursor to screen coordinates `(X, Y)`.

**Example:**

```winscript
move_mouse 500 300
```

---

### `click "button"`

Performs a mouse click.

**Example:**

```winscript
click "left"
click "right"
click "middle"
```

---

## ⏱ Timing

### `wait DURATION`

Pauses script execution for a specified duration.

**Example:**

```winscript
wait 2s
wait 500ms
```

Supports durations like `ms`, `s`, `m`.

---

## 🛠 Utility

### `log "message"`

Outputs a message to the console for debugging or information.

**Example:**

```winscript
log "Script started."
```

---

## ✅ Coming Soon

- `set` (Variables)
- `if` / `else` (Conditionals)
- `repeat` / `while` (Loops)
- Error handling
- Comments (`# this is a comment`)

---

*Last updated: April 2025*

> Want to suggest a new command? [Open an issue](https://github.com/1strewave/winscript/issues).
