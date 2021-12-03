package main

import (
	"bufio"
	"fmt"
	"os"
)

func get_lines(file_name string) ([]string, error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	instruction_lines := make([]string, 0, 10)

	for scanner.Scan() {
		instruction_lines = append(instruction_lines, scanner.Text())
	}

	return instruction_lines, nil
}

func part1(lines []string) {
	column_count := len(lines[0])
	line_count := len(lines)
	ones := 0
	zeros := 0
	gamma := 0
	epsilon := 0
	for column := 0; column < column_count; column++ {
		ones = 0
		zeros = 0
		for line := 0; line < line_count; line++ {
			if lines[line][column] == '0' {
				zeros++
			} else {
				ones++
			}
		}
		gamma <<= 1
		epsilon <<= 1
		if ones > zeros {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}

	fmt.Printf("Part1 - Gamma: %d - Epsilon: %d\n", gamma, epsilon)
	fmt.Printf("Part1 Solution: %d", gamma*epsilon)
}

func part2(lines []string) {

}

func main() {
	lines, err := get_lines("input.txt")

	if err != nil {
		fmt.Printf("Could not read file: %s", err.Error())
		return
	}

	part1(lines)
	part2(lines)
}
