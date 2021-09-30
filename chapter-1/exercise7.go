package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func e7() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		// io.Copy(resp.Body, os.Stdout)
		e7Copy(resp.Body, os.Stdout)
		resp.Body.Close()
	}
}

// io.Copy implementation
func e7Copy(in io.Reader, out io.Writer) {
	buffer := make([]byte, 32*1024)
	var err error
	for err == nil {
		bs, e := in.Read(buffer)
		err = e
		out.Write(buffer[:bs])
	}
}
