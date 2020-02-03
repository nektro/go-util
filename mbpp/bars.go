package mbpp

import (
	"github.com/nektro/go-util/util"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
)

//
func barA(task, name string) *mpb.Bar {
	return progress.AddBar(1,
		mpb.BarStyle(barStyle),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WCSyncSpace), ""),
			decor.Name(": ", decor.WC{W: 2}),
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(util.TrimLen(name, 160), decor.WCSyncSpaceR),
		),
	)
}

func barB(task, name string) *mpb.Bar {
	return progress.AddBar(1,
		mpb.BarStyle(barStyle),
		mpb.BarRemoveOnComplete(),
		mpb.PrependDecorators(
			decor.Name(task, decor.WCSyncSpaceR),
			decor.Name(": ", decor.WC{W: 2}),
			decor.CountersKibiByte("% .2f / % .2f", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.OnComplete(decor.Percentage(decor.WCSyncSpace), ""),
			decor.Name(": ", decor.WC{W: 2}),
			decor.OnComplete(decor.EwmaETA(decor.ET_STYLE_MMSS, 0, decor.WCSyncWidth), ""),
			decor.Name(": ", decor.WC{W: 2}),
			decor.Name(util.TrimLen(name, 160), decor.WCSyncSpaceR),
		),
	)
}
