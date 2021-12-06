package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_population(file_name string) (population []int, err error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")
	population = make([]int, 9)
	for _, str := range splitted {
		age, _ := strconv.Atoi(str)
		population[age]++
	}

	return population, nil
}

func simulate_population(population []int, days int) (total_fishies int64) {

	for i := 0; i < days; i++ {
		new_fishes := population[0]
		for j := 1; j < len(population); j++ {
			population[j-1] = population[j]
		}
		population[6] = new_fishes + population[6] // sum of fishes with counter < 0 and those that were counted down from 7
		population[8] = new_fishes
	}

	var total int64 = 0
	for _, count := range population {
		total += (int64)(count)
	}

	return total
}

func main() {
	file := "input.txt"
	population, _ := parse_population(file)
	fish_count := simulate_population(population, 80)
	fmt.Printf("Solution Part 1: %d\n", fish_count)

	population, _ = parse_population(file)
	fish_count = simulate_population(population, 256)
	fmt.Printf("Solution Part 2: %d\n", fish_count)
}
