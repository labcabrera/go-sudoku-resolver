package main

import (
	"fmt"
	"log"
	"strings"
	"os"
	"bufio"
	"errors"
	"strconv"
)

func main() {
	fmt.Printf("Sudoku demo resolver\n")	
	var board[100] int
	var err = readBoard(&board, os.Args[1])
	if err != nil {
		log.Fatalf("Error reading board: %s", err)
	}
	
	fmt.Print("<...>\n")
	printBoard(board)
	eval(&board)
	
	fmt.Printf("<...After eval...>\n")
	printBoard(board)
	
	tryResolveIteration(&board)
	fmt.Printf("<...After first iteration...>\n")
	eval(&board)
	printBoard(board)
	
	tryResolve(&board)
	fmt.Printf("<...After try resolve...>\n")
	eval(&board)
	printBoard(board)
	
	
	var availableValues [10]int 
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			availableValues = evalCell(&board, row, col)
			printCellAvailableValues(availableValues)
			fmt.Print("   ")
		}
		fmt.Print("\n")
	}
}

func writePosition(board *[100] int, row int, col int, value int) {
	board[row * 10 + col] = value
}

func readPosition(board [100] int, row int, col int) int {
	return board[row * 10 + col]
}

func inc(board *[100] int, row int, col int, val int) {
	board[row * 10 + col] = board[row * 10 + col] + val
}

func readBoard(board *[100] int, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineReaded string
	var count = 0
	for scanner.Scan() {
		lineReaded = strings.TrimSpace(scanner.Text())
		if count > 9 {
			return errors.New("Invalid line count")
		} else if strings.HasPrefix(lineReaded, "#") {
			// Ignore comments on txt file
			log.Print("Readed comment: ", lineReaded, "\n")
			continue
		} else if len(lineReaded) != 9 {
			return errors.New("Invalid line size: " + lineReaded + ". Expected 9 digits")
		} else {
			log.Print("Parsing line: ", lineReaded, "\n")
			for i := 0; i < 9; i++ {
				var ri, err = strconv.Atoi(lineReaded[i:i+1])
				if err != nil {
					return err
				}
				writePosition(board, count, i, ri)
			}
		}
		count++;
	}
	return nil
}

func tryResolve(board *[100] int) {
	var success bool
	for {
		success = tryResolveIteration(board)
		if(!success) {
			break
		}
	}
}

func tryResolveIteration(board *[100] int) bool {
	var availableValues [10]int 
	for row := 0; row < 9; row ++ {
		for col :=0; col < 9; col++ {
			availableValues = evalCell(board, row, col)
			if availableValues[0] == 1 {
				for i := 1; i < 10; i++ {
					if(availableValues[i] > 0) {
						fmt.Print("Setting cell value: ", row, col, availableValues[i], "\n")
						writePosition(board, row, col, availableValues[i])
						return true
					}
				}
			}
		}
	}
	return false
}

func printBoard(board[100] int) {
	for i := 0; i < 10; i++ {
		if i > 0 && i % 3 == 0 {
			fmt.Print("\n")
		}
        for j := 0; j < 10; j++ {
			if j > 0 && j % 3 == 0 {
				fmt.Print(" ")
			}
			var x = readPosition(board, i, j)
			if(x > 0) {
				fmt.Print(x)
			} else {
				fmt.Print("_")
			}
        }
		fmt.Println()
    }
	fmt.Print("  Empty cells: ", readPosition(board, 9, 9), "\n")
}

func printCellAvailableValues(val [10]int) {
	for i := 1; i < 10; i++ {
		if val[i] > 0 {
			fmt.Print(val[i])
		} else {
			fmt.Print("_")
		}
	}
	fmt.Print("(", val[0], ")")
}

func eval(board *[100] int) {
	for i := 0; i < 10; i++ {
		writePosition(board, i, 9, 9)
		writePosition(board, 9, i, 9)
	}
	writePosition(board, 9, 9, 81)
	count := 0
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if readPosition(*board, row, col) > 0 {
				inc(board, row, 9, -1)
				inc(board, 9, col, -1)
				count++
			}
		}
	}
	writePosition(board, 9, 9, 9 * 9 - count)
}

func resolve(board*[100] int) {
}

func evalCell(board*[100] int, row int, col int) [10]int {
	if readPosition(*board, row, col) != 0 {
		return [10]int{0,0,0,0,0,0,0,0,0,0}
	}
	res := [10]int{0,1,2,3,4,5,6,7,8,9}
	for i := 0; i < 9; i++ {
		res[readPosition(*board, row, i)] = 0
		res[readPosition(*board, i, col)] = 0
	}
	var row0 int = 3 * (row / 3)
	var col0 int = 3 * (col / 3)
	for i := row0; i < (row0 + 3); i++ {
		for z := col0; z < (col0 + 3); z++ {
			res[readPosition(*board, i, z)] = 0
		}
	}
	var count = 0
	for i := 1 ; i<10; i++ {
		if(res[i] > 0) {
			count = count + 1
		}
	}
	res[0] = count
	return res
}