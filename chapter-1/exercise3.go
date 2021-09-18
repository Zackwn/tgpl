package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func e3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	end := time.Since(start)
	fmt.Printf("strings.Join: %v\n", end.String())

	start = time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	end = time.Since(start)
	fmt.Printf("for: %v\n", end.String())

	start = time.Now()
	s, sep = "", ""
	for _, e := range os.Args[1:] {
		s += sep + e
		sep = " "
	}
	fmt.Println(s)
	end = time.Since(start)
	fmt.Printf("for_range: %v\n", end.String())
}
