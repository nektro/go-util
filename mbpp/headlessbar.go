package mbpp

import (
	"bytes"
	"io"
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
		b.Done()
	}
}

func (b *HeadlessBar) Done() {
	b.w.Close()
}

func (b *HeadlessBar) AddToMax(by int64) {
	b.m += by
}
