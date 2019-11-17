package style

import (
	"github.com/nektro/go-util/ansi"
)

var (
	ResetAll = ansi.SGR(0)

	Bold      = ansi.SGR(1)
	Faint     = ansi.SGR(2)
	Italic    = ansi.SGR(3)
	Underline = ansi.SGR(4)
	BlinkSlow = ansi.SGR(5)
	BlinkFast = ansi.SGR(6)

	ResetFont = ansi.SGR(10)
	Font1     = ansi.SGR(11)
	Font2     = ansi.SGR(12)
	Font3     = ansi.SGR(13)
	Font4     = ansi.SGR(14)
	Font5     = ansi.SGR(15)
	Font6     = ansi.SGR(16)
	Font7     = ansi.SGR(17)
	Font8     = ansi.SGR(18)
	Font9     = ansi.SGR(19)

	UnderlineDouble = ansi.SGR(21)
	ResetIntensity  = ansi.SGR(22)
	ResetItalic     = ansi.SGR(23)
	ResetUnderline  = ansi.SGR(24)
	ResetBlink      = ansi.SGR(25)

	FgBlack      = ansi.SGR(30)
	FgRed        = ansi.SGR(31)
	FgGreen      = ansi.SGR(32)
	FgYellow     = ansi.SGR(33)
	FgBlue       = ansi.SGR(34)
	FgMagenta    = ansi.SGR(35)
	FgCyan       = ansi.SGR(36)
	FgWhite      = ansi.SGR(37)
	ResetFgColor = ansi.SGR(39)

	BgBlack      = ansi.SGR(40)
	BgRed        = ansi.SGR(41)
	BgGreen      = ansi.SGR(42)
	BgYellow     = ansi.SGR(43)
	BgBlue       = ansi.SGR(44)
	BgMagenta    = ansi.SGR(45)
	BgCyan       = ansi.SGR(46)
	BgWhite      = ansi.SGR(47)
	ResetBgColor = ansi.SGR(49)

	Framed         = ansi.SGR(51)
	Encircled      = ansi.SGR(52)
	Overlined      = ansi.SGR(53)
	ResetFrameEnci = ansi.SGR(54)
	ResetOverlined = ansi.SGR(55)
)
