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

// func main() {
// 	http.HandleFunc("/test/", func(w http.ResponseWriter, r *http.Request) {
// 		e11()
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("OK"))
// 	})
// 	http.HandleFunc("/info/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("current goroutines:", runtime.NumGoroutine())
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("OK"))
// 	})
// 	fmt.Println("initial goroutines:", runtime.NumGoroutine())
// 	http.ListenAndServe(":8080", nil)
// }

func e11() {
	start := time.Now()
	ch := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer close(ch)
	for _, url := range os.Args[1:] {
		go e11Fetch(url, ch, ctx)
	}
	fmt.Println(<-ch)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return
}

func e11Fetch(url string, ch chan<- string, ctx context.Context) {
	start := time.Now()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// ch will be closed when the context is cancelled
		// therefore can't send to ch
		// ch <- err.Error()
		return
	}
	defer resp.Body.Close()
	bytesWritten, err := io.Copy(ioutil.Discard, resp.Body)
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, bytesWritten, url)
}
