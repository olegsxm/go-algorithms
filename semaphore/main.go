package main

import (
	"log"
	"sync"
	"time"
)

type Semaphore struct {
	ch chan struct{}
}

func (s *Semaphore) Acquire() {
	s.ch <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.ch
}

func NewSemaphore(maxReq int) *Semaphore {
	return &Semaphore{
		make(chan struct{}, maxReq),
	}
}

func main() {
	var wg sync.WaitGroup

	semaphore := NewSemaphore(2)

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func(id int) {
			semaphore.Acquire()
			defer wg.Done()
			defer semaphore.Release()

			log.Printf("Run process %d", id)
			time.Sleep(time.Second * 1)
		}(i)
	}

	wg.Wait()
}
