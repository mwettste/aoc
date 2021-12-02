package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	direction string
	units     int
}

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

func parse_instruction(instruction string) (cmd command, err error) {
	splitted_inst := strings.Split(instruction, " ")
	cmd.direction = splitted_inst[0]
	cmd.units, err = strconv.Atoi(splitted_inst[1])

	if err != nil {
		return cmd, err
	}

	return cmd, nil
}

func part1(instructions []string) {
	pos_x, pos_y := 0, 0

	for _, instruction := range instructions {
		cmd, err := parse_instruction(instruction)

		if err != nil {
			fmt.Printf("Error parsing instruction: %s\n", err.Error())
			return
		}

		switch cmd.direction {
		case "forward":
			pos_x += cmd.units
		case "down":
			pos_y += cmd.units
		case "up":
			pos_y -= cmd.units
		default:
			fmt.Printf("Invalid direction: %s", cmd.direction)
		}
	}

	fmt.Printf("Horizontal: %d, Depth: %d\n", pos_x, pos_y)
	fmt.Printf("Solution (h * d) of Part 1: %d\n", pos_x*pos_y)
}

func part2(instructions []string) {
	pos_x, pos_y := 0, 0
	aim := 0
	for _, instruction := range instructions {
		cmd, err := parse_instruction(instruction)

		if err != nil {
			fmt.Printf("Error parsing instruction: %s\n", err.Error())
			return
		}

		switch cmd.direction {
		case "forward":
			pos_x += cmd.units
			pos_y += aim * cmd.units
		case "down":
			aim += cmd.units
		case "up":
			aim -= cmd.units
		default:
			fmt.Printf("Invalid direction: %s", cmd.direction)
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
