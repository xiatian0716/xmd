package templatef

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"runtime"
)

const (
	OsWindows = "windows"
)

func Blue(s string) string {
	if runtime.GOOS == OsWindows {
		return s
	}

	return aurora.BrightBlue(s).String()
}

func Green(s string) string {
	if runtime.GOOS == OsWindows {
		return s
	}

	return aurora.BrightGreen(s).String()
}

func Rainbow(s string) string {
	if runtime.GOOS == OsWindows {
		return s
	}
	s0 := s[0]
	return colorRender[int(s0)%(len(colorRender)-1)](s)
}

// rpadx adds padding to the right of a string.
func Rpadx(s string, padding int) string {
	template := fmt.Sprintf("%%-%ds", padding)
	return Rainbow(fmt.Sprintf(template, s))
}

var colorRender = []func(v interface{}) string{
	func(v interface{}) string {
		return aurora.BrightRed(v).String()
	},
	func(v interface{}) string {
		return aurora.BrightGreen(v).String()
	},
	func(v interface{}) string {
		return aurora.BrightYellow(v).String()
	},
	func(v interface{}) string {
		return aurora.BrightBlue(v).String()
	},
	func(v interface{}) string {
		return aurora.BrightMagenta(v).String()
	},
	func(v interface{}) string {
		return aurora.BrightCyan(v).String()
	},
}
