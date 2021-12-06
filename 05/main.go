package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func get_file(file_name string) ([]string, error) {
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

func parse_commaseparted_coords(coords string) point {
	points := strings.Split(coords, ",")
	x, _ := strconv.Atoi(points[0])
	y, _ := strconv.Atoi(points[1])
	return point{x, y}
}

func parse_points(line string) (p1 point, p2 point) {
	// line == 0,9 -> 5,9
	line = strings.ReplaceAll(line, " ", "")
	substrings := strings.Split(line, "->")
	p1 = parse_commaseparted_coords(substrings[0])
	p2 = parse_commaseparted_coords(substrings[1])
	return p1, p2
}

func draw_line(p1, p2 point, diagram map[point]int) {
	diagram[p1]++
	for p1.x != p2.x || p1.y != p2.y {
		if p1.x > p2.x {
			p1.x--
		} else if p1.x < p2.x {
			p1.x++
		}

		if p1.y > p2.y {
			p1.y--
		} else if p1.y < p2.y {
			p1.y++
		}
		diagram[p1]++
	}
}

func number_of_overlaps(diagram map[point]int) int {
	count := 0
	for _, i := range diagram {
		if i > 1 {
			count++
		}
	}

	return count
}

func part1(file []string) {
	diagram := make(map[point]int)
	for _, line := range file {
		p1, p2 := parse_points(line)

		if p1.x == p2.x || p1.y == p2.y {
			draw_line(p1, p2, diagram)
		}
	}

	fmt.Printf("Solution Part1: %d", number_of_overlaps(diagram))
}

func part2(file []string) {
	diagram := make(map[point]int)
	for _, line := range file {
		p1, p2 := parse_points(line)
		draw_line(p1, p2, diagram)
	}

	fmt.Printf("Solution Part1: %d", number_of_overlaps(diagram))
}

func main() {
	file, _ := get_file("input.txt")
	part1(file)
	part2(file)
}
