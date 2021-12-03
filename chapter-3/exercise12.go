package main

import "fmt"

func e12() {
	fmt.Println(e12IsAnagram("ananba", "abanan")) // true
	fmt.Println(e12IsAnagram("ac", "abb"))        // false
	fmt.Println(e12IsAnagram("abb", "abbc"))      // false
}

func e12IsAnagram(s1, s2 string) bool {
	s1c, s2c := make(map[rune]int), make(map[rune]int)
	for _, c := range s1 {
		s1c[c]++
	}
	for _, c := range s2 {
		s2c[c]++
	}
	for char, count := range s1c {
		if s2c[char] != count {
			return false
		}
	}
	return len(s1c) == len(s2c)
}
