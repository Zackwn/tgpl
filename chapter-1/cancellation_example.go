package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

var max = 10
var min = 2

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	n := 2
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < n; i++ {
		go worker(fmt.Sprintf("worker %v", i), ch, ctx)
	}
	result := <-ch
	fmt.Println(result)
	cancel()
	close(ch)
	time.Sleep(10 * time.Second)
	return
}

func worker(name string, ch chan<- string, ctx context.Context) {
	timeout := time.Second * time.Duration(rand.Intn(max-min)+min)
	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled with %v timeout\n", name, timeout)
		return
	// case <-time.After(timeout):
	case <-work(timeout):
		ch <- fmt.Sprintf("%v done within %v\n", name, timeout)
	}
}

func work(timeout time.Duration) chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		time.Sleep(timeout)
		ch <- "done"
	}()
	return ch
}
