package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Show all supported keywords",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available commands:")
		fmt.Println("- open \"program.exe\"")
		fmt.Println("- type \"text\"")
		fmt.Println("- wait 1s")
		fmt.Println("- press \"enter\"")
		fmt.Println("- hotkey \"ctrl+s\"")
		fmt.Println("- move_mouse 500 300")
		fmt.Println("- click \"left\"")
	},
}

func init() {
	rootCmd.AddCommand(docsCmd)
}
