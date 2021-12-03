package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	fmt.Printf("Part1 Solution: %d\n", gamma*epsilon)
}

func carbon_filter(most rune, least rune) rune {
	return most
}

func co2_filter(most rune, least rune) rune {
	return least
}

func filter_numbers(lines []string, filter func(rune, rune) rune) int64 {
	remaining_lines := make([]string, len(lines))
	copy(remaining_lines, lines)

	column_index := 0
	for len(remaining_lines) > 1 {
		new_remaining := make([]string, 0, len(remaining_lines))
		target_rune := filter(most_and_least(remaining_lines, column_index))
		for _, line := range remaining_lines {
			if rune(line[column_index]) == target_rune {
				new_remaining = append(new_remaining, line)
			}
		}

		remaining_lines = new_remaining[:]
		column_index++
	}

	rating, _ := strconv.ParseInt(remaining_lines[0], 2, 32)
	fmt.Printf("Rating: %d\n", rating)
	return rating
}

func part2(lines []string) {
	result := filter_numbers(lines, carbon_filter) * filter_numbers(lines, co2_filter)
	fmt.Printf("Part2 Solution: %d\n", result)
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
