package protoToCode

import (
	_ "embed"
	"errors"
	"github.com/spf13/cobra"
	"github.com/xiatian0716/xmd/cmd"
	"github.com/xiatian0716/xmd/gRPC/protoToCode/hello"
	"github.com/xiatian0716/xmd/gRPC/protoToCode/hello/client"
	"github.com/xiatian0716/xmd/gRPC/protoToCode/hello/service"
	"github.com/xiatian0716/xmd/gRPC/protoToCode/protoGen"
	"github.com/xiatian0716/xmd/gRPC/protoToCode/tools"
	"github.com/xiatian0716/xmd/util/enhance"
	"strings"
)

var (
	GoModInit *string
	ProtoFile *string
)

var protoToCodeCmd = &cobra.Command{
	Use:   "protoC",
	Short: "通过proto模板生成代码模板",
	Long:  `通过指定的proto模板生成对应的代码模板`,
	RunE:  protoToCode,
}

func ProtoToCodeCmdSetup() {
	cmd.RootCmd.AddCommand(protoToCodeCmd)
	ProtoFile = protoToCodeCmd.Flags().StringP("protofile", "f", "", "选择Proto文件")
	GoModInit = protoToCodeCmd.Flags().StringP("gomodinit", "g", "example.com", "初始化go mod init")
}

func protoToCode(_ *cobra.Command, args []string) error {
	var projectName string
	protoFileName := *ProtoFile
	gomodinit := *GoModInit
	//pwd, err := os.Getwd()

	// 创建输出文件夹
	splitProtoFileByDot := strings.Split(protoFileName, ".")
	if splitProtoFileByDot[len(splitProtoFileByDot)-1] == "proto" {
		// 获取proto的名称 作为输出项目的名称
		projectName = strings.Replace(protoFileName, ".proto", "", -1) // -1是全部替换
		err := tools.MakeGencodeDir(projectName)
		if err != nil {
			return err
		}
	} else {
		return errors.New("请输入正确的proto文件名称(--protofile)")
	}

	if len(gomodinit) == 0 {
		return errors.New("ZRPC: missing --gomodinit")
	}
	gomoddata := map[string]string{"gomodinit": gomodinit}

	// 生成proto
	err := protoGen.GenProto(protoFileName)
	if err != nil {
		return err
	}
	// 生成client
	outClientFileName := "./" + projectName + "/client/client.go" // hello/client/client
	err = protoGen.GenCodeTmpl(client.GetClientTpl(), outClientFileName, gomoddata)
	if err != nil {
		return err
	}
	// 生成service
	outServiceFileName := "./" + projectName + "/service/service.go"
	err = protoGen.GenCodeTmpl(service.GetServiceTpl(), outServiceFileName, gomoddata)
	if err != nil {
		return err
	}
	// 生成go mod
	gomodFileName := "./" + projectName + "/go.mod"
	err = protoGen.GenCodeTmpl(hello.GetGoModTpl(), gomodFileName, gomoddata)
	if err != nil {
		return err
	}
	workDir := "./" + projectName
	_, err = enhance.Run("go mod tidy", workDir)

	return nil
}
