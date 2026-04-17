package day08

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Coordinates struct {
	X int
	Y int
	Z int
}

func (c Coordinates) distance(c1 Coordinates) float64 {
	deltaX := c1.X - c.X
	deltaY := c1.Y - c.Y
	deltaZ := c1.Z - c.Z

	return math.Sqrt(float64(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ))
}

func BuildInput(input []string) ([]Coordinates, [][]float64, error) {
	coordinates := []Coordinates{}

	for _, l := range input {
		rawValues := strings.Split(l, ",")

		x, err := strconv.Atoi(rawValues[0])
		if err != nil {
			return nil, nil, err
		}

		y, err := strconv.Atoi(rawValues[1])
		if err != nil {
			return nil, nil, err
		}

		z, err := strconv.Atoi(rawValues[2])
		if err != nil {
			return nil, nil, err
		}

		coordinates = append(coordinates, Coordinates{X: x, Y: y, Z: z})
	}

	distances := make([][]float64, len(coordinates))
	for i := 0; i < len(coordinates); i++ {
		distances[i] = make([]float64, len(coordinates))
		for j := i; j < len(coordinates); j++ {
			d := coordinates[i].distance(coordinates[j])
			distances[i][j] = d
		}
		fmt.Printf("%v\n", distances[i])
	}

	return coordinates, distances, nil
}

func Solve(part int, input []string) (string, error) {
	coordinates, distances, err := BuildInput(input)
	if err != nil {
		return "", err
	}
	switch part {
	case 1:
		return solvePart1(coordinates, distances)
	case 2:
		return "day 08 part 2 not implemented\n", nil
	default:
		return "", nil
	}
}

func solvePart1(coordinates []Coordinates, distances [][]float64) (string, error) {
	return "TEST", nil
}
