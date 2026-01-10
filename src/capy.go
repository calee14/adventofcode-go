package src

import (
	"fmt"
	"strconv"
	"strings"
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

func Day2Part1() {
	data, err := ReadFile("data/day2.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0

	for id, game := range data {
		possible := true

		// Strip the 'Game <id>:'
		// from game string
		game, _ := strings.CutPrefix(strings.Split(game, ":")[1], " ")
		sets := strings.SplitSeq(game, ";")
		for set := range sets {
			red, green, blue := 12, 13, 14
			mred, mgreen, mblue := 0, 0, 0
			for color := range strings.SplitSeq(set, ",") {
				color, _ = strings.CutPrefix(color, " ")
				num, err := strconv.Atoi(strings.Split(color, " ")[0])
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				if strings.Contains(color, "red") {
					red -= num
					mred += num
				} else if strings.Contains(color, "green") {
					green -= num
					mgreen += num
				} else if strings.Contains(color, "blue") {
					blue -= num
					mblue += num
				}
			}

			if red < 0 || green < 0 || blue < 0 {
				possible = false
			}
		}

		if possible {
			result += id + 1
		}
	}

	fmt.Println(result)
}

func Day2Part2() {
	data, err := ReadFile("data/day2.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0

	for _, game := range data {
		mred, mgreen, mblue := 0, 0, 0

		// Strip the 'Game <id>:'
		// from game string
		game, _ := strings.CutPrefix(strings.Split(game, ":")[1], " ")
		sets := strings.SplitSeq(game, ";")
		for set := range sets {
			for color := range strings.SplitSeq(set, ",") {
				color, _ = strings.CutPrefix(color, " ")
				num, err := strconv.Atoi(strings.Split(color, " ")[0])
				if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				if strings.Contains(color, "red") {
					mred = max(mred, num)
				} else if strings.Contains(color, "green") {
					mgreen = max(mgreen, num)
				} else if strings.Contains(color, "blue") {
					mblue = max(mblue, num)
				}
			}
		}

		result += mred * mgreen * mblue
	}

	fmt.Println(result)
}
