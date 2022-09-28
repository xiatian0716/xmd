package main

import (
	"github.com/xiatian0716/xmd/cmd"
	"github.com/xiatian0716/xmd/gRPC/getProtoTmpl"
	"github.com/xiatian0716/xmd/gRPC/protoToCode"
)

func main() {
	cmd.Execute()
}

func init() {
	getProtoTmpl.PrototmplGenCmdSetup()
	protoToCode.ProtoToCodeCmdSetup()
	protoToCode.ProtoToCodeCmdProSetup()
}
