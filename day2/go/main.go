package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafeSequence(numbers []int) bool {
	if len(numbers) < 2 {
		return false
	}

	increasing := numbers[1] > numbers[0]

	for i := 0; i < len(numbers)-1; i++ {
		diff := numbers[i+1] - numbers[i]

		if diff == 0 {
			return false
		}
		if increasing {
			if diff <= 0 || diff > 3 {
				return false
			}
		} else {
			if diff >= 0 || diff < -3 {
				return false
			}
		}
	}

	return true
}

func checkWithDampener(numbers []int) bool {
	if isSafeSequence(numbers) {
		return true
	}

	for i := range numbers {
		reduced := make([]int, 0, len(numbers)-1)
		reduced = append(reduced, numbers[:i]...)
		reduced = append(reduced, numbers[i+1:]...)

		if isSafeSequence(reduced) {
			return true
		}
	}

	return false
}

func readNumbers(line string) ([]int, error) {
	fields := strings.Fields(line)
	numbers := make([]int, 0, len(fields))

	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, fmt.Errorf("error converting string to int: %v", err)
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

func partA() (int, error) {
	file, err := os.Open("../data/data.txt")
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers, err := readNumbers(scanner.Text())
		if err != nil {
			return 0, err
		}

		if isSafeSequence(numbers) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeCount, nil
}

func partB() (int, error) {
	file, err := os.Open("../data/data.txt")
	if err != nil {
		return 0, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	safeCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		numbers, err := readNumbers(scanner.Text())
		if err != nil {
			return 0, err
		}

		if checkWithDampener(numbers) {
			safeCount++
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file: %v", err)
	}

	return safeCount, nil
}

func main() {
	resultA, err := partA()
	if err != nil {
		fmt.Printf("Error in part A: %v\n", err)
		return
	}
	fmt.Printf("Part A: %d\n", resultA)

	resultB, err := partB()
	if err != nil {
		fmt.Printf("Error in part B: %v\n", err)
		return
	}
	fmt.Printf("Part B: %d\n", resultB)
}
