package main

import (
	"fmt"
)

func main() {
label:
	var width int
	fmt.Print("Введите размер доски: ")
	fmt.Scan(&width)

	if width < 100 {
		if width >= 0 {
			if width%1 == 0 {
				if width%2 == 0 {
					for i := 0; i < width; i++ {

						if i%2 == 0 {
							for j := 0; j < width/2; j++ {
								fmt.Print("# ")
							}
						} else {
							for j := 0; j < width/2; j++ {
								fmt.Print(" #")
							}
						}
						fmt.Println()
					}
				} else {
					fmt.Print("Ошибка: нечетные числа не поддерживаются.\n")
					fmt.Print("Введите размеры заново.\n")
					goto label
				}
			} else {
				fmt.Print("Ошибка: дробные числа не поддерживаются.\n")
				fmt.Print("Введите размеры заново.\n")
				goto label
			}
		} else {
			fmt.Print("Ошибка: отрицательные или равные 0 числа не поддерживаются.\n")
			fmt.Print("Введите размеры заново.\n")
			goto label
		}

	} else {
		fmt.Print("Ошибка: число должно быть меньше 100.\n")
		fmt.Print("Введите размеры заново.\n")
		goto label
	}
}
