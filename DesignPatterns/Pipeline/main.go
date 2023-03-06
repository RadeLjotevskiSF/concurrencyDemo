package main

import "fmt"

func generator(nums ...int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func square(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func sum(in chan int) int {
	sumValue := 0
	for n := range in {
		sumValue += n
	}
	return sumValue
}

func main() {
	// create slice of some numbers
	numbers := []int{1, 2, 3, 4, 5}

	// create the pipeline
	c1 := generator(numbers...)
	c2 := square(c1)
	result := sum(c2)

	fmt.Println(result) //55
}
