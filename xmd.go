package main

import (
	"github.com/xiatian0716/xmd/cmd"
	"github.com/xiatian0716/xmd/gRPC/getProtoTmpl"
)

func main() {
	cmd.Execute()
}

func init() {
	getProtoTmpl.PrototmplGenCmdSetup()
}
