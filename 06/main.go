package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_file(file_name string) (fishies []int, err error) {
	file, err := os.Open(file_name)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	splitted := strings.Split(line, ",")
	fishies = make([]int, 0, len(splitted))
	for _, str := range splitted {
		number, _ := strconv.Atoi(str)
		fishies = append(fishies, number)
	}

	return fishies, nil
}

func simulate_population(fishies []int, days int) (total_fishies int) {
	newFishies := make([]int, 0)
	for i := 0; i < days; i++ {
		for j := 0; j < len(fishies); j++ {
			fishies[j]--

			if fishies[j] == -1 {
				fishies[j] = 6
				newFishies = append(newFishies, 8)
			}
		}

		if len(newFishies) > 0 {
			fishies = append(fishies, newFishies...)
			newFishies = make([]int, 0)
		}
	}

	return len(fishies)
}

func main() {
	input_fishies, _ := parse_file("testinput.txt")
	fishies := input_fishies
	fish_count := simulate_population(fishies, 80)
	fmt.Printf("Solution Part 1: %d\n", fish_count)

	fishies = input_fishies
	fish_count = simulate_population(fishies, 256)
	fmt.Printf("Solution Part 2: %d\n", fish_count)
}
