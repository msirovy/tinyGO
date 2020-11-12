package main

import (
  "time"
	"fmt"
	"strings"
)

func main() {
	min := 100
	max := 1000

	step := 50


	for {
		for i := min; i < max; i=i+step {
			fmt.Println(strings.Repeat("+", i/50))
      time.Sleep(time.Millisecond * time.Duration(i))
			fmt.Println(strings.Repeat("-", i/50))
      time.Sleep(time.Millisecond * time.Duration(i))
		}

		for i := max; i > min; i=i-step {
			fmt.Println(strings.Repeat("+", i/50))
      time.Sleep(time.Millisecond * time.Duration(i))
			fmt.Println(strings.Repeat("-", i/50))
      time.Sleep(time.Millisecond * time.Duration(i))
		}

	}
}
