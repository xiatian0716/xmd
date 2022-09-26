package getProtoTmpl

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/xiatian0716/xmd/cmd"
	"github.com/xiatian0716/xmd/gRPC/getProtoTmpl/prototmplGen"
	"github.com/xiatian0716/xmd/util/enhance"
)

var (
	VarStringOutput string
)

// RPCTemplate is the entry for generate rpc template
func RPCTemplate(_ *cobra.Command, _ []string) error {
	enhance.Warning("记得传入文件名字")
	protoFile := VarStringOutput

	if len(protoFile) == 0 {
		return errors.New("missing -o")
	}

	return prototmplGen.ProtoTmpl(protoFile)
}

var prototmplGenCmd = &cobra.Command{
	Use:   "protoT",
	Short: "proto模板",
	Long:  `proto模板，并指定模板名称`,
	RunE:  RPCTemplate,
}

func PrototmplGenCmdSetup() {
	cmd.RootCmd.AddCommand(prototmplGenCmd)
	VarStringOutput = *prototmplGenCmd.Flags().StringP("protoname", "p", "hello.proto", "生成proto文件的名字")
	//.BoolP(“toggle”，“t”，false，“切换的帮助消息”)
	//在这里，您将定义标志和配置设置。
	//Cobra支持用于此命令的Flag以及所有子命令，例如：
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")
	//Cobra支持本地标志，只有在执行此命令时才会运行直接调用，例如：
	//helloCmd.Flags().BoolP（“toggle”，“t”，false，“切换的帮助消息”）
}
