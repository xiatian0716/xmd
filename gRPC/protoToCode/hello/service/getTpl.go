package service

import (
	_ "embed"
)

//go:embed service.tpl
var serviceTemplateText string

// GetTpl
func GetServiceTpl() string {
	return serviceTemplateText
}
