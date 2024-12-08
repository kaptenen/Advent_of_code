package main

import (
	"bufio"
    "fmt"
    "os"
    "sort"
)


func diff(a, b int) int{
    if a < b{
        return b - a
    }
    return a - b
}

func main(){
	file, osErr := os.Open("input.txt")

    if osErr != nil {
        fmt.Print(osErr)
    }
    defer file.Close()


    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var leftSlice []int
    var rightSlice []int

    for scanner.Scan() {
        line := scanner.Text()

        var left, right int
        fmt.Sscanf(line, "%d %d", &left, &right)

        leftSlice = append(leftSlice, left)
        rightSlice = append(rightSlice, right)
    }

    sort.Ints(leftSlice)
    sort.Ints(rightSlice)

    var totalDiff int

    for i := 0; i < len(leftSlice); i++ {
        difference := diff(leftSlice[i], rightSlice[i])

        totalDiff += difference
    }

    fmt.Print(totalDiff)
}