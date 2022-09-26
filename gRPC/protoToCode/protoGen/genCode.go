package protoGen

import (
	_ "embed"
	"github.com/xiatian0716/xmd/util"
	"github.com/xiatian0716/xmd/util/filex"
	"path/filepath"
)

// GenCodeTmpl returns a sample of a hello project file
func GenCodeTmpl(embedTemplateFile, out string, data map[string]string) error {
	text := embedTemplateFile

	dir := filepath.Dir(out)
	err := filex.MkdirIfNotExist(dir)
	if err != nil {
		return err
	}

	err = util.With("t").Parse(text).SaveTo(data, out, false)
	return err
}
