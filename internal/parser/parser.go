package parser

import (
	"bufio"
	"github.com/1strewave/winscript/internal/models"
	"os"
	"strings"
)

func ParseFile(path string) ([]models.Command, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []models.Command
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		cmd := models.Command{Name: parts[0]}

		if len(parts) > 1 {
			args := parseArgs(parts[1])
			cmd.Args = args
		}

		commands = append(commands, cmd)
	}

	return commands, nil
}

func parseArgs(s string) []string {
	args := []string{}
	current := ""
	inQuote := false

	for _, r := range s {
		switch r {
		case '"':
			inQuote = !inQuote
			if !inQuote {
				args = append(args, current)
				current = ""
			}
		default:
			if inQuote {
				current += string(r)
			}
		}
	}
	return args
}
