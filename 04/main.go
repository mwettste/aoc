package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// i really should clean this up some day...

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

func get_boards(file []string) (boards [][][]int) {
	board := make([][]int, 0)
	for i, line := range file {
		if i < 2 {
			continue
		}

		if line == "" {
			// append board from previous loop and init a new one
			boards = append(boards, board)
			board = make([][]int, 0)
			continue
		}

		string_numbers := strings.Fields(line)
		numbers := make([]int, 0, len(string_numbers))
		for _, str := range string_numbers {
			n, _ := strconv.Atoi(str)
			numbers = append(numbers, n)
		}

		board = append(board, numbers)
	}

	boards = append(boards, board)

	return boards
}

func check_row(row []int) bool {
	for _, number := range row {
		if number != -1 {
			return false
		}
	}

	return true
}

func check_column(column_number int, board [][]int) bool {
	for _, row := range board {
		if row[column_number] != -1 {
			return false
		}
	}

	return true
}

func mark_boards(boards [][][]int, bingo_number int) (has_winner bool, winning_board [][]int, winner_index int) {
	for board_index, board := range boards {
		for i, row := range board {
			for j, number := range row {
				if number == bingo_number {
					board[i][j] = -1

					if (check_row(row) || check_column(j, board)) && !has_winner {
						has_winner = true
						winning_board = board
						winner_index = board_index
					}
				}
			}

		}
	}

	return has_winner, winning_board, winner_index
}

// i'm getting tired and just want to get this over with -.-
func mark_boards2(boards [][][]int, bingo_number int) (has_winner bool, winning_board_indices map[int]bool) {
	winning_board_indices = make(map[int]bool)
	for board_index, board := range boards {
		for i, row := range board {
			for j, number := range row {
				if number == bingo_number {
					board[i][j] = -1

					if check_row(row) || check_column(j, board) {
						has_winner = true
						winning_board_indices[board_index] = true
					}
				}
			}

		}
	}

	return has_winner, winning_board_indices
}

func board_score(board [][]int) int {
	score := 0
	for i, board2 := range board {
		for j := range board2 {
			if board[i][j] != -1 {
				score += board[i][j]
			}
		}
	}

	return score
}

func part1(file []string, bingo_numbers []int) {
	boards := get_boards(file)
	fmt.Printf("Found %d boards\n", len(boards))

	for _, bingo_number := range bingo_numbers {
		if has_winner, winner, _ := mark_boards(boards, bingo_number); has_winner {
			fmt.Printf("Solution Part 1: %d\n", board_score(winner)*bingo_number)
			return
		}
	}
}

func part2(file []string, bingo_numbers []int) {
	boards := get_boards(file)
	fmt.Printf("Found %d boards\n", len(boards))

	for _, bingo_number := range bingo_numbers {
		if has_winner, winner_indices := mark_boards2(boards, bingo_number); has_winner {
			remaining_boards := make([][][]int, 0)
			for i, board := range boards {
				if !winner_indices[i] {
					remaining_boards = append(remaining_boards, board)
				}
			}

			if len(remaining_boards) == 0 {
				fmt.Printf("Solution Part 2: %d\n", board_score(boards[0])*bingo_number)
				return
			}
			boards = remaining_boards
		}
	}
}

func main() {
	file, _ := get_file("input.txt")

	bingo_strings := strings.Split(file[0], ",")
	bingo_numbers := make([]int, 0, len(bingo_strings))
	for _, str := range bingo_strings {
		number, _ := strconv.Atoi(str)
		bingo_numbers = append(bingo_numbers, number)
	}

	part1(file, bingo_numbers)
	part2(file, bingo_numbers)
}
