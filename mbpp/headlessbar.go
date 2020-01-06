package mbpp

import (
	"bytes"
	"io"
	"math"
)

type HeadlessBar struct {
	w *io.PipeWriter
	m int64
	p int64
}

func (b *HeadlessBar) Increment(by int) {
	b.p += int64(by)
	b.w.Write(bytes.Repeat([]byte{1}, by))
	if b.p >= b.m {
		b.w.Close()
	}
}

// Done really should be called FinishNow to follow my own naming convention
// but here we are. Done completes the current HeadlessBar regardless of its
// current progress.
func (b *HeadlessBar) Done() {
	b.Increment(int(math.Max(float64(b.m-b.p), 0)))
}

func (b *HeadlessBar) AddToMax(by int64) {
	b.m += by
}
