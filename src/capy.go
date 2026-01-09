package src

import (
	"fmt"
	"unicode"
)

func Day1Part1() {
	data, err := ReadFile("data/day1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0

	for _, inputStr := range data {
		ld, rd := -1, -1
		for _, char := range inputStr {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				if ld == -1 {
					ld = digit
				}
				rd = digit
			}
		}
		result += ld*10 + rd
	}

	fmt.Println(result)
}

func Day1Part2() {
	data, err := ReadFile("data/day1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0

	is_str_number := func(str string, start int) int {
		numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
		for i, num := range numbers {
			if len(str) >= start+len(num) {
				if str[start:start+len(num)] == num {
					return i + 1
				}
			}
		}
		return 0
	}

	for _, inputStr := range data {
		ld, rd := -1, -1
		for i, char := range inputStr {
			if unicode.IsDigit(char) {
				digit := int(char - '0')
				if ld == -1 {
					ld = digit
				}
				rd = digit
			} else {
				digit := is_str_number(inputStr, i)
				if digit > 0 {
					if ld == -1 {
						ld = digit
					}
					rd = digit
				}
			}
		}
		result += ld*10 + rd
	}

	fmt.Println(result)
}
