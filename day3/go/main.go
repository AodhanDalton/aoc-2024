package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var mulRegex = regexp.MustCompile(`mul\(([0-9]+).([0-9]+)\)`)

func loadFile() string {
	data, err := os.ReadFile("../data/data.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func partA(fileData string) int {
	total := 0
	matches := mulRegex.FindAllStringSubmatch(fileData, -1)
	for _, match := range matches {
		left, _ := strconv.Atoi(match[1])
		right, _ := strconv.Atoi(match[2])
		total += left * right
	}
	return total
}

func partB(fileData string) int {
	total := 0
	enabled := true
	pos := 0

	for pos < len(fileData) {
		if strings.HasPrefix(fileData[pos:], "don't()") {
			enabled = false
			pos += 7
			continue
		}
		if strings.HasPrefix(fileData[pos:], "do()") {
			enabled = true
			pos += 4
			continue
		}
		if enabled {
			match := mulRegex.FindStringSubmatchIndex(fileData[pos:])
			if match != nil {
				submatch := mulRegex.FindStringSubmatch(fileData[pos:])
				left, _ := strconv.Atoi(submatch[1])
				right, _ := strconv.Atoi(submatch[2])
				total += left * right
                length := 4 + len(submatch[1]) + 1 + len(submatch[2]) + 1
				pos += length 
			} else {
				pos++
			}
		} else {
			pos++
		}
	}
	return total
}

func main() {
	fileData := strings.ReplaceAll(loadFile(), "\n", "")
	
	start := time.Now()
	fmt.Println(partA(fileData))
	fmt.Printf("Time Taken: %v\n", time.Since(start))
	
	start = time.Now()
	fmt.Println(partB(fileData))
	fmt.Printf("Time Taken: %v\n", time.Since(start))
}
