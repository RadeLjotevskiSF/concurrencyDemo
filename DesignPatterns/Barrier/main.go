package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	barrier := make(chan struct{})

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Printf("goroutine %d is doing some work\n", id)
			wg.Done()
			<-barrier // wait for all goroutines to come to this barrier
			fmt.Printf("goroutine %d is continuing\n", id)
		}(i)
	}

	wg.Wait()

	// signal all goroutines to continue
	close(barrier)

	time.Sleep(time.Millisecond)
}
