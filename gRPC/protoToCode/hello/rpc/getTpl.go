package rpc

import (
	_ "embed"
)

//go:embed rpc.tpl
var rpcTemplateText string

// GetTpl
func GetRpcTpl() string {
	return rpcTemplateText
}
