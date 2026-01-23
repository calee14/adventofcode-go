package main

import (
	"adventofcode/src"
	"fmt"
)

func main() {
	solutions := []func(){
		src.Day1Part1,
		src.Day1Part2,
		src.Day2Part1,
		src.Day2Part2,
		src.Day3Part1,
		src.Day3Part2,
		src.Day4Part1,
		src.Day4Part2,
		src.Day5Part1,
	}
	for i, sol := range solutions {
		fmt.Printf("Solving Day %d Part %d\n", i/2, i%2+1)
		sol()
		fmt.Println()
	}
}
