package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func e2() {
	bits := flag.String("b", "256", "sha256/384/512 digest")
	flag.Parse()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		switch *bits {
		case "384":
			fmt.Printf("%x\n", sha512.Sum384([]byte(text)))
		case "512":
			fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		}
	}
}
