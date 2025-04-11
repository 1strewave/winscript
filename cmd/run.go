package cmd

import (
	"fmt"
	"github.com/1strewave/winscript/internal/parser"
	"github.com/1strewave/winscript/internal/runtime"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run file",
	Short: "Run a .ws script file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		script, err := parser.ParseFile(args[0])
		if err != nil {
			fmt.Println("Error parsing script", err)
			return
		}
		runtime.Execute(script)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
