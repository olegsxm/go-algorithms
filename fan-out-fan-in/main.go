package main

import (
	"fmt"
	"sync"
	"time"
)

func generator(doneCh chan struct{}, data []int) chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)
		for _, n := range data {
			select {
			case <-doneCh:
				return
			case stream <- n:
			}
		}
	}()

	return stream
}

func add(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)

		for num := range inputCh {
			time.Sleep(time.Second * 3)

			result := num + 1

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}

	}()

	return resultStream
}

func multiply(doneCh chan struct{}, inputCh chan int) chan int {
	resultStream := make(chan int)

	go func() {
		defer close(resultStream)
		for num := range inputCh {
			result := num * 2

			select {
			case <-doneCh:
				return
			case resultStream <- result:
			}
		}
	}()

	return resultStream
}

func fanOut(doneCh chan struct{}, inputCh chan int, workers int) []chan int {
	resultChannels := make([]chan int, workers)

	for i := 0; i < workers; i++ {
		resultChannels[i] = add(doneCh, inputCh)
	}

	return resultChannels
}

func fanIn(doneCh chan struct{}, channels ...chan int) chan int {
	var wg sync.WaitGroup
	resultStream := make(chan int)

	for _, channel := range channels {
		copyCh := channel
		wg.Add(1)

		go func() {
			defer wg.Done()
			for val := range copyCh {
				select {
				case <-doneCh:
					return
				case resultStream <- val:
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resultStream)
	}()

	return resultStream
}

func main() {
	data := make([]int, 0, 100)

	for i := range 100 {
		data = append(data, i)
	}

	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, data)

	chanals := fanOut(doneCh, inputCh, 5)

	addResultCh := fanIn(doneCh, chanals...)

	resultCh := multiply(doneCh, addResultCh)

	for result := range resultCh {
		fmt.Println(result)
	}
}
