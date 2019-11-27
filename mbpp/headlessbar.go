package mbpp

import (
	"bytes"
	"io"
)

type HeadlessBar struct {
	w *io.PipeWriter
}

func (b *HeadlessBar) Increment(by int) {
	b.w.Write(bytes.Repeat([]byte{1}, by))
}

func (b *HeadlessBar) Done() {
	b.w.Close()
}
