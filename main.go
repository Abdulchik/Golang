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
		fmt.Print("Введите размер доски (1–99): ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте ещё раз.")
			continue
		}

		input = strings.TrimSpace(input)

		// Проверяем, что строка — целое число
		if !isValidInteger(input) {
			fmt.Println("Ошибка: введите целое число.")
			continue
		}

		// Преобразуем в число
		width, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка преобразования. Попробуйте ещё раз.")
			continue
		}

		// Проверяем диапазон: 1 ≤ width ≤ 99
		if width > 0 && width < 100 {
			break // Всё корректно — выходим из цикла
		} else {
			fmt.Println("Ошибка: размер доски должен быть от 1 до 99.")
		}
	}

	// Рисуем шахматную доску
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

// Проверяет, что строка — корректное целое число (с возможным знаком)
func isValidInteger(s string) bool {
	if s == "" {
		return false
	}

	// Проверяем знак
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
		if s == "" {
			return false // только знак без цифр
		}
	}

	// Все символы должны быть цифрами
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
