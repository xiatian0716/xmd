package hellopro

import (
	_ "embed"
)

//go:embed config.tpl
var configTemplateText string

//go:embed go.tpl
var gomodTemplateText string

// GetTpl
func GetConfigTpl() string {
	return configTemplateText
}

// GetTpl
func GetGoModTpl() string {
	return gomodTemplateText
}
