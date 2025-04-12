# Open Notepad
log "Script started"
open "notepad.exe"
wait 2s
type "Hello from WinScript! sigma boy"
wait 2s

# Change focus on window -> "type.ws - Notepad"
focus "type.ws – Notepad"
wait 1s  # Wait until focus done