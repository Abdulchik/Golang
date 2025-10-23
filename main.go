package main

import (
	"fmt"
)

func main() {
	var a int = 7
	var b int = 7

	if a%2 == 0 && b%2 == 0 {
		for i := 0; i < a; i++ {
			fmt.Print("\n")
			if i%2 == 0 {
				for o := 0; o < b; o++ {
					fmt.Print("# ")

				}
			} else {
				for j := 0; j < b-1; j++ {
					fmt.Print(" #")
				}
			}
		}

	} else {
		fmt.Print("Не четное число")
	}
}
