package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_instruction_lines(file_name string) ([]string, error) {
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

func parse_instruction(instruction string) (direction string, distance int, err error) {
	splitted_inst := strings.Split(instruction, " ")
	direction = splitted_inst[0]
	distance, err = strconv.Atoi(splitted_inst[1])

	if err != nil {
		return "", 0, err
	}

	return direction, distance, nil
}

func part1(instructions []string) {
	pos_x, pos_y := 0, 0

	for _, instruction := range instructions {
		direction, distance, err := parse_instruction(instruction)

		if err != nil {
			fmt.Printf("Error parsing instruction: %s\n", err.Error())
			return
		}

		switch direction {
		case "forward":
			pos_x += distance
		case "down":
			pos_y += distance
		case "up":
			pos_y -= distance
		default:
			fmt.Printf("Invalid direction: %s", direction)
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d\n", pos_x, pos_y)
	fmt.Printf("Solution (h * d) of Part 1: %d\n", pos_x*pos_y)
}

func part2(instructions []string) {
	pos_x, pos_y := 0, 0
	aim := 0
	for _, instruction := range instructions {
		direction, distance, err := parse_instruction(instruction)

		if err != nil {
			fmt.Printf("Error parsing instruction: %s\n", err.Error())
			return
		}

		switch direction {
		case "forward":
			pos_x += distance
			pos_y += aim * distance
		case "down":
			aim += distance
		case "up":
			aim -= distance
		default:
			fmt.Printf("Invalid direction: %s", direction)
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d\n", pos_x, pos_y)
	fmt.Printf("Solution (h * d) of Part 2: %d\n", pos_x*pos_y)
}

func main() {
	instructions, err := get_instruction_lines("input.txt")

	if err != nil {
		fmt.Printf("Could not read file: %s", err.Error())
		return
	}

	part1(instructions)
	part2(instructions)
}
