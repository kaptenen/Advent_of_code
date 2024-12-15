package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    file, osErr := os.Open("../input.txt")

    if osErr != nil {
        fmt.Print(osErr)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var leftSlice, rightSlice []int

    for scanner.Scan() {
        line := scanner.Text()

        var left, right int
        fmt.Sscanf(line, "%d %d", &left, &right)

        leftSlice = append(leftSlice, left)
        rightSlice = append(rightSlice, right)
    }

    hashMap := make(map[int]int)

    for i := 0; i < len(leftSlice); i++ {
        var freq int
        lookup := leftSlice[i]

        for j := 0; j < len(rightSlice); j++ {

            if rightSlice[j] == lookup {
                freq++
            }
        }
        hashMap[lookup] = freq
        freq = 0
    }

    holder := 0

    for key, value := range hashMap {
        holder += key * value
    }

    fmt.Print(holder)
}
