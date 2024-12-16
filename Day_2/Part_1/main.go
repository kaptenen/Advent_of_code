package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func isValidLine(seq []int) bool {
	isIncreasing := seq[0] < seq[1]

	for i := 0; i < len(seq)-1; i++ {
        curr, next := seq[i], seq[i+1]

        if curr == next || diff(curr, next) > 3 {
            return false
        }

        if isIncreasing && curr >= next || !isIncreasing && curr <= next {
            return false
        }
    }

    return true
}

func main() {
	file, osErr := os.Open("../input.txt")

	if osErr != nil {
		fmt.Print(osErr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers [][]int

	for scanner.Scan() {
		line := scanner.Text()
		lineNums := strings.Split(line, " ")
		intNums := make([]int, len(lineNums))

		for i, num := range lineNums {
			intNums[i],_ = strconv.Atoi(num)
		}

		numbers = append(numbers, intNums)
	}

	var safeLines int
	for i := 0; i < len(numbers); i++ {
		line := numbers[i]

		if isValidLine(line) {
			safeLines++
		}
	}
	fmt.Println(safeLines)
}
