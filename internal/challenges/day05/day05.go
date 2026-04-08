package day05

import (
	"fmt"
	"strconv"
	"strings"
)

func BuildInput(input []string) ([]int, []int, error) {
	ranges := []int{}
	ingredients := []int{}

	checkingRanges := true
	for _, line := range input {
		if line == "" {
			checkingRanges = false
			continue
		}

		if checkingRanges {
			currentRange := strings.Split(line, "-")
			min, err := strconv.Atoi(currentRange[0])
			if err != nil {
				return ranges, ingredients, err
			}
			max, err := strconv.Atoi(currentRange[1])
			if err != nil {
				return ranges, ingredients, err
			}
			ranges = append(ranges, min, max)

		} else {
			ingredient, err := strconv.Atoi(line)
			if err != nil {
				return ranges, ingredients, err
			}
			ingredients = append(ingredients, ingredient)
		}
	}

	return ranges, ingredients, nil
}

func Solve(part int, input []string) (string, error) {
	ranges, ingredients, err := BuildInput(input)
	if err != nil {
		return "", err
	}
	fmt.Printf("Ranges: %v\n", ranges)
	fmt.Printf("Ingredients: %v\n", ingredients)
	switch part {
	case 1:
		return solvePart1(ranges, ingredients)
	case 2:
		return solvePart2(ranges)
	default:
		return "", nil
	}
}

func solvePart1(ranges []int, ingredients []int) (string, error) {
	fresh := 0
	for _, i := range ingredients {
		for r := 0; r < len(ranges); r += 2 {
			isFresh := i >= ranges[r] && i <= ranges[r+1]

			if isFresh {
				fresh++
				break
			}
		}
	}

	return strconv.Itoa(fresh), nil
}

func solvePart2(ranges []int) (string, error) {
	fresh := 0
	merged := normalizeRanges(ranges)

	for i := 0; i < len(merged); i += 2 {
		fresh += merged[i+1] - merged[i] + 1
	}

	return strconv.Itoa(fresh), nil
}

func normalizeRanges(ranges []int) []int {
	if len(ranges) == 0 {
		return nil
	}

	sorted := make([]int, len(ranges))
	copy(sorted, ranges)

	for i := 0; i < len(sorted)-2; i += 2 {
		minIndex := i
		for j := i + 2; j < len(sorted); j += 2 {
			if sorted[j] < sorted[minIndex] || (sorted[j] == sorted[minIndex] && sorted[j+1] < sorted[minIndex+1]) {
				minIndex = j
			}
		}
		if minIndex != i {
			sorted[i], sorted[minIndex] = sorted[minIndex], sorted[i]
			sorted[i+1], sorted[minIndex+1] = sorted[minIndex+1], sorted[i+1]
		}
	}

	newRanges := []int{sorted[0], sorted[1]}

	for i := 2; i < len(sorted); i += 2 {
		currentMin := sorted[i]
		currentMax := sorted[i+1]

		lastMaxIndex := len(newRanges) - 1
		if currentMin <= newRanges[lastMaxIndex]+1 {
			if currentMax > newRanges[lastMaxIndex] {
				newRanges[lastMaxIndex] = currentMax
			}
			continue
		}

		newRanges = append(newRanges, currentMin, currentMax)
	}

	return newRanges
}
