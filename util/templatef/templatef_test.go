package templatef

import (
	"fmt"
	"testing"
)

func TestTemplatef(t *testing.T) {
	t.Run("有颜色的cmd", func(t *testing.T) {
		fmt.Println(Blue("测试"))
	})

}
