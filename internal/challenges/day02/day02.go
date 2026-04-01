package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func BuildInput(input []string) ([][]string, error) {
	intervals := [][]string{}

	// expecting only one line
	if len(input) != 1 {
		return nil, fmt.Errorf("Expecting exactly one line, got %d", len(input))
	}
	rawIntervals := strings.Split(input[0], ",")

	for _, rawInterval := range rawIntervals {
		intervals = append(intervals, strings.Split(rawInterval, "-"))
	}

	return intervals, nil
}

func Solve(part int, input []string) (string, error) {
	parsedInput, err := BuildInput(input)
	if err != nil {
		return "", err
	}

	switch part {
	case 1:
		return solvePart1(parsedInput)
	case 2:
		return solvePart2(parsedInput)
	default:
		return "", nil
	}
}

func solvePart1(intervals [][]string) (string, error) {
	invalidIDsSum := 0

	for _, rawInterval := range intervals {
		start, err := strconv.Atoi(rawInterval[0])
		if err != nil {
			return "", fmt.Errorf("Invalid start %s", rawInterval[0])
		}
		end, err := strconv.Atoi(rawInterval[1])
		if err != nil {
			return "", fmt.Errorf("Invalid end %s", rawInterval[1])
		}

		for i := start; i <= end; i++ {
			current := strconv.Itoa(i)

			// patterns will only happen in even lenght values
			currentLength := len(current)
			if currentLength%2 != 0 {
				continue
			}

			splitAt := currentLength / 2
			firstHalf := current[:splitAt]
			secondHalf := current[splitAt:]

			if firstHalf == secondHalf {
				invalidIDsSum += i
			}
		}
	}

	return strconv.Itoa(invalidIDsSum), nil
}

func solvePart2(intervals [][]string) (string, error) {
	invalidIDsSum := 0

	for _, rawInterval := range intervals {
		start, err := strconv.Atoi(rawInterval[0])
		if err != nil {
			return "", fmt.Errorf("Invalid start %s", rawInterval[0])
		}
		end, err := strconv.Atoi(rawInterval[1])
		if err != nil {
			return "", fmt.Errorf("Invalid end %s", rawInterval[1])
		}

		for i := start; i <= end; i++ {
			current := strconv.Itoa(i)

			for size := len(current) / 2; size > 0; size-- {
				if isInvalidIDForPart2(current, size) {
					invalidIDsSum += i
					break
				}
			}
		}
	}

	return strconv.Itoa(invalidIDsSum), nil
}

func isInvalidIDForPart2(value string, size int) bool {
	if size < 1 || size > len(value) {
		return false
	}

	base := value[:size]
	remain := value

	for len(remain) >= size {
		secondBase := remain[:size]

		if secondBase != base {
			return false
		}

		remain = remain[size:]
	}

	return len(remain) == 0
}
