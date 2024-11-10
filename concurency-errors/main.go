package main

import (
	"errors"
	"log"
)

func generator(input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			inputCh <- data
		}
	}()

	return inputCh
}

func consumer(dataCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range dataCh {
		resp, err := request(data)
		resultCh <- Result{resp, err}
	}
}

func request(data int) (int, error) {
	return data, errors.New("some error")
}

type Result struct {
	data int
	err  error
}

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	resultCH := make(chan Result)

	go consumer(generator(data), resultCH)

	for result := range resultCH {
		if result.err != nil {
			log.Println("error:", result.err)
		} else {
			log.Println("result:", result.data)
		}

	}
}
