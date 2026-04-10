package day07

import (
	"fmt"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func BuildInputPart1(input []string) (Position, Position, map[Position]bool, error) {
	initialPosition := Position{X: 0, Y: 0}
	splitters := map[Position]bool{}
	mapSize := Position{X: len(input[0]), Y: len(input)}

	for y, l := range input {
		for x, c := range l {
			switch c {
			case 'S':
				initialPosition.X = x
				initialPosition.Y = y
			case '^':
				splitters[Position{X: x, Y: y}] = true
			}
		}
	}

	return initialPosition, mapSize, splitters, nil
}

func Solve(part int, input []string) (string, error) {
	initialPosition, mapSize, splitters, _ := BuildInputPart1(input)
	switch part {
	case 1:
		return solvePart1(initialPosition, mapSize, splitters)
	case 2:
		return solvePart2(initialPosition, mapSize, splitters)
	default:
		return "", nil
	}
}

func solvePart1(initialPosition, mapSize Position, splitters map[Position]bool) (string, error) {
	previousPositions := map[Position]bool{initialPosition: true}
	splits := 0

	for y := initialPosition.Y + 1; y < mapSize.Y; y++ {
		newPreviousPositions := map[Position]bool{}
		for x := 0; x < mapSize.X; x++ {
			previousPosition := Position{X: x, Y: y - 1}
			currentPosition := Position{X: x, Y: y}

			_, hasPreviousPosition := previousPositions[previousPosition]
			if !hasPreviousPosition {
				continue
			}

			_, hasSplitter := splitters[currentPosition]
			if !hasSplitter {
				newPreviousPositions[currentPosition] = true
				continue
			}
			splits++

			leftPosition := Position{X: x - 1, Y: y}
			if _, hasLeftPosition := newPreviousPositions[leftPosition]; !hasLeftPosition {
				newPreviousPositions[leftPosition] = true
			}

			rightPosition := Position{X: x + 1, Y: y}
			newPreviousPositions[rightPosition] = true
		}
		fmt.Printf("%v\n", newPreviousPositions)
		previousPositions = newPreviousPositions
	}
	return strconv.Itoa(splits), nil
}

func solvePart2(initialPosition, mapSize Position, splitters map[Position]bool) (string, error) {
	paths := map[Position]int{initialPosition: 1}

	for y := initialPosition.Y; y < mapSize.Y; y++ {
		newPaths := map[Position]int{}
		for position, counter := range paths {
			nextPosition := Position{X: position.X, Y: position.Y + 1}
			if _, hit := splitters[nextPosition]; !hit {
				newPaths[nextPosition] += counter
				continue
			}

			left := Position{X: nextPosition.X - 1, Y: nextPosition.Y}
			newPaths[left] += counter
			right := Position{X: nextPosition.X + 1, Y: nextPosition.Y}
			newPaths[right] += counter
		}
		paths = newPaths
	}

	result := 0
	for _, counter := range paths {
		result += counter
	}
	return strconv.Itoa(result), nil
}
