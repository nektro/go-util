package mbpp

import "io"

type passThru struct {
	bar    *BarProxy
	reader io.Reader
	determ bool
}

func (pt *passThru) Read(p []byte) (int, error) {
	n, err := pt.reader.Read(p)
	taskSize += int64(n)
	if !pt.determ {
		pt.bar.addRaw(int64(n))
	}
	pt.bar.incRaw(n)
	return n, err
}
