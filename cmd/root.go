package cmd

import (
	"fmt"
	"github.com/1strewave/winscript/internal/parser"
	"github.com/1strewave/winscript/internal/runtime"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "winscript [file]",
	Short: "WinScript is a Windows automation script runner",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			script, err := parser.ParseFile(args[0])
			if err != nil {
				fmt.Println("Error parsing script:", err)
				return
			}
			runtime.Execute(script)
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
