package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func e8() {
	for _, url := range os.Args[1:] {
		if !e8HasPrefix(url, "https://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}

// strings.HasPrefix implementation
func e8HasPrefix(str, prefix string) bool {
	for i := 0; i < len(str) && i < len(prefix); i++ {
		if str[i] != prefix[i] {
			return false
		}
	}
	return true
}
