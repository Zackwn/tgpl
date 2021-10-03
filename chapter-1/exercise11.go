package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func e11() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go e11Fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func e11Fetch(url string, ch chan<- string) {
	// Cancelling = Context
	// To make a request with a specified context.Context, use NewRequestWithContext and DefaultClient.Do
	start := time.Now()
	timeout := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	fmt.Println(req.Cancel, cancel)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- err.Error()
		return
	}
	defer resp.Body.Close()
	bytesWritten, err := io.Copy(ioutil.Discard, resp.Body)
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, bytesWritten, url)
}
