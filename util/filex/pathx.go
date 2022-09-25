package filex

import (
	"os/exec"
	"runtime"
	"strings"
)

const (
	binProtoc          = "protoc"
	binProtocGenGo     = "protoc-gen-go"
	binProtocGenGrpcGo = "protoc-gen-go-grpc"
)

// LookUpProtoc searches an executable protoc in the directories
// named by the PATH environment variable.
func LookUpProtoc() (string, error) {
	suffix := getExeSuffix()
	xProtoc := binProtoc + suffix
	return LookPath(xProtoc)
}

// LookUpProtocGenGo searches an executable protoc-gen-go in the directories
// named by the PATH environment variable.
func LookUpProtocGenGo() (string, error) {
	suffix := getExeSuffix()
	xProtocGenGo := binProtocGenGo + suffix
	return LookPath(xProtocGenGo)
}

// LookUpProtocGenGoGrpc searches an executable protoc-gen-go-grpc in the directories
// named by the PATH environment variable.
func LookUpProtocGenGoGrpc() (string, error) {
	suffix := getExeSuffix()
	xProtocGenGoGrpc := binProtocGenGrpcGo + suffix
	return LookPath(xProtocGenGoGrpc)
}

// LookPath searches for an executable named file in the
// directories named by the PATH environment variable,
// for the os windows, the named file will be spliced with the
// .exe suffix.
func LookPath(xBin string) (string, error) {
	suffix := getExeSuffix()
	if len(suffix) > 0 && !strings.HasSuffix(xBin, suffix) {
		xBin = xBin + suffix
	}

	bin, err := exec.LookPath(xBin)
	if err != nil {
		return "", err
	}
	return bin, nil
}

func getExeSuffix() string {
	if runtime.GOOS == "windows" {
		return ".exe"
	}
	return ""
}
