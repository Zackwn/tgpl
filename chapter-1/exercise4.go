package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func e4() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		// CTRL + Z then Enter to exit Stdin
		e4CountLines(os.Stdin, counts)
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				log.Fatal(err)
			}
			e4CountLines(file, counts)
			file.Close()
		}
	}
	for line, count := range counts {
		c := 0
		filenames := ""
		sep := ""
		for filename, fc := range count {
			c += fc
			filenames += fmt.Sprintf("%v%v: %v", sep, filename, fc)
			sep = " "
		}
		if c > 1 {
			fmt.Printf("[%v] %v\n", filenames, line)
		}
	}
}

func e4CountLines(f *os.File, c map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		name, text := f.Name(), input.Text()
		// fmt.Println(name, text)
		if c[text] == nil {
			c[text] = make(map[string]int)
		}
		c[text][name]++
	}
}
