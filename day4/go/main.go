package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var directions = [][2]int{
	{0, 1}, {1, 1}, {1, 0}, {1, -1},
	{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
}

func loadFile() []string {
	file, err := os.Open("../data/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	return grid
}

func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func partA(grid []string, rows, cols int) int {
	count := 0
	checkDirection := func(x, y, dx, dy int) bool {
		word := ""
		for i := 0; i < 4; i++ {
			currX, currY := x+i*dx, y+i*dy
			if !isValid(currX, currY, rows, cols) {
				return false
			}
			word += string(grid[currX][currY])
		}
		return word == "XMAS"
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for _, dir := range directions {
				if checkDirection(x, y, dir[0], dir[1]) {
					count++
				}
			}
		}
	}
	return count
}

func partB(grid []string, rows, cols int) int {
	count := 0
	checkDirection := func(x, y, dx, dy int) bool {
		x = x - dx
		y = y - dy
		word := ""
		for i := 0; i < 3; i++ {
			currX, currY := x+i*dx, y+i*dy
			if !isValid(currX, currY, rows, cols) {
				return false
			}
			word += string(grid[currX][currY])
		}
		return word == "MAS" || word == "SAM"
	}

	checkXPattern := func(x, y int) bool {
		if grid[x][y] != 'A' {
			return false
		}
		return checkDirection(x, y, -1, 1) &&
			checkDirection(x, y, 1, -1) &&
			checkDirection(x, y, -1, -1) &&
			checkDirection(x, y, 1, 1)
	}

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			if checkXPattern(x, y) {
				count++
			}
		}
	}
	return count
}

func main() {
	grid := loadFile()
	rows := len(grid)
	cols := len(grid[0])

	start := time.Now()
	fmt.Println(partA(grid, rows, cols))
	fmt.Printf("Time Taken: %v\n", time.Since(start))

	start = time.Now()
	fmt.Println(partB(grid, rows, cols))
	fmt.Printf("Time Taken: %v\n", time.Since(start))
}
