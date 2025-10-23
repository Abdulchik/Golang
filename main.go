package main

import (
	"fmt"
)

func main() {
	var width, length int
	fmt.Print("Введите ширину: ")
	fmt.Scan(&width)
	fmt.Print("Введите длину: ")
	fmt.Scan(&length)
	if width == length {
		for i := 0; i < length; i++ {

			if i%2 == 0 {
				for j := 0; j < width; j++ {
					fmt.Print("# ")
				}
			} else {
				for j := 0; j < width; j++ {
					fmt.Print(" #")
				}
			}
			fmt.Println()
		}

	} else {

		fmt.Print("Длина и ширина должны быть равны для создания шахматной доски.")
	}
}
