package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var width int
	var input string

	for {
		fmt.Print("Введите размер доски (от 1 до 100): ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте ещё раз.")
			continue
		}

		input = strings.TrimSpace(input)

		if !isValidInteger(input) {
			fmt.Println("Ошибка: введите целое число.")
			continue
		}

		width, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка преобразования. Попробуйте ещё раз.")
			continue
		}

		if width > 0 && width < 101 {
			break
		} else {
			fmt.Println("Ошибка: размер доски должен быть от 1 до 100.")
		}
	}

	for row := 0; row < width; row++ {
		for col := 0; col < width; col++ {
			if (row+col)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func isValidInteger(s string) bool {
	if s == "" {
		return false
	}

	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if s == "" {
			return false
		}
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
