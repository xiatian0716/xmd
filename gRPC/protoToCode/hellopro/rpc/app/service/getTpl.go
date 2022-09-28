package service

import (
	_ "embed"
)

//go:embed hello_service.tpl
var helloServiceTemplateText string

// GetTpl
func GetHelloServiceTpl() string {
	return helloServiceTemplateText
}
