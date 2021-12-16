package main

import (
	"crypto/sha256"
	"fmt"
)

func e1() {
	ch1 := sha256.Sum256([]byte("ch1"))
	ch2 := sha256.Sum256([]byte("ch2"))
	fmt.Printf("%x\n%x\n", ch1, ch2)
	fmt.Println(e1bitsDiff(&ch1, &ch2))
	fmt.Println(ch1 == ch2)
}

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func e1bitsDiff(ch1, ch2 *[sha256.Size]byte) int {
	var count int
	for i := 0; i < sha256.Size; i++ {
		count += int(pc[ch1[i]^ch2[i]])
	}
	return count
}
