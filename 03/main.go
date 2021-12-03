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

func most_and_least(lines []string, index int) (most rune, least rune) {
	zeros, ones := 0, 0
	for _, line := range lines {
		if line[index] == '0' {
			zeros++
		} else {
			ones++
		}
	}

	if zeros > ones {
		return '0', '1'
	} else {
		return '1', '0'
	}
}

func part1(lines []string) {
	gamma := 0
	epsilon := 0

	for column_index := range lines[0] {
		most, least := most_and_least(lines, column_index)

		gamma <<= 1
		epsilon <<= 1

		if most == '1' {
			gamma |= 1
		}

		if least == '1' {
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
