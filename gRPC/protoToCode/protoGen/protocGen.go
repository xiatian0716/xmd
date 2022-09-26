package protoGen

import (
	"errors"
	"fmt"
	"github.com/xiatian0716/xmd/util/enhance"
	"os"
)

var (
	ProtoInstallMessage = `1.安装proto编译器 https://github.com/protocolbuffers/protobuf/releases
2.安装protoc的Golang gRPC插件
  go install google.golang.org/protobuf/cmd/protoc-gen-go
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
3.添加到环境变量的 PATH 变量中`
)

func GenProto(ProtoFileName string) error {
	// protoc --go_out=.  --go-grpc_out=. helloworld.proto
	if FileExists(ProtoFileName) {
		codeLine := fmt.Sprintf(`protoc --go_out=.  --go-grpc_out=. %s`, ProtoFileName)
		_, err := enhance.Run(codeLine, ".") // 生成proto文件
		if err != nil {
			return errors.New(ProtoInstallMessage)
		}
	} else {
		return errors.New("erro: proto文件不存在，请检查传入参数(-f)")
	}

	return nil
}

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}
