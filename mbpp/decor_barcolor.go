package mbpp

import (
	"github.com/nektro/go-util/ansi/style"
	"github.com/vbauerster/mpb/decor"
)

type decorBarColor struct {
}

func (w *decorBarColor) Decor(s *decor.Statistics) string {
	if !ColoredBar {
		return ""
	}
	p := float64(s.Current) / float64(s.Total)
	if p < .5 {
		// red -> yellow
		x := int(p / .5 * 255)
		return style.Fg24bit(255, x, 0)
	}
	// yello -> green
	x := int((p - .5) / .5 * 255)
	return style.Fg24bit(255-x, 255, 0)
}

func (w *decorBarColor) Sync() (chan int, bool) {
	return nil, true
}

func (w *decorBarColor) GetConf() decor.WC {
	return decor.WC{W: 0}
}

func (w *decorBarColor) SetConf(decor.WC) {
}
