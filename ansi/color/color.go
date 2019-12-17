package color

import (
	"github.com/nektro/go-util/ansi/style"
)

func FgBlack(s string) string {
	return style.FgBlack + s + style.ResetFgColor
}
func FgRed(s string) string {
	return style.FgRed + s + style.ResetFgColor
}
func FgGreen(s string) string {
	return style.FgGreen + s + style.ResetFgColor
}
func FgYellow(s string) string {
	return style.FgYellow + s + style.ResetFgColor
}
func FgBlue(s string) string {
	return style.FgBlue + s + style.ResetFgColor
}
func FgMagenta(s string) string {
	return style.FgMagenta + s + style.ResetFgColor
}
func FgCyan(s string) string {
	return style.FgCyan + s + style.ResetFgColor
}
func FgWhite(s string) string {
	return style.FgWhite + s + style.ResetFgColor
}

func BgBlack(s string) string {
	return style.BgBlack + s + style.ResetBgColor
}
func BgRed(s string) string {
	return style.BgRed + s + style.ResetBgColor
}
func BgGreen(s string) string {
	return style.BgGreen + s + style.ResetBgColor
}
func BgYellow(s string) string {
	return style.BgYellow + s + style.ResetBgColor
}
func BgBlue(s string) string {
	return style.BgBlue + s + style.ResetBgColor
}
func BgMagenta(s string) string {
	return style.BgMagenta + s + style.ResetBgColor
}
func BgCyan(s string) string {
	return style.BgCyan + s + style.ResetBgColor
}
func BgWhite(s string) string {
	return style.BgWhite + s + style.ResetBgColor
}
