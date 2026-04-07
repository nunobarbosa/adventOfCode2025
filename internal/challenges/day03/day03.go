package day03

import (
	"strconv"
)

func BuildInput(input []string) ([][]int64, error) {
	banks := make([][]int64, len(input))

	for _, rawBank := range input {
		newBank := []int64{}
		for _, rawBattery := range rawBank {
			newBank = append(newBank, int64(rawBattery-'0'))
		}
		banks = append(banks, newBank)
	}

	return banks, nil
}

func Solve(part int, input []string) (string, error) {
	banks, _ := BuildInput(input)

	switch part {
	case 1:
		return solvePart1(banks)
	case 2:
		return solvePart2(banks)
	default:
		return "", nil
	}
}

func solvePart1(banks [][]int64) (string, error) {
	var maxJoltage int64 = 0

	for _, bank := range banks {
		maxJoltage += maxJoltageForBattery(bank, 2)
	}

	return strconv.FormatInt(maxJoltage, 10), nil
}

func solvePart2(banks [][]int64) (string, error) {
	var maxJoltage int64 = 0

	for _, bank := range banks {
		maxJoltage += maxJoltageForBattery(bank, 12)
	}

	return strconv.FormatInt(maxJoltage, 10), nil
}

func maxJoltageForBattery(batteries []int64, fullSize int) int64 {
	if len(batteries) < fullSize {
		return 0
	}

	minWindow := 0
	maxWindow := len(batteries) - fullSize
	var bankJoltage int64 = 0
	for range fullSize {
		if minWindow > len(batteries) {
			break
		}

		var currentJoltage int64 = batteries[minWindow]
		for selectedBattery := minWindow + 1; selectedBattery < maxWindow+1; selectedBattery++ {
			if currentJoltage < batteries[selectedBattery] {
				minWindow = selectedBattery
				currentJoltage = batteries[selectedBattery]
			}
		}
		bankJoltage *= 10
		bankJoltage += currentJoltage
		minWindow++
		maxWindow++
	}
	return bankJoltage
}
