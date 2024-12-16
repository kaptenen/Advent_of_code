package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	text, err := os.ReadFile("../input.txt")

	if err != nil {
		panic("Error reading file:")
		return
	}

	fileContent := string(text)

	match := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	numPattern := regexp.MustCompile(`\d+`)

	result := match.FindAllString(fileContent, -1)

	isEnabled := true
	var valueHolder int

	for i := 0; i < len(result); i++ {

		switch result[i] {
		case "do()":
			isEnabled = true
		case "don't()":
			isEnabled = false
		}

		if isEnabled {
			numbersMatch := numPattern.FindAllString(result[i], -1)

			if len(numbersMatch) > 1 {
				numOne, _ := strconv.Atoi(numbersMatch[0])
				numTwo, _ := strconv.Atoi(numbersMatch[1])
				valueHolder += numOne * numTwo
			}
		}
	}
	fmt.Print(valueHolder)
}
