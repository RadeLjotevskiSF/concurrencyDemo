package main

import (
	"errors"
	"fmt"
	"time"
)

type FutureResult struct {
	success chan any
	failure chan error
}

func (fr *FutureResult) onSuccess(result any) {
	fr.success <- result
}

func (fr *FutureResult) onFailure(err error) {
	fr.failure <- err
}

func doSomething() (string, error) {
	// simulate some work
	time.Sleep(time.Second)

	// return a random result or an error
	if time.Now().Unix()%2 == 0 {
		return "success", nil
	} else {
		return "", errors.New("failure")
	}
}

func main() {
	futureResult := &FutureResult{
		success: make(chan any),
		failure: make(chan error),
	}

	// run the operation in a goroutine
	go func() {
		result, err := doSomething()
		if err != nil {
			futureResult.onFailure(err)
		} else {
			futureResult.onSuccess(result)
		}
	}()

	// do other stuff while waiting for the future result
	for i := 1; i <= 5; i++ {
		fmt.Printf("working on something else... (%d/5)\n", i)
		time.Sleep(time.Second)
	}

	// wait for the result from the future
	select {
	case result := <-futureResult.success:
		fmt.Printf("result: %v\n", result)
	case err := <-futureResult.failure:
		fmt.Printf("err: %v\n", err)
	}
}
