package src

import (
	"fmt"
	"math"
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
		return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
	}

	result := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			temp_i := i
			temp_j := j
			if grid[i][j] != '.' && unicode.IsDigit(rune(grid[i][j])) {
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

func read_number(i int, j int, grid [][]byte, in_region func(int, int) bool) int {
	number := 0
	// Start at the beginning
	// of the number
	for in_region(i, j) && unicode.IsDigit(rune(grid[i][j])) {
		j -= 1
	}
	j += 1

	for in_region(i, j) && unicode.IsDigit(rune(grid[i][j])) {

		number = number*10 + int(grid[i][j]-'0')
		grid[i][j] = '.'
		j += 1
	}

	return number
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

	result := 0

	for i := range len(grid) {
		for j := range len(grid[0]) {
			temp_i := i
			temp_j := j
			if grid[temp_i][temp_j] == '*' {
				in_region := func(x int, y int) bool {
					return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
				}
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

				prod := 1
				adj_numbers := 0
				for _, dir := range dirs {
					if in_region(i+dir.First, j+dir.Second) &&
						unicode.IsDigit(rune(grid[i+dir.First][j+dir.Second])) {
						prod *= read_number(i+dir.First, j+dir.Second, grid, in_region)
						adj_numbers += 1
					}
				}

				if adj_numbers == 2 {
					result += prod
				}
			}
		}
	}

	fmt.Println(result)
}

func Map[T any, R any](in []T, fn func(T) R) []R {
	out := make([]R, len(in))
	for i, v := range in {
		out[i] = fn(v)
	}
	return out
}

func Filter[T any](in []T, pred func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if pred(v) {
			out = append(out, v)
		}
	}
	return out
}

func ToSet[T comparable](items []T) map[T]struct{} {
	set := make(map[T]struct{}, len(items))
	for _, v := range items {
		set[v] = struct{}{}
	}
	return set
}

func Day4Part1() {
	data, err := ReadFile("data/day4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0
	for _, card := range data {
		winning_nums := Map(Filter(strings.Split(strings.Split(
			strings.Split(card, "|")[0],
			":")[1], " "),
			func(card string) bool {
				return card != ""
			}),
			func(card string) int {
				num, err := strconv.Atoi(card)
				if err != nil {
					fmt.Println("Error: ", err)
					return -1
				}
				return num
			})

		our_nums := Map(Filter(strings.Split(
			strings.Split(card, "|")[1], " "),
			func(card string) bool {
				return card != ""
			}),
			func(card string) int {
				num, err := strconv.Atoi(card)
				if err != nil {
					fmt.Println("Error: ", err)
					return -1
				}
				return num
			})
		winning_set := ToSet(winning_nums)
		score := 0
		for _, num := range our_nums {
			if _, ok := winning_set[num]; ok {
				score += 1
			}
		}

		if score > 2 {
			score = int(math.Pow(2, float64(score-1)))
		}
		result += score
	}

	fmt.Println(result)
}

func Day4Part2() {
	data, err := ReadFile("data/day4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	result := 0
	num_cards := make([]int, len(data))
	for i := range num_cards {
		num_cards[i] = 1
	}

	for i, card := range data {
		winning_nums := Map(Filter(strings.Split(strings.Split(
			strings.Split(card, "|")[0],
			":")[1], " "),
			func(card string) bool {
				return card != ""
			}),
			func(card string) int {
				num, err := strconv.Atoi(card)
				if err != nil {
					fmt.Println("Error: ", err)
					return -1
				}
				return num
			})

		our_nums := Map(Filter(strings.Split(
			strings.Split(card, "|")[1], " "),
			func(card string) bool {
				return card != ""
			}),
			func(card string) int {
				num, err := strconv.Atoi(card)
				if err != nil {
					fmt.Println("Error: ", err)
					return -1
				}
				return num
			})
		winning_set := ToSet(winning_nums)
		matches := 0
		for _, num := range our_nums {
			if _, ok := winning_set[num]; ok {
				matches += 1
			}
		}
		for d := 1; d <= matches; d++ {
			if i+d < len(data) {
				num_cards[i+d] += num_cards[i]
			}
		}
		result += num_cards[i]
	}

	fmt.Println(result)
}
