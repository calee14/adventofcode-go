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

type IntPair struct {
	First  int
	Second int
}

func Day3Part1() {
	data, err := ReadFile("data/day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	grid := make([][]byte, len(data)) // Use byte for ASCII
	for i, line := range data {
		grid[i] = []byte(line)
	}

	in_region := func(x int, y int) bool {
		return x >= 0 && x < len(data) && y >= 0 && y < len(data[0])
	}

	result := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			temp_i := i
			temp_j := j
			if data[i][j] != '.' && unicode.IsDigit(rune(grid[i][j])) {
				has_adj_symbol := false
				number := 0
				dirs := []IntPair{
					{1, 0},
					{1, 1},
					{0, 1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
					{0, -1},
					{1, -1},
				}
				for in_region(temp_i, temp_j) && unicode.IsDigit(rune(grid[temp_i][temp_j])) {
					for _, dir := range dirs {
						if in_region(temp_i+dir.First, temp_j+dir.Second) &&
							grid[temp_i+dir.First][temp_j+dir.Second] != '.' &&
							!unicode.IsDigit(rune(grid[temp_i+dir.First][temp_j+dir.Second])) {
							has_adj_symbol = true
						}
					}

					number = number*10 + int(grid[temp_i][temp_j]-'0')
					grid[temp_i][temp_j] = '.'
					temp_j += 1

				}

				if has_adj_symbol {
					result += number
				}
			}
		}
	}

	fmt.Println(result)
}

func Day3Part2() {
	data, err := ReadFile("data/day3.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	grid := make([][]byte, len(data)) // Use byte for ASCII
	for i, line := range data {
		grid[i] = []byte(line)
	}

	in_region := func(x int, y int) bool {
		return x >= 0 && x < len(data) && y >= 0 && y < len(data[0])
	}

	result := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			temp_i := i
			temp_j := j
			if data[i][j] != '.' && unicode.IsDigit(rune(grid[i][j])) {
				has_adj_symbol := false
				number := 0
				dirs := []IntPair{
					{1, 0},
					{1, 1},
					{0, 1},
					{-1, 1},
					{-1, 0},
					{-1, -1},
					{0, -1},
					{1, -1},
				}
				for in_region(temp_i, temp_j) && unicode.IsDigit(rune(grid[temp_i][temp_j])) {
					for _, dir := range dirs {
						if in_region(temp_i+dir.First, temp_j+dir.Second) &&
							grid[temp_i+dir.First][temp_j+dir.Second] != '.' &&
							!unicode.IsDigit(rune(grid[temp_i+dir.First][temp_j+dir.Second])) {
							has_adj_symbol = true
						}
					}

					number = number*10 + int(grid[temp_i][temp_j]-'0')
					grid[temp_i][temp_j] = '.'
					temp_j += 1

				}

				if has_adj_symbol {
					result += number
				}
			}
		}
	}

	fmt.Println(result)
}
