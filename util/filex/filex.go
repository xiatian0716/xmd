package filex

import (
	"github.com/xiatian0716/xmd/util/version"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NL defines a new line.
const (
	NL       = "\n"
	XmdDir   = ".Xmd"
	cacheDir = "cache"
)

var XmdHome string

// FileExists returns true if the specified file is exists.
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

// GetXmdHome returns the path value of the Xmd, the default path is ~/.Xmd, if the path has
// been set by calling the RegisterXmdHome method, the user-defined path refers to.
func GetXmdHome() (home string, err error) {
	defer func() {
		if err != nil {
			return
		}
		info, err := os.Stat(home)
		if err == nil && !info.IsDir() {
			os.Rename(home, home+".old")
			MkdirIfNotExist(home)
		}
	}()
	if len(XmdHome) != 0 {
		home = XmdHome
		return
	}
	home, err = GetDefaultXmdHome()
	return
}

// GetDefaultXmdHome returns the path value of the Xmd home where Join $HOME with .Xmd.
func GetDefaultXmdHome() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, XmdDir), nil
}

// GetCacheDir returns the cache dit of Xmd.
func GetCacheDir() (string, error) {
	XmdH, err := GetXmdHome()
	if err != nil {
		return "", err
	}

	return filepath.Join(XmdH, cacheDir), nil
}

// GetTemplateDir returns the category path value in XmdHome where could get it by GetXmdHome.
func GetTemplateDir(category string) (string, error) {
	home, err := GetXmdHome()
	if err != nil {
		return "", err
	}
	if home == XmdHome {
		// backward compatible, it will be removed in the feature
		// backward compatible start.
		beforeTemplateDir := filepath.Join(home, version.GetXmdVersion(), category)
		fs, _ := ioutil.ReadDir(beforeTemplateDir)
		var hasContent bool
		for _, e := range fs {
			if e.Size() > 0 {
				hasContent = true
			}
		}
		if hasContent {
			return beforeTemplateDir, nil
		}
		// backward compatible end.

		return filepath.Join(home, category), nil
	}

	return filepath.Join(home, version.GetXmdVersion(), category), nil
}

// InitTemplates creates template files XmdHome where could get it by GetXmdHome.
func InitTemplates(category string, templates map[string]string) error {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return err
	}

	if err := MkdirIfNotExist(dir); err != nil {
		return err
	}

	for k, v := range templates {
		if err := createTemplate(filepath.Join(dir, k), v, false); err != nil {
			return err
		}
	}

	return nil
}

// CreateTemplate writes template into file even it is exists.
func CreateTemplate(category, name, content string) error {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return err
	}
	return createTemplate(filepath.Join(dir, name), content, true)
}

// Clean deletes all templates and removes the parent directory.
func Clean(category string) error {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return err
	}
	return os.RemoveAll(dir)
}

// LoadTemplate gets template content by the specified file.
func LoadTemplate(category, file, builtin string) (string, error) {
	dir, err := GetTemplateDir(category)
	if err != nil {
		return "", err
	}

	file = filepath.Join(dir, file)
	if !FileExists(file) {
		return builtin, nil
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func createTemplate(file, content string, force bool) error {
	if FileExists(file) && !force {
		return nil
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func Copy(src, dest string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	defer f.Close()

	dir := filepath.Dir(dest)
	err = MkdirIfNotExist(dir)
	if err != nil {
		return err
	}
	w, err := os.Create(dest)
	if err != nil {
		return err
	}
	w.Chmod(os.ModePerm)
	defer w.Close()
	_, err = io.Copy(w, f)
	return err
}

func ReadLink(name string) (string, error) {
	return name, nil
}

// MkdirIfNotExist makes directories if the input path is not exists
func MkdirIfNotExist(dir string) error {
	if len(dir) == 0 {
		return nil
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}

	return nil
}
