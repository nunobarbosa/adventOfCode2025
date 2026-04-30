package day10

import (
	"fmt"
	"strconv"
	"strings"
)

type Machine struct {
	Expected  uint
	Sequences []uint
	Joltages  []uint
}

type State struct {
	Value uint
	Steps uint
}

func Solve(part int, input []string) (string, error) {
	machines, err := BuildInput(input)
	if err != nil {
		return "", nil
	}

	switch part {
	case 1:
		return solvePart1(machines)
	case 2:
		return "day 10 part 2 not implemented\n", nil
	default:
		return "", nil
	}
}

func BuildInput(input []string) ([]Machine, error) {
	machines := []Machine{}
	for _, l := range input {
		rawValues := strings.Split(l, " ")
		machine := Machine{
			Expected: 0,
		}
		for _, r := range rawValues {
			if strings.HasPrefix(r, "[") {
				// Expected
				lights := r[1 : len(r)-1]
				for p, c := range lights {
					if c != '.' {
						machine.Expected |= 1 << p
					}
				}
				continue
			}

			if strings.HasPrefix(r, "{") {
				joltages := strings.Split(r[1:len(r)-1], ",")
				machine.Joltages = make([]uint, 0, len(joltages))
				for _, j := range joltages {
					joltage, err := strconv.Atoi(j)
					if err != nil {
						return nil, err
					}
					machine.Joltages = append(machine.Joltages, uint(joltage))
				}
				continue
			}

			lights := strings.Split(r[1:len(r)-1], ",")
			var sequence uint = 0
			for _, j := range lights {
				s, err := strconv.Atoi(j)
				if err != nil {
					return nil, err
				}
				sequence |= 1 << s
			}
			machine.Sequences = append(machine.Sequences, sequence)
		}
		machines = append(machines, machine)
	}

	return machines, nil
}

func solvePart1(machines []Machine) (string, error) {
	var checkSum uint = 0

	for _, m := range machines {
		combinations, found := m.Find()
		if !found {
			return "", fmt.Errorf("machine does not have a solution")
		}
		fmt.Printf("Combinations: %d\n", combinations)
		checkSum += combinations
	}

	return strconv.Itoa(int(checkSum)), nil
}

func (m *Machine) Find() (uint, bool) {
	queue := []State{{Value: 0, Steps: 0}}
	visited := map[uint]bool{
		0: true,
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Value == m.Expected {
			return current.Steps, true
		}

		for _, sequence := range m.Sequences {
			nextValue := current.Value ^ sequence

			if visited[nextValue] {
				continue
			}

			visited[nextValue] = true
			queue = append(queue, State{
				Value: nextValue,
				Steps: current.Steps + 1,
			})
		}
	}

	return 0, false
}
