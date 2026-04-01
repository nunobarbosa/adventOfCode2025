package day01

import (
	"fmt"
	"strconv"
)

func BuildInput(input []string) ([]rune, []int, error) {
	directions := []rune{}
	distances := []int{}

	for _, rotation := range input {
		directions = append(directions, rune(rotation[0]))

		s, err := strconv.Atoi(rotation[1:])
		if err != nil {
			return nil, nil, err
		}
		distances = append(distances, s)
	}

	return directions, distances, nil
}

func Solve(part int, input []string) (string, error) {
	switch part {
	case 1:
		return solvePart1(input)
	case 2:
		return solvePart2(input)
	default:
		return "", fmt.Errorf("unsupported part: %d", part)
	}
}

func solvePart1(input []string) (string, error) {
	directions, distances, err := BuildInput(input)
	if err != nil {
		return "", err
	}

	pointingTo := 50
	hits := 0
	for i, direction := range directions {
		distance := distances[i]

		switch direction {
		case 'L':
			// subtract
			pointingTo = pointingTo - distance
		case 'R': // sum
			pointingTo = pointingTo + distance
		default:
			return "", fmt.Errorf("unsupported direction: %c", direction)
		}

		pointingTo %= 100

		if pointingTo < 0 {
			pointingTo += 100
		} else if pointingTo == 0 {
			hits += 1
		}

		fmt.Printf("The dial is rotated %c%d to point at %d\n", direction, distance, pointingTo)
	}

	return fmt.Sprintf("%d", hits), nil
}
func solvePart2(input []string) (string, error) {
	directions, distances, err := BuildInput(input)
	if err != nil {
		return "", err
	}

	pointingTo := 50
	var clicks int = 0
	for i, direction := range directions {
		previousPointingTo := pointingTo
		distance := distances[i]
		midClicks := 0

		switch direction {
		case 'L':
			// subtract
			pointingTo = pointingTo - distance

			// how many times does it pass 100?
			for i = previousPointingTo; i > pointingTo; i-- {
				if i%100 == 0 {
					midClicks++
				}
			}

		case 'R': // sum
			pointingTo = pointingTo + distance

			// how many time does it pass 100?
			for i = previousPointingTo; i < pointingTo; i++ {
				if i%100 == 0 {
					midClicks++
				}
			}
		default:
			return "", fmt.Errorf("unsupported direction: %c", direction)
		}

		pointingTo %= 100

		if pointingTo < 0 {
			pointingTo += 100
		}

		clicks += midClicks

		fmt.Printf("The dial is rotated %c%d to point at %d; during this rotation, it pointed at 0 %d times\n", direction, distance, pointingTo, midClicks)
	}

	return fmt.Sprintf("%d", clicks), nil
}

// The solution for part 2 I presented is very lazy.
// We could use math to do this, as shown below by AI.
// The above has O(n * distance), while below solution does O(n)

/*
From AI:
func countZeroHits(start, end int) int {
	if start < end {
			return floorDiv(end, 100) - floorDiv(start, 100)
	}
	if start > end {
			return floorDiv(start-1, 100) - floorDiv(end-1, 100)
	}
	return 0
}

func floorDiv(a, b int) int {
    q := a / b
	r := a % b
	if r != 0 && ((r < 0) != (b < 0)) {
			q--
	}
	return q
}
*/
