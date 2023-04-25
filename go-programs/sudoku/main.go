package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	boardSize = 9
	boxSize   = 3
	empty     = 0
)

type Board [boardSize][boardSize]int

func main() {
	rand.Seed(time.Now().UnixNano())

	board := generateBoard()
	fmt.Println("Generated Board:")
	printBoard(board)

	solvedBoard := solveBoard(board)
	fmt.Println("Solved Board:")
	printBoard(solvedBoard)
}

func generateBoard() Board {
	board := Board{}
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			board[i][j] = empty
		}
	}

	// Fill the board with random numbers
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board[i][j] == empty {
				for {
					num := rand.Intn(boardSize) + 1
					if isValid(board, i, j, num) {
						board[i][j] = num
						break
					}
				}
			}
		}
	}

	// Remove some numbers to create a puzzle
	emptyCells := boardSize * boardSize / 2
	for emptyCells > 0 {
		i := rand.Intn(boardSize)
		j := rand.Intn(boardSize)
		if board[i][j] != empty {
			emptyCells--
			board[i][j] = empty
		}
	}

	return board
}

func solveBoard(board Board) Board {
	var solve func(Board) bool
	solve = func(board Board) bool {
		row, col := -1, -1
		isEmpty := true

		// Find the next empty cell in the board
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				if board[i][j] == empty {
					row, col = i, j
					isEmpty = false
					break
				}
			}
			if !isEmpty {
				break
			}
		}

		// If all cells are filled, the board is solved
		if isEmpty {
			return true
		}

		// Try to fill the empty cell with a valid number
		for num := 1; num <= boardSize; num++ {
			if isValid(board, row, col, num) {
				board[row][col] = num

				// Recursively solve the updated board
				if solve(board) {
					return true
				}

				board[row][col] = empty
			}
		}

		// Backtrack if no valid number is found
		return false
	}

	// Copy the original board to a new board and solve it
	solvedBoard := Board{}
	copy(solvedBoard[:], board[:])
	solve(solvedBoard)

	return solvedBoard
}

func isValid(board Board, row, col, num int) bool {
	// Check if num is present in the same row
	for j := 0; j < boardSize; j++ {
		if board[row][j] == num {
			return false
		}
	}

	// Check if num is present in the same column
	for i := 0; i < boardSize; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check if num is present in the same box
	boxRow := row - row%boxSize
	boxCol := col - col%boxSize
	for i := boxRow; i < boxRow+boxSize; i++ {
		for j := boxCol; j < boxCol+boxSize; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

func printBoard(board Board) {
	for i := 0; i < boardSize; i++ {
		if i%boxSize == 0 && i != 0 {
			fmt.Println("---------------------")
		}
		for j := 0; j < boardSize; j++ {
			if j%boxSize == 0 && j != 0 {
				fmt.Print("| ")
			}
			fmt.Print(board[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
