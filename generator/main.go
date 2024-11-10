package main

import (
	"fmt"
	"strings"
)

func generator(data []string) chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for _, e := range data {
			ch <- strings.ToUpper(e[:1]) + e[1:]
		}

	}()

	return ch
}

func handler(c chan string) {
	for s := range c {
		fmt.Println(s)
	}
}

func main() {
	dataSet := strings.Split("The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using 'Content here, content here'", " ")

	dataChan := generator(dataSet)

	handler(dataChan)
}
