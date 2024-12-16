package main

import (
    "fmt"
	"os"
    "regexp"
	"strconv"
)

func main(){
	text, err := os.ReadFile("../input.txt")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fileContent := string(text)

	textPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	numPattern := regexp.MustCompile(`\d+`)
	textMatches := textPattern.FindAllString(fileContent, -1)

	var valueHolder int

	if len(textMatches) > 0 {
		for _, match := range textMatches {
			numMatch := numPattern.FindAllString(match, -1)

			num_one, _ := strconv.Atoi(numMatch[0])
			num_two, _ := strconv.Atoi(numMatch[1])

			valueHolder += num_one * num_two
		}
	}
	fmt.Print(valueHolder)
}