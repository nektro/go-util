package mbpp

import (
	"github.com/nektro/go-util/ansi/style"
	"github.com/nektro/go-util/util"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

//
func barA(task, name string) *mpb.Bar {
	return progress.AddBar(1,
		mpb.BarStyle(BarStyle),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
			&decorBarColor{},
		),
		mpb.AppendDecorators(
			decor.Name(style.ResetFgColor, decor.WC{W: 0}),
			decor.Name(PercentageStyle, decor.WC{W: 0}),
			decor.OnComplete(decor.NewPercentage("%.3f", decor.WCSyncSpace), "100%"),
			decor.Name(style.ResetFgColor, decor.WC{W: 0}),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(ETAStyle, decor.WC{W: 0}),
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
			decor.Name(style.ResetFgColor, decor.WC{W: 0}),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(util.TrimLen(name, 160), decor.WCSyncSpaceR),
		),
	)
}

func barB(task, name string) *mpb.Bar {
	return progress.AddBar(1,
		mpb.BarStyle(BarStyle),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.CountersKibiByte("% .2f / % .2f", decor.WCSyncWidth),
			&decorBarColor{},
		),
		mpb.AppendDecorators(
			decor.Name(style.ResetFgColor, decor.WC{W: 0}),
			decor.Name(PercentageStyle, decor.WC{W: 0}),
			decor.OnComplete(decor.NewPercentage("%.3f", decor.WCSyncSpace), "100%"),
			decor.Name(style.ResetFgColor, decor.WC{W: 0}),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(ETAStyle, decor.WC{W: 0}),
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
			decor.Name(style.ResetFgColor, decor.WC{W: 0}),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(util.TrimLen(name, 160), decor.WCSyncSpaceR),
		),
	)
}
