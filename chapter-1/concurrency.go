package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	results := make(chan int)
	go worker(jobs, results)
	t := 20
	for i := 0; i < t; i++ {
		jobs <- i
	}
	close(jobs)
	for r := range results {
		fmt.Println(r)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	var wg sync.WaitGroup
	for n := range jobs {
		wg.Add(1)
		go func(n int) {
			results <- fib(n)
			wg.Done()
		}(n)
	}
	wg.Wait()
	close(results)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
