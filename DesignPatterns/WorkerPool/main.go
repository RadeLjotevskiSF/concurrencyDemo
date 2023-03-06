package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Returning result
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 2
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// starting workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// sending jobs to workers
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)

	// collecting results
	for i := 1; i <= numJobs; i++ {
		result := <-results
		fmt.Printf("got result: %d\n", result)
	}
}
