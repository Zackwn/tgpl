package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/zackwn/tgpl/chapter-2/tempconv"
)

func main() {
	e2()
}

func e2() {
	if len(os.Args) >= 2 {
		for _, arg := range os.Args[1:] {
			unit, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(e2Convert(unit))
		}
	} else {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			text := input.Text()
			unit, err := strconv.ParseFloat(text, 64)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(e2Convert(unit))
		}
	}
}

func e2Convert(unit float64) string {
	c := tempconv.Celsius(unit)
	f := tempconv.Fahrenheit(unit)
	k := tempconv.Kelvin(unit)
	cs := fmt.Sprintf("%s = %s = %s\n", c, tempconv.CtoF(c), tempconv.CtoK(c))
	fs := fmt.Sprintf("%s = %s = %s\n", f, tempconv.FtoC(f), tempconv.FtoK(f))
	ks := fmt.Sprintf("%s = %s = %s", k, tempconv.KtoC(k), tempconv.KtoF(k))
	return cs + fs + ks
}
