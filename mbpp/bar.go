package mbpp

import (
	"github.com/vbauerster/mpb"
)

type BarProxy struct {
	T int64
	B *mpb.Bar
}

func (b *BarProxy) AddToTotal(by int64) {
	b.T += by
	b.B.SetTotal(b.T, false)
}

func (b *BarProxy) Increment(by int) {
	b.B.IncrBy(by)
}

func (b *BarProxy) FinishNow() {
	b.B.SetTotal(b.T, true)
}
