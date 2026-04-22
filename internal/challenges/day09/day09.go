package day09

import (
	"fmt"
	"strconv"
	"strings"
)

type Tile struct {
	X int
	Y int
}

func Solve(part int, input []string) (string, error) {
	tiles, err := BuildInput(input)
	if err != nil {
		return "", err
	}
	switch part {
	case 1:
		return solvePart1(tiles)
	case 2:
		return "day 09 part 2 not implemented\n", nil
	default:
		return "", nil
	}
}

func BuildInput(input []string) ([]Tile, error) {
	tiles := []Tile{}
	for _, l := range input {
		rawValues := strings.Split(l, ",")
		if len(rawValues) != 2 {
			return nil, fmt.Errorf("invalid coordinate %q", l)
		}

		x, err := strconv.Atoi(rawValues[0])
		if err != nil {
			return nil, err
		}

		y, err := strconv.Atoi(rawValues[1])
		if err != nil {
			return nil, err
		}

		tiles = append(tiles, Tile{X: x, Y: y})
	}

	return tiles, nil
}

func solvePart1(tiles []Tile) (string, error) {
	maxArea := 0

	for i := range tiles {
		for j := i + 1; j < len(tiles); j++ {
			area := tiles[i].area(tiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return strconv.Itoa(maxArea), nil
}

func (t Tile) area(t1 Tile) int {
	minX := t.X
	maxX := t1.X
	if minX > maxX {
		minX, maxX = maxX, minX
	}

	minY := t.Y
	maxY := t1.Y
	if minY > maxY {
		minY, maxY = maxY, minY
	}

	return (maxX - minX + 1) * (maxY - minY + 1)
}
