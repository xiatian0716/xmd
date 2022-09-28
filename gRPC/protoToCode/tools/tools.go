package tools

import (
	"fmt"
	"github.com/xiatian0716/xmd/util/enhance"
	"golang.org/x/text/encoding/simplifiedchinese"
	"strings"
)

//
//func MakeGencodeDir(projectName string) error {
//	// 项目目录
//	GenProjectOutAbs, err := filepath.Abs("./" + projectName) //当前目录
//	if err != nil {
//		return err
//	}
//	err = filex.MkdirIfNotExist(GenProjectOutAbs)
//	if err != nil {
//		return err
//	}
//	// 服务端目录
//	rpcOutAbs := filepath.Join(GenProjectOutAbs, "./rpc")
//	err = filex.MkdirIfNotExist(rpcOutAbs)
//	if err != nil {
//		return err
//	}
//	// 客户端目录
//	ClientOutAbs := filepath.Join(GenProjectOutAbs, "./client")
//	err = filex.MkdirIfNotExist(ClientOutAbs)
//	if err != nil {
//		return err
//	}
//	// Proto生成文件
//	ProtoOutAbs := filepath.Join(GenProjectOutAbs, "./proto")
//	err = filex.MkdirIfNotExist(ProtoOutAbs)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func MakeGencodeDirPro(projectName string) error {
//	// 项目目录
//	GenProjectOutAbs, err := filepath.Abs("./" + projectName) //当前目录
//	if err != nil {
//		return err
//	}
//	err = filex.MkdirIfNotExist(GenProjectOutAbs)
//	if err != nil {
//		return err
//	}
//	// 服务端目录
//	rpcOutAbs := filepath.Join(GenProjectOutAbs, "./rpc")
//	err = filex.MkdirIfNotExist(rpcOutAbs)
//	if err != nil {
//		return err
//	}
//	appOutAbs := filepath.Join(GenProjectOutAbs, "./rpc/app")
//	err = filex.MkdirIfNotExist(appOutAbs)
//	if err != nil {
//		return err
//	}
//	modelsOutAbs := filepath.Join(GenProjectOutAbs, "./rpc/models")
//	err = filex.MkdirIfNotExist(modelsOutAbs)
//	if err != nil {
//		return err
//	}
//	serviceOutAbs := filepath.Join(GenProjectOutAbs, "./rpc/service")
//	err = filex.MkdirIfNotExist(serviceOutAbs)
//	if err != nil {
//		return err
//	}
//	// 客户端目录
//	ClientOutAbs := filepath.Join(GenProjectOutAbs, "./client")
//	err = filex.MkdirIfNotExist(ClientOutAbs)
//	if err != nil {
//		return err
//	}
//	// Proto生成文件
//	ProtoOutAbs := filepath.Join(GenProjectOutAbs, "./proto")
//	err = filex.MkdirIfNotExist(ProtoOutAbs)
//	if err != nil {
//		return err
//	}
//	return nil
//}

func CheckSoftInstall() {
	//解决Windows下在CMD下执行Go出现中文乱码的解决方法
	_, err := enhance.Run("123s", ".")
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(fmt.Sprintf(`%s`, err)))
	str := string(decodeBytes)
	fmt.Println(str)
}

// Title 首字母大写 HelloWorld
func Title(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

// Untitle 驼峰命名法 helloWord
func Untitle(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToLower(s[:1]) + s[1:]
}
