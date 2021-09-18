package main

import (
	"fmt"
	"os"
)

func e2() {
	s, sep := "", ""
	for i, e := range os.Args {
		fmt.Println(i, ":", e)
		s += sep + e
		sep = " "
	}
	fmt.Println(s)
}
