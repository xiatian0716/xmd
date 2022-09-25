package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aCmd cobra样例
var aCmd = &cobra.Command{
	Use:   "a",
	Short: "xmd样例",
	Long:  `xmd子命令样例`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("你好，这是xmd子命令样例！")
	},
}

func init() {
	RootCmd.AddCommand(aCmd)
	//在这里，您将定义标志和配置设置。
	//Cobra支持用于此命令的Flag以及所有子命令，例如：
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")
	//Cobra支持本地标志，只有在执行此命令时才会运行直接调用，例如：
	//helloCmd.Flags（）。BoolP（“toggle”，“t”，false，“切换的帮助消息”）
}
