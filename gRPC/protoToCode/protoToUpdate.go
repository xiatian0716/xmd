package protoToCode

import (
	_ "embed"
	"errors"
	"github.com/spf13/cobra"
	"github.com/xiatian0716/xmd/cmd"
	"github.com/xiatian0716/xmd/gRPC/protoToCode/protoGen"
	"strings"
)

var (
	protoFileU *string
)

var protoToUpdateCmd = &cobra.Command{
	Use:   "protoCU",
	Short: "通过proto模板更新grpc代码",
	Long:  `通过proto模板更新grpc代码`,
	RunE:  protoToUpdate,
}

func ProtoToUpdateCmdSetup() {
	cmd.RootCmd.AddCommand(protoToUpdateCmd)
	protoFileU = protoToUpdateCmd.Flags().StringP("protofile", "f", "", "选择Proto文件")
}

func protoToUpdate(_ *cobra.Command, args []string) error {
	protoFileName := *protoFileU

	// 获取proto的名称
	splitProtoFileByDot := strings.Split(protoFileName, ".")
	if splitProtoFileByDot[len(splitProtoFileByDot)-1] == "proto" {
		// 生成proto
		err := protoGen.GenProto(protoFileName)
		if err != nil {
			return err
		}
	} else {
		return errors.New("请输入正确的proto文件名称(--protofile)")
	}

	return nil
}
