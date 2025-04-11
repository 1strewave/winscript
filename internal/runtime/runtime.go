package runtime

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/1strewave/winscript/internal/models"
)

func Execute(commands []models.Command) {
	for _, cmd := range commands {
		switch cmd.Name {
		case "open":
			if len(cmd.Args) == 1 {
				exec.Command("cmd", "/C", "start", cmd.Args[0]).Start()
			}
		case "type":
			if len(cmd.Args) == 1 {
				fmt.Printf("Typing: %s\n", cmd.Args[0])
				//TODO:add keyboad thing
			}
		case "wait":
			if len(cmd.Args) == 1 {
				d, _ := time.ParseDuration(cmd.Args[0])
				time.Sleep(d)
			}
		case "log":
			if len(cmd.Args) == 1 {
				fmt.Println(cmd.Args[0])
			}
		default:
			fmt.Printf("Unknown command: %s\n", cmd.Name)
		}
	}
}
