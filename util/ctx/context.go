package ctx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xiatian0716/xmd/util/enhance"
	"github.com/xiatian0716/xmd/util/filex"
	"go/build"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var errModuleCheck = errors.New("the work directory must be found in the go mod or the $GOPATH")

const goModuleWithoutGoFiles = "command-line-arguments"

var errInvalidGoMod = errors.New("invalid go module")

// ProjectContext is a structure for the project,
// which contains WorkDir, Name, Path and Dir
type ProjectContext struct {
	WorkDir string
	// Name is the root name of the project
	// eg: go-zero、greet
	Name string
	// Path identifies which module a project belongs to, which is module value if it's a go mod project,
	// or else it is the root name of the project, eg: github.com/zeromicro/go-zero、greet
	Path string
	// Dir is the path of the project, eg: /Users/keson/goland/go/go-zero、/Users/keson/go/src/greet
	Dir string
}

type Module struct {
	Path      string
	Main      bool
	Dir       string
	GoMod     string
	GoVersion string
}

func (m *Module) validate() error {
	if m.Path == goModuleWithoutGoFiles || m.Dir == "" {
		return errInvalidGoMod
	}
	return nil
}

// projectFromGoMod is used to find the go module and project file path
// the workDir flag specifies which folder we need to detect based on
// only valid for go mod project
func projectFromGoMod(workDir string) (*ProjectContext, error) {
	if len(workDir) == 0 {
		return nil, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return nil, err
	}

	workDir, err := filex.ReadLink(workDir)
	if err != nil {
		return nil, err
	}

	m, err := getRealModule(workDir, enhance.Run)
	if err != nil {
		return nil, err
	}
	if err := m.validate(); err != nil {
		return nil, err
	}

	var ret ProjectContext
	ret.WorkDir = workDir
	ret.Name = filepath.Base(m.Dir)
	dir, err := filex.ReadLink(m.Dir)
	if err != nil {
		return nil, err
	}

	ret.Dir = dir
	ret.Path = m.Path
	return &ret, nil
}

func getRealModule(workDir string, execRun enhance.RunFunc) (*Module, error) {
	data, err := execRun("go list -json -m", workDir)
	if err != nil {
		return nil, err
	}
	modules, err := decodePackages(strings.NewReader(data))
	if err != nil {
		return nil, err
	}
	for _, m := range modules {
		if strings.HasPrefix(workDir, m.Dir) {
			return &m, nil
		}
	}
	return nil, errors.New("no matched module")
}

func decodePackages(rc io.Reader) ([]Module, error) {
	var modules []Module
	decoder := json.NewDecoder(rc)
	for decoder.More() {
		var m Module
		if err := decoder.Decode(&m); err != nil {
			return nil, fmt.Errorf("invalid module: %v", err)
		}
		modules = append(modules, m)
	}

	return modules, nil
}

// Prepare checks the project which module belongs to,and returns the path and module.
// workDir parameter is the directory of the source of generating code,
// where can be found the project path and the project module,
func Prepare(workDir string) (*ProjectContext, error) {
	ctx, err := background(workDir)
	if err == nil {
		return ctx, nil
	}

	name := filepath.Base(workDir)
	_, err = enhance.Run("go mod init "+name, workDir)
	if err != nil {
		return nil, err
	}
	return background(workDir)
}

func background(workDir string) (*ProjectContext, error) {
	isGoMod, err := IsGoMod(workDir)
	if err != nil {
		return nil, err
	}

	if isGoMod {
		return projectFromGoMod(workDir)
	}
	return projectFromGoPath(workDir)
}

// IsGoMod is used to determine whether workDir is a go module project through command `go list -json -m`
func IsGoMod(workDir string) (bool, error) {
	if len(workDir) == 0 {
		return false, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return false, err
	}

	data, err := enhance.Run("go list -m -f '{{.GoMod}}'", workDir)
	if err != nil || len(data) == 0 {
		return false, nil
	}

	return true, nil
}

// projectFromGoPath is used to find the main module and project file path
// the workDir flag specifies which folder we need to detect based on
// only valid for go mod project
func projectFromGoPath(workDir string) (*ProjectContext, error) {
	if len(workDir) == 0 {
		return nil, errors.New("the work directory is not found")
	}
	if _, err := os.Stat(workDir); err != nil {
		return nil, err
	}

	workDir, err := filex.ReadLink(workDir)
	if err != nil {
		return nil, err
	}

	buildContext := build.Default
	goPath := buildContext.GOPATH
	goPath, err = filex.ReadLink(goPath)
	if err != nil {
		return nil, err
	}

	goSrc := filepath.Join(goPath, "src")
	if !filex.FileExists(goSrc) {
		return nil, errModuleCheck
	}

	wd, err := filepath.Abs(workDir)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(wd, goSrc) {
		return nil, errModuleCheck
	}

	projectName := strings.TrimPrefix(wd, goSrc+string(filepath.Separator))
	return &ProjectContext{
		WorkDir: workDir,
		Name:    projectName,
		Path:    projectName,
		Dir:     filepath.Join(goSrc, projectName),
	}, nil
}
