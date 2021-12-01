package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getDepthsFromFile(fileName string) ([]int, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	depths := make([]int, 0, 2000)
	for scanner.Scan() {
		currentDepth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		depths = append(depths, currentDepth)
	}

	return depths, nil
}

func part1(depths []int) {
	increase_counter := 0
	last_depth := math.MaxInt32

	for _, depth := range depths {
		if depth > last_depth {
			increase_counter++
		}

		last_depth = depth
	}

	fmt.Printf("Depth increases %d times (Part 1)\n", increase_counter)
}

func sum_of_slice(slice []int) int {
	sum := 0
	for _, value := range slice {
		sum += value
	}
	return sum
}

func part2(depths []int) {
	window_len := 3
	last_sum := math.MaxInt32
	increase_counter := 0

	for i := 0; i <= len(depths)-window_len; i++ {
		window := depths[i : i+window_len]
		current_sum := sum_of_slice(window)
		if current_sum > last_sum {
			increase_counter++
		}

		last_sum = current_sum
	}

	fmt.Printf("Depth increases %d times (Part 2)\n", increase_counter)
}

func main() {
	depths, err := getDepthsFromFile("input.txt")

	if err != nil {
		fmt.Printf("Could not parse file: %s", err.Error())
	}

	part1(depths)
	part2(depths)
}
