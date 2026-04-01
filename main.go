package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"adventofcode2025/internal/challenges"
)

func main() {
	challenge := flag.Int("challenge", 0, "Challenge number to run (1-25)")
	part := flag.Int("part", 0, "Challenge part to run (1-2)")
	inputPath := flag.String("input", "", "Path to the input file")
	flag.Parse()

	if err := run(*challenge, *part, *inputPath); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(challenge int, part int, inputPath string) error {
	if challenge < 1 || challenge > 25 {
		return errors.New("challenge must be between 1 and 25")
	}

	if part < 1 || part > 2 {
		return errors.New("part must be 1 or 2")
	}

	if inputPath == "" {
		return errors.New("input file is required")
	}

	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("read input file: %w", err)
	}

	solver, err := challenges.Get(challenge)
	if err != nil {
		return err
	}

	result, err := solver(part, parseInputLines(inputData))
	if err != nil {
		return fmt.Errorf("run challenge %d part %d: %w", challenge, part, err)
	}

	if err := os.MkdirAll("results", 0o755); err != nil {
		return fmt.Errorf("create results directory: %w", err)
	}

	outputPath := filepath.Join("results", "day_"+strconv.Itoa(challenge)+"_part"+strconv.Itoa(part)+"_"+filepath.Base(inputPath))
	if err := os.WriteFile(outputPath, []byte(result), 0o644); err != nil {
		return fmt.Errorf("write result file: %w", err)
	}

	fmt.Printf("Challenge %d part %d result written to %s\nResult: %s\n", challenge, part, outputPath, result)
	return nil
}

func parseInputLines(inputData []byte) []string {
	normalized := strings.ReplaceAll(string(inputData), "\r\n", "\n")
	normalized = strings.TrimRight(normalized, "\n")

	if normalized == "" {
		return []string{}
	}

	return strings.Split(normalized, "\n")
}
