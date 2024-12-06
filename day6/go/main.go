package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const OBSTACLE = '#'

type Point struct {
	x, y int
}

type State struct {
	x, y, direction int
}

type Grid [][]rune

func loadFile(filename string) (Grid, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var grid Grid
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	return grid, scanner.Err()
}

func isValid(grid Grid, x, y int) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}

func getStart(grid Grid) (int, int) {
	for i := range grid {
		for j, ch := range grid[i] {
			if ch == '^' {
				return i, j
			}
		}
	}
	return -1, -1
}

func simulatePath(grid Grid, startX, startY, startDir int) bool {
	directions := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make(map[State]bool)
	x, y, dir := startX, startY, startDir

	for {
		state := State{x, y, dir}
		if visited[state] {
			return true
		}
		visited[state] = true

		nextX := x + directions[dir].x
		nextY := y + directions[dir].y

		if !isValid(grid, nextX, nextY) {
			return false
		}

		if grid[nextX][nextY] == OBSTACLE {
			dir = (dir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}
}

func partA(grid Grid) int {
	visited := make(map[Point]bool)
	directions := []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	startX, startY := getStart(grid)
	x, y, dir := startX, startY, 0

	for {
		point := Point{x, y}
		visited[point] = true

		nextX := x + directions[dir].x
		nextY := y + directions[dir].y

		if !isValid(grid, nextX, nextY) {
			break
		}

		if grid[nextX][nextY] == OBSTACLE {
			dir = (dir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}

	return len(visited)
}

func partB(grid Grid) int {
	startX, startY := getStart(grid)
	loopPositions := 0
	gridCopy := make(Grid, len(grid))
	
	// Make a copy of the grid
	for i := range grid {
		gridCopy[i] = make([]rune, len(grid[i]))
		copy(gridCopy[i], grid[i])
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' || (i == startX && j == startY) {
				continue
			}

			// Simulate roadblock
			gridCopy[i][j] = OBSTACLE
			if simulatePath(gridCopy, startX, startY, 0) {
				loopPositions++
			}
			gridCopy[i][j] = '.'
		}
	}

	return loopPositions
}

func main() {
	grid, err := loadFile("../data/data.txt")
	if err != nil {
		fmt.Printf("Error loading file: %v\n", err)
		return
	}

	start := time.Now()
	fmt.Printf("Part A: %d\n", partA(grid))
	fmt.Printf("Time taken: %v\n", time.Since(start))

	start = time.Now()
	fmt.Printf("Part B: %d\n", partB(grid))
	fmt.Printf("Time taken: %v\n", time.Since(start))
}
