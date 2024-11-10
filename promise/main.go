package main

import (
	"fmt"
	"time"
)

type PromiseResult[T any] struct {
	Result T
	Error  error
}

func Promise[T any](task func() (T, error)) chan PromiseResult[T] {
	result := make(chan PromiseResult[T])

	go func() {
		taskResult, err := task()
		result <- PromiseResult[T]{Result: taskResult, Error: err}
		close(result)
	}()

	return result
}

func main() {
	task := func() (int64, error) {
		time.Sleep(3 * time.Second)
		return time.Now().Unix(), nil
	}

	promise := Promise(task)

	fmt.Println("Doing some thing else")

	result := <-promise
	if result.Error != nil {
		fmt.Println("Promise error: ", result.Error)
		return
	}

	fmt.Println(result.Result)

}
