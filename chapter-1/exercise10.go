package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func e10() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go e10Fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func e10Fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- err.Error()
		return
	}
	defer resp.Body.Close()
	wd, _ := os.Getwd()
	f, err := os.Create(filepath.Join(wd, e10GetDomain(url)+"-content.txt"))
	if err != nil {
		ch <- err.Error()
		return
	}
	// write content to report file
	bytesWritten, err := io.Copy(f, resp.Body)
	if err != nil {
		ch <- err.Error()
		return
	}
	elapsed := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, bytesWritten, url)
}

func e10GetDomain(url string) string {
	domain := ""
	i := 0
	for i < len(url) && url[i] != ':' {
		i++
	}
	i += 2
	for i < len(url) && url[i] != '.' {
		domain += string(url[i])
		i++
	}
	return domain
}
