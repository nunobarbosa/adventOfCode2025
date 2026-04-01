package challenges

import (
	"fmt"

	"adventofcode2025/internal/challenges/day01"
	"adventofcode2025/internal/challenges/day02"
	"adventofcode2025/internal/challenges/day03"
	"adventofcode2025/internal/challenges/day04"
	"adventofcode2025/internal/challenges/day05"
	"adventofcode2025/internal/challenges/day06"
	"adventofcode2025/internal/challenges/day07"
	"adventofcode2025/internal/challenges/day08"
	"adventofcode2025/internal/challenges/day09"
	"adventofcode2025/internal/challenges/day10"
	"adventofcode2025/internal/challenges/day11"
	"adventofcode2025/internal/challenges/day12"
	"adventofcode2025/internal/challenges/day13"
	"adventofcode2025/internal/challenges/day14"
	"adventofcode2025/internal/challenges/day15"
	"adventofcode2025/internal/challenges/day16"
	"adventofcode2025/internal/challenges/day17"
	"adventofcode2025/internal/challenges/day18"
	"adventofcode2025/internal/challenges/day19"
	"adventofcode2025/internal/challenges/day20"
	"adventofcode2025/internal/challenges/day21"
	"adventofcode2025/internal/challenges/day22"
	"adventofcode2025/internal/challenges/day23"
	"adventofcode2025/internal/challenges/day24"
	"adventofcode2025/internal/challenges/day25"
)

type Solver func(part int, input []string) (string, error)

func Get(day int) (Solver, error) {
	switch day {
	case 1:
		return day01.Solve, nil
	case 2:
		return day02.Solve, nil
	case 3:
		return day03.Solve, nil
	case 4:
		return day04.Solve, nil
	case 5:
		return day05.Solve, nil
	case 6:
		return day06.Solve, nil
	case 7:
		return day07.Solve, nil
	case 8:
		return day08.Solve, nil
	case 9:
		return day09.Solve, nil
	case 10:
		return day10.Solve, nil
	case 11:
		return day11.Solve, nil
	case 12:
		return day12.Solve, nil
	case 13:
		return day13.Solve, nil
	case 14:
		return day14.Solve, nil
	case 15:
		return day15.Solve, nil
	case 16:
		return day16.Solve, nil
	case 17:
		return day17.Solve, nil
	case 18:
		return day18.Solve, nil
	case 19:
		return day19.Solve, nil
	case 20:
		return day20.Solve, nil
	case 21:
		return day21.Solve, nil
	case 22:
		return day22.Solve, nil
	case 23:
		return day23.Solve, nil
	case 24:
		return day24.Solve, nil
	case 25:
		return day25.Solve, nil
	default:
		return nil, fmt.Errorf("challenge %d is not available", day)
	}
}
