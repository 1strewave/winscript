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
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, " ", 2)
		cmd := models.Command{
			Name:     parts[0],
			LineNum:  lineNum,
			OrigText: line,
		}

		if len(parts) > 1 {
			cmd.Args = parseArgs(parts[1])
		}

		commands = append(commands, cmd)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return commands, nil
}

func parseArgs(s string) []string {
	var args []string
	var current strings.Builder
	inQuote := false
	escaped := false

	for i, r := range s {
		if escaped {
			current.WriteRune(r)
			escaped = false
			continue
		}

		switch r {
		case '\\':
			escaped = true
		case '"':
			inQuote = !inQuote
		case ' ':
			if inQuote {
				current.WriteRune(r)
			} else if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		default:
			current.WriteRune(r)
		}

		if i == len(s)-1 && current.Len() > 0 {
			args = append(args, current.String())
		}
	}

	return args
}
