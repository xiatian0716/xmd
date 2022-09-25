package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var testString string

// RootCmd 根命令行
var RootCmd = &cobra.Command{
	Use:     "xmd",
	Short:   "加强命令行工具集(cmd-plus)",
	Long:    `欢迎使用xmd`,
	Example: `xmd --argBool=True`,
	Version: `v 1.0.0`,
}

// Execute 将所有的子命令添加到根命令上，并且设置参数(flags)标签
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	testRootCmdBool := RootCmd.Flags().BoolP("argBool", "a", false, "通过命令行传给xmd递参数(Bool)")
	RootCmd.PersistentFlags().StringVar(&testString, "argString", "", "通过命令行传给xmd递参数(String)")
	if *testRootCmdBool {
		fmt.Printf(`打印命令行传递的参数：%t`, *testRootCmdBool)
	}

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	//
}
