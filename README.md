# WinScript

![Platform](https://img.shields.io/badge/platform-Windows-blue?logo=windows)
![Language](https://img.shields.io/badge/language-Go-00ADD8?logo=go)
![License](https://img.shields.io/github/license/1strewave/winscript)
![Status](https://img.shields.io/badge/status-alpha-red)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)

> **WinScript** is a simple scripting language for automating actions on Windows.  
> A lightweight AppleScript-style tool written in Go with a Cobra-powered CLI.

---

## Features

- Launch and control applications
- Simulate keystrokes and text input
- Mouse movement and clicks
- Simple delays (`wait`)
- Clean and readable syntax
- Command-line interface via `winscript`
- Built entirely in Go, easily extendable

---

## Example Script

```winscript
open "notepad.exe"
wait 2s
type "Hello from WinScript!"
move_mouse 400 300
click "left"
```

Scripts are saved with the `.ws` extension.

---

## Installation

```bash
git clone https://github.com/yourusername/winscript.git
cd winscript
go build -o winscript
```

Now the `winscript` CLI is available for use.

---

## CLI

WinScript uses a Cobra-based CLI. Available commands:

```bash
winscript [file]             # Run a script
winscript run script.ws      # Explicit script run
winscript version            # Show version
winscript help               # Display help
```

### Examples

```bash
winscript hello.ws
winscript run scripts/boot.ws
winscript version
```

---

## Language Syntax

| Command       | Description                      | Example                           |
|---------------|----------------------------------|-----------------------------------|
| `open`        | Launch an application            | `open "notepad.exe"`              |
| `type`        | Type a string                    | `type "Hello, world!"`            |
| `press`       | Press a key                      | `press "enter"`                   |
| `hotkey`      | Press key combination            | `hotkey "ctrl+s"`                 |
| `move_mouse`  | Move the mouse cursor            | `move_mouse 500 300`              |
| `click`       | Perform a mouse click            | `click "left"`                    |
| `wait`        | Pause execution                  | `wait 2s`                          |
| `log`         | Output to console                | `log "Script started"`            |

Full command list coming soon in `docs/commands.md`.

---

## Roadmap

- [x] Basic interpreter
- [x] CLI interface using Cobra
- [ ] Variables (`set name = "Alice"`)
- [ ] Conditionals (`if`, `else`)
- [ ] Loops (`repeat`, `while`)
- [ ] Functions and blocks
- [ ] REPL support
- [ ] GUI interface
- [ ] Installer (.msi) support

---

## Contributing

Pull requests are welcome!  
Open issues, suggest features, report bugs — all contributions are appreciated.

---

## Author

WinScript is a personal project designed to explore scripting language development and Windows automation.  
Made with love for simplicity, Go, and productivity.
