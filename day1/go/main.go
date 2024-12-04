package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
    "time"
)

func readDataFile() ([]int, []int, error) {
    file, err := os.Open("../data/data.txt")
    if err != nil {
        return nil, nil, fmt.Errorf("error opening file: %v", err)
    }
    defer file.Close()

    var left, right []int
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        parts := strings.Split(scanner.Text(), "   ")
        if len(parts) != 2 {
            continue
        }

        leftNum, err := strconv.Atoi(strings.TrimSpace(parts[0]))
        if err != nil {
            return nil, nil, fmt.Errorf("error parsing left number: %v", err)
        }

        rightNum, err := strconv.Atoi(strings.TrimSpace(parts[1]))
        if err != nil {
            return nil, nil, fmt.Errorf("error parsing right number: %v", err)
        }

        left = append(left, leftNum)
        right = append(right, rightNum)
    }

    if err := scanner.Err(); err != nil {
        return nil, nil, fmt.Errorf("error reading file: %v", err)
    }

    sort.Ints(left)
    sort.Ints(right)

    return left, right, nil
}

func partA(left, right []int) int {
    totalDistance := 0
    for i := 0; i < len(left); i++ {
        if left[i] > right[i] {
            totalDistance += left[i] - right[i]
        } else if right[i] > left[i] {
            totalDistance += right[i] - left[i]
        }
    }
    return totalDistance
}

func partB(left, right []int) int {
    similarityScore := 0
    for _, leftNum := range left {
        count := 0
        for _, rightNum := range right {
            if leftNum == rightNum {
                count++
            }
        }
        similarityScore += leftNum * count
    }
    return similarityScore
}

func main() {
    start := time.Now()
    left, right, err := readDataFile()
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }

    fmt.Println("__________________")
    fmt.Printf("A: %d\n", partA(left, right))
    fmt.Printf("B: %d\n", partB(left, right))
    elapsed := time.Since(start)
    fmt.Println("Time taken:", elapsed)
    fmt.Println("__________________")
}
