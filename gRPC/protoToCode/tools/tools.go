package tools

import (
	"fmt"
	"github.com/xiatian0716/xmd/util/enhance"
	"github.com/xiatian0716/xmd/util/filex"
	"golang.org/x/text/encoding/simplifiedchinese"
	"path/filepath"
)

func MakeGencodeDir(projectName string) error {
	// 项目目录
	GenProjectOutAbs, err := filepath.Abs("./" + projectName) //当前目录
	if err != nil {
		return err
	}
	err = filex.MkdirIfNotExist(GenProjectOutAbs)
	if err != nil {
		return err
	}
	// 服务端目录
	ServiceOutAbs := filepath.Join(GenProjectOutAbs, "./service")
	err = filex.MkdirIfNotExist(ServiceOutAbs)
	if err != nil {
		return err
	}
	// 客户端目录
	ClientOutAbs := filepath.Join(GenProjectOutAbs, "./client")
	err = filex.MkdirIfNotExist(ClientOutAbs)
	if err != nil {
		return err
	}
	// Proto生成文件
	ProtoOutAbs := filepath.Join(GenProjectOutAbs, "./proto")
	err = filex.MkdirIfNotExist(ProtoOutAbs)
	if err != nil {
		return err
	}
	return nil
}

func CheckSoftInstall() {
	//解决Windows下在CMD下执行Go出现中文乱码的解决方法
	_, err := enhance.Run("123s", ".")
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(fmt.Sprintf(`%s`, err)))
	str := string(decodeBytes)
	fmt.Println(str)
}
