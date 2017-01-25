package main

import (
	"fmt"
	"github.com/labcabrera/go-sudoku-resolver/sudoku"
)

func main() {
	fmt.Print("It works\n")
	board := new(sudoku.Board)
	sudoku.LoadFromFile(board, "./src/github.com/labcabrera/go-sudoku-resolver/resources/sudoku.txt")
	
	fmt.Print("Loaded board:\n")
	sudoku.PrintBoard(board)
	
	sudoku.Resolve(board)
	
	fmt.Print("Resolved board:\n")
	sudoku.PrintBoard(board)
	
	if !board.Completed {	
		printAvailableValues(board)
	}
}

func printAvailableValues(board *sudoku.Board) {
	for i := 0 ; i < 9 ; i++ {
		for z := 0 ; z < 9 ; z++ {
			fmt.Print("cell[", i, "][", z, "]=", sudoku.AvailableValues(board,i,z), "\n")
		}
	}
}
