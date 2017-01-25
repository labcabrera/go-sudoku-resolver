package sudoku

import (
	"log"
)


type Resolver struct {
}

func Resolve(board *Board) {
	log.Print("Resolving board\n")
	var success bool
	for {
		success = ResolveIteration(board)
		if(!success) {
			break
		}
	}
	board.Completed = isCompleted(board)
}

func isCompleted(board *Board) bool {
	for i := 0; i < 9; i++ {
		for z := 0; z < 9; z++ {
			if board.Cells[i][z] == 0 {
				return false
			}
		}
	}
	return true
}

func ResolveIteration(board *Board) bool {
	var availableValues []int 
	for row := 0; row < 9; row ++ {
		for col :=0; col < 9; col++ {
			availableValues = AvailableValues(board, row, col)
			//if len(availableValues) == 1 {
			if len(availableValues) == 1 && availableValues[0] > 0 {
				
				log.Print("Setting value [", row, "][", col, "]: ", availableValues[0], "\n")
				//log.Print(board)
				//log.Print(availableValues)
				//PrintBoard(board)
				
				board.Cells[row][col] = availableValues[0]
				return true
			}
		}
	}
	return false
}
func AvailableValues(board *Board, row int, col int) []int{
	var tmp []int
	if board.Cells[row][col] > 0 {
		return tmp
	}
	var v int
	for i := 1; i < 10; i++ {
		tmp = append(tmp, i)
	}
	for i := 0; i < 9; i++ {
		v = board.Cells[row][i]
		if v > 0 {
			tmp[v - 1] = 0
		}
		v = board.Cells[i][col]
		if v > 0 {
			tmp[v - 1] = 0
		}
	}
	var row0 int = 3 * (row / 3)
	var col0 int = 3 * (col / 3)
	for i := row0; i < (row0 + 3); i++ {
		for z := col0; z < (col0 + 3); z++ {
			v = board.Cells[i][z]
			if v > 0 {
				tmp[v - 1] = 0
			}
		}
	}
	var result []int
	for i:= 0; i < 9; i++ {
		if tmp[i] > 0 {
			result = append(result, tmp[i])
		}
	}
	return result
	
	
	
}