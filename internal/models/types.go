package models

type Command struct {
	Name     string
	Args     []string
	LineNum  int
	OrigText string
}
