package prototmplGen

import (
	_ "embed"
	"path/filepath"
	"strings"

	"github.com/xiatian0716/xmd/util"
	"github.com/xiatian0716/xmd/util/filex"
	"github.com/xiatian0716/xmd/util/stringx"
)

//go:embed rpc.tpl
var rpcTemplateText string
var category = "rpc"
var rpcTemplateFile = "template.tpl"

// ProtoTmpl returns a sample of a proto file
func ProtoTmpl(out string) error {
	protoFilename := filepath.Base(out)
	serviceName := stringx.From(strings.TrimSuffix(protoFilename, filepath.Ext(protoFilename)))
	text, err := filex.LoadTemplate(category, rpcTemplateFile, rpcTemplateText)
	if err != nil {
		return err
	}

	dir := filepath.Dir(out)
	err = filex.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

	err = util.With("t").Parse(text).SaveTo(map[string]string{
		"package":     serviceName.Untitle(),
		"serviceName": serviceName.Title(),
	}, out, false)
	return err
}
