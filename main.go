package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b int = 6

	if a%2 == 0 && b%2 == 0 {
		for i := 0; i < a; i++ {

			if i%2 == 0 {
				for o := 0; o < b; o++ {
					fmt.Print("# ")

				}
			} else {
				for j := 0; j < b-1; j++ {
					fmt.Print(" #")
				}
			}
			fmt.Print("\n")
		}

	} else {
		if a%2 != 0 {
			fmt.Print("Первое заданное число не четное")
		} else {
			fmt.Print("Второе заданное число не четное")
		}
	}
}
