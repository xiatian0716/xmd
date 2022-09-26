package protoToCode

import (
	"fmt"
	"os/exec"
	"testing"
)

// exec.Command("cmd.exe", "/c", arg)
func CmdCheck(b *testing.B) {
	//要执行的Windows的指令
	//command := exec.Command("notepad")
	c := exec.Command("cmd.exe", "/c", fmt.Sprint("D:\\soft\\protoc-3.14.0-win64\\bin\\protoc.exe --version"))
	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}
