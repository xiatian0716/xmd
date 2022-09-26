package client

import (
	_ "embed"
)

//go:embed client.tpl
var clientTemplateText string

// GetTpl
func GetClientTpl() string {
	return clientTemplateText
}
