package mbpp

import (
	"time"

	"github.com/vbauerster/mpb"
)

type BarProxy struct {
	T int64
	B *mpb.Bar
	s time.Time
}

func (b *BarProxy) AddToTotal(by int64) {
	b.T += by
	b.B.SetTotal(b.T, false)
}

func (b *BarProxy) Increment(by int) {
	b.B.IncrBy(by, time.Since(b.s))
}

func (b *BarProxy) FinishNow() {
	b.B.SetTotal(b.T, true)
}
