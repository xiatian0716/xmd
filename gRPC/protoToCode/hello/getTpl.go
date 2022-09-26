package hello

import (
	_ "embed"
)

//go:embed go.tpl
var gomodTemplateText string

// GetTpl
func GetGoModTpl() string {
	return gomodTemplateText
}
