package filex

import (
	"fmt"
	"testing"
)

func TestGetXmdHome(t *testing.T) {
	//category := "foo"

	t.Run("GetDefaultXmdHome", func(t *testing.T) {
		//home := t.TempDir()
		XmdHome, _ := GetDefaultXmdHome()
		fmt.Println(XmdHome)
		//dir := filepath.Join(home, category)
		//err := MkdirIfNotExist(dir)
		//if err != nil {
		//	return
		//}
		//templateDir, err := GetTemplateDir(category)
		//if err != nil {
		//	return
		//}
		//assert.Equal(t, dir, templateDir)
	})

}
