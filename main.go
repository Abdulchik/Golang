package main

import (
	"fmt"
)

func main() {
	var a int = 6

	if a%2 == 0 {
		for i := 0; i < a; i++ {

			if i%2 == 0 {
				for j := 0; j < a; j++ {
					fmt.Print("# ")

				}
			} else {
				for j := 0; j < a-1; j++ {
					fmt.Print(" #")
				}
			}
			fmt.Println()
		}

	} else {

		fmt.Print("Заданное число не четное")

	}
}
