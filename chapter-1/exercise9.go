package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func e9() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(os.Stdout, resp.Body)
		fmt.Println("status:", resp.Status)
		resp.Body.Close()
	}
}
