package day04

import (
	"strconv"
)

func BuildInput(input []string) ([][]int, error) {
	grid := [][]int{}

	for _, l := range input {
		line := []int{}
		for _, roll := range l {
			if roll == '@' {
				line = append(line, 1)
			} else {
				line = append(line, 0)
			}
		}
		grid = append(grid, line)
	}

	return grid, nil
}

func Solve(part int, input []string) (string, error) {
	grid, _ := BuildInput(input)
	switch part {
	case 1:
		return solvePart1(grid)
	case 2:
		return solvePart2(grid)
	default:
		return "", nil
	}
}

func solvePart1(grid [][]int) (string, error) {
	rolls := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != 1 {
				continue
			}
			neighbours := countNeighbours(grid, x, y)
			if neighbours < 4 {
				rolls++
			}
		}
	}
	return strconv.Itoa(rolls), nil
}

func solvePart2(grid [][]int) (string, error) {
	rolls := 0
	removedRolls := len(grid) * len(grid[0])
	for removedRolls > 0 {
		currentRun := 0
		for y := range grid {
			for x := range grid[y] {
				if grid[y][x] != 1 {
					continue
				}
				neighbours := countNeighbours(grid, x, y)
				if neighbours < 4 {
					currentRun++
					grid[y][x] = 0
				}
			}
		}
		removedRolls = currentRun
		rolls += currentRun
	}
	return strconv.Itoa(rolls), nil
}

func countNeighbours(grid [][]int, x, y int) int {
	minY := max(0, y-1)
	minX := max(0, x-1)
	maxY := min(len(grid)-1, y+1)
	maxX := min(len(grid[0])-1, x+1)

	neighbours := 0
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if i == y && j == x {
				continue
			}
			neighbours += grid[i][j]
		}
	}
	return neighbours
}
