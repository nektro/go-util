package types

import "sync"

type Semaphore struct {
	ch chan int
}

func NewSemaphore(max int) *Semaphore {
	return &Semaphore{
		ch: make(chan int, max),
	}
}

var mt = new(sync.Mutex)

func (sem *Semaphore) Add() {
	sem.ch <- 1
}

func (sem *Semaphore) Done() {
	mt.Lock()
	<-sem.ch
	mt.Unlock()
}
