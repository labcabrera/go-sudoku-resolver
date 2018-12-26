package sudoku

import (
	"fmt"
	"log"
	"strings"
	"os"
	"bufio"
	"errors"
	"strconv"
)

type Board struct {
	Cells [9][9]int
	AvailableCells int

	StatusRows [9]int
	StatusCells [9]int
	
	Completed bool
}

func LoadFromFile(board *Board, path string) error {
	board.Completed = false
	log.Print("Loading file: ", path, "\n")
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
			log.Print("Readed comment: ", lineReaded, "\n")
			continue
		} else if len(lineReaded) != 9 {
			return errors.New("Invalid line size: " + lineReaded + ". Expected 9 digits")
		} else {
			log.Print("Parsing line: ", lineReaded, "\n")
			for i := 0; i < 9; i++ {
				var value, err = strconv.Atoi(lineReaded[i:i+1])
				if err != nil {
					return err
				}
				board.Cells[count][i] = value
			}
		}
		count++;
	}
	return nil
}

func PrintBoard(board *Board) {
	for i := 0; i < 9; i++ {
		if i > 0 && i % 3 == 0 {
			fmt.Print("\n")
		}
        for j := 0; j < 9; j++ {
			if j > 0 && j % 3 == 0 {
				fmt.Print(" ")
			}
			var x = board.Cells[i][j]
			if(x > 0) {
				fmt.Print(x)
			} else {
				fmt.Print("_")
			}
        }
		fmt.Println()
    }
}
