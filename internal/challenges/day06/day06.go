package day06

import (
	"fmt"
	"strconv"
	"strings"
)

type OperationStruct struct {
	Operation string
	Values    []int64
}

func (o *OperationStruct) AddValue(l string) error {
	cleanLine := strings.Join(strings.Fields(strings.TrimSpace(l)), " ")
	v, err := strconv.Atoi(cleanLine)
	if err != nil {
		return err
	}
	o.Values = append(o.Values, int64(v))
	return nil
}

func (o *OperationStruct) GetResult() int64 {
	result := o.Values[0]
	for i := 1; i < len(o.Values); i++ {
		if o.Operation == "*" {
			result *= o.Values[i]
			continue
		}
		result += o.Values[i]
	}
	return result
}

func BuildInputPart1(input []string) ([]int, []string, error) {
	inputValues := []int{}
	inputOperations := []string{}

	for _, l := range input {
		cleanLine := strings.Join(strings.Fields(strings.TrimSpace(l)), " ")
		values := strings.Split(cleanLine, " ")

		if values[0] == "+" || values[0] == "*" {
			inputOperations = values
			break
		}

		for _, v := range values {
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, nil, err
			}
			inputValues = append(inputValues, n)
		}
	}

	return inputValues, inputOperations, nil
}

func BuildInputPart2(input []string) ([]OperationStruct, error) {
	// Getting operations signs
	rawSigns := input[len(input)-1]
	cleanLine := strings.Join(strings.Fields(strings.TrimSpace(rawSigns)), " ")
	opSigns := strings.Split(cleanLine, " ")

	input = input[:len(input)-1]
	ops := []OperationStruct{{Operation: "", Values: []int64{}}}
	currentOps := 0

	for i := 0; i < len(input[0]); i++ {
		var sb strings.Builder
		for l := 0; l < len(input); l++ {
			singleChar := string(input[l][i])
			sb.WriteString(singleChar)
		}

		line := sb.String()
		if len(strings.TrimSpace(line)) == 0 {
			currentOps++
			ops = append(ops, OperationStruct{Operation: "", Values: []int64{}})
			continue
		}
		err := ops[currentOps].AddValue(sb.String())
		if err != nil {
			return nil, err
		}
		ops[currentOps].Operation = opSigns[currentOps]
	}

	fmt.Printf("%v\n", ops)

	return ops, nil
}

func Solve(part int, input []string) (string, error) {
	switch part {
	case 1:
		values, operations, err := BuildInputPart1(input)
		if err != nil {
			return "", err
		}
		return solvePart1(values, operations)
	case 2:
		ops, err := BuildInputPart2(input)
		if err != nil {
			return "", err
		}
		return solvePart2(ops)
	default:
		return "", nil
	}
}

func solvePart1(values []int, operations []string) (string, error) {
	jump := len(operations)

	var grandTotal int64 = 0
	for v, o := range operations {
		total := values[v]
		for i := v + jump; i < len(values); i += jump {
			currentValue := values[i]
			switch o {
			case "*":
				total *= currentValue
			case "+":
				total += currentValue
			}
		}
		grandTotal += int64(total)
	}

	return strconv.FormatInt(grandTotal, 10), nil
}

func solvePart2(ops []OperationStruct) (string, error) {
	var result int64 = 0

	for _, o := range ops {
		result += o.GetResult()
	}

	return strconv.FormatInt(result, 10), nil
}
