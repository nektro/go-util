package mbpp

import (
	"sync"
	"time"

	"github.com/vbauerster/mpb"
)

type BarProxy struct {
	T int64
	B *mpb.Bar
	s time.Time
	w *sync.WaitGroup
}

func (b *BarProxy) addRaw(by int64) {
	b.T += by
	b.B.SetTotal(b.T, false)
}

var mt = new(sync.Mutex)

func (b *BarProxy) incRaw(by int) {
	mt.Lock()
	b.B.IncrBy(by, time.Since(b.s))
	b.s = time.Now()
	mt.Unlock()
}

func (b *BarProxy) AddToTotal(by int64) {
	b.addRaw(by)
	b.w.Add(int(by))
}

func (b *BarProxy) Increment(by int) {
	b.incRaw(by)
	for i := 0; i < by; i++ {
		b.w.Done()
	}
}

func (b *BarProxy) FinishNow() {
	b.B.SetTotal(b.T, true)
}

func (b *BarProxy) Wait() {
	b.w.Wait()
}
