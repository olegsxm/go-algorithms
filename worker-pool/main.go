package main

import (
	"log"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		log.Printf("worker %d starting job %d", id, job)
		time.Sleep(time.Second * 2)
		log.Printf("worker %d finish job %d", id, job)
		results <- job * 2
	}
}

func main() {
	taskCount := 5
	jobs := make(chan int, taskCount)
	results := make(chan int, taskCount)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= taskCount; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= taskCount; a++ {
		res := <-results
		log.Printf("Result: %d", res)
	}
}
