# WinScript

![Platform](https://img.shields.io/badge/platform-Windows-blue?logo=windows)
![Language](https://img.shields.io/badge/language-Go-00ADD8?logo=go)
![License](https://img.shields.io/github/license/1strewave/winscript)
![Status](https://img.shields.io/badge/status-alpha-red)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)

> ⚡ **WinScript** is a lightweight scripting language for Windows automation.  
> Inspired by AppleScript, written in Go, and designed to be simple, readable, and powerful.

---

## ✨ Features

- 🖥 Launch and control Windows applications
- 🎯 Simulate mouse movement and clicks
- ⌨️ Simulate keystrokes and text input
- ⏱ Simple `wait` and delay functionality
- 📜 Human-readable `.ws` syntax
- 🔧 Cross-compatible CLI with Cobra
- 🧩 Easy to extend with Go

---

## 🚀 Quick Example

```winscript
open "notepad.exe"
wait 2s
type "Hello from WinScript!"
move_mouse 400 300
click "left"
```

> 📁 Save your scripts as `.ws` files.

---

## 🔧 Installation

```bash
git clone https://github.com/yourusername/winscript.git
cd winscript
go build -o winscript
```

> ✅ Optionally, add the built binary folder to your system `PATH` for global access.

To install the path permanently via script, you can use `install.bat`.

---

## 💻 CLI Usage

WinScript provides a command-line interface using Cobra.

### 🔹 Available Commands:

```bash
winscript [file]           # Run a script file directly
winscript run script.ws    # Explicit 'run' command
winscript help             # Show available commands
winscript docs             # Show documentation keywords
winscript version          # Print current version
```

### 🔸 Examples:

```bash
winscript hello.ws
winscript run scripts/boot.ws
winscript docs
```

---

## 🧠 Language Keywords

| Keyword       | Description                      | Example                           |
|---------------|----------------------------------|-----------------------------------|
| `open`        | Launch an application            | `open "notepad.exe"`              |
| `type`        | Type a string                    | `type "Hello, world!"`            |
| `press`       | Press a single key               | `press "enter"`                   |
| `hotkey`      | Key combination                  | `hotkey "ctrl+s"`                 |
| `move_mouse`  | Move mouse cursor                | `move_mouse 500 300`              |
| `click`       | Mouse click                      | `click "left"`                    |
| `wait`        | Wait for time                    | `wait 2s`                         |
| `log`         | Output text to console           | `log "Starting script..."`        |

📘 More details coming soon in [`docs/commands.md`](docs/commands.md)

---

## 🛣 Roadmap

- [x] CLI interface with Cobra
- [x] Script parser and interpreter
- [x] Core automation commands
- [ ] Variables (`set name = "Alice"`)
- [ ] Conditional logic (`if`, `else`)
- [ ] Loops (`repeat`, `while`)
- [ ] User-defined functions
- [ ] REPL / interactive mode
- [ ] GUI interface
- [ ] Windows Installer (`.msi`)

---

## 🤝 Contributing

Pull requests, issues, ideas, and feedback are welcome.  
If you’d like to help shape the future of WinScript — let’s build it together!

---

## 🧑‍💻 Author

**WinScript** is a hobby project created by [@1strewave](https://github.com/1strewave) to explore language design and automation on Windows using Go.  
Made with ❤️ for code, simplicity, and creativity.
