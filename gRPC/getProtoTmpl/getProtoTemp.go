package getProtoTmpl

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiatian0716/xmd/cmd"
	"github.com/xiatian0716/xmd/gRPC/getProtoTmpl/prototmplGen"
)

var (
	protoFileName *string
)

// RPCTemplate is the entry for generate rpc template
func RPCTemplate(_ *cobra.Command, _ []string) error {
	protoFile := *protoFileName
	if len(protoFile) == 0 {
		return errors.New("缺少flag -f")
	}
	fmt.Println(fmt.Sprintf("正在生成：%s", protoFile))

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
	protoFileName = prototmplGenCmd.Flags().StringP("protofilename", "f", "hello.proto", "生成proto文件的名字")
	//.BoolP(“toggle”，“t”，false，“切换的帮助消息”)
	//在这里，您将定义标志和配置设置。
	//Cobra支持用于此命令的Flag以及所有子命令，例如：
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")
	//Cobra支持本地标志，只有在执行此命令时才会运行直接调用，例如：
	//helloCmd.Flags().BoolP（“toggle”，“t”，false，“切换的帮助消息”）
}
