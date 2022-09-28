package models

import (
	_ "embed"
)

//go:embed SysPing.tpl
var sysPingTemplateText string

// GetTpl
func GetSysPingTpl() string {
	return sysPingTemplateText
}
