package main

import (
	"fmt"
	"os"
	"strconv"
)

// Pos 座標
type Pos struct {
	x int
	y int
}

// Positions 座標の集合
type Positions []Pos

// 利き筋を作る
func getControls(p Pos) Positions {
	var controls Positions
	col := p.x - 1
	for 0 <= col {
		diff := p.x - col
		controls = append(controls,
			Pos{col, p.y},        // 西方向  のマス
			Pos{col, p.y + diff}, // 南西方向のマス
			Pos{col, p.y - diff}) // 北西方向のマス
		col--
	}
	return controls
}

func (q Positions) findPos(p Pos) bool {
	for _, v := range q {
		if v.x == p.x && v.y == p.y {
			// found
			return true
		}
	}
	return false
}
func isSafePos(p Pos, positions Positions) bool {
	// 利き筋を作る
	controls := getControls(p)
	// 利き筋にすでに置いているqueenがあるか?
	for _, q := range positions {
		if controls.findPos(q) {
			return false
		}
	}
	return true
}

func queensRecursive(boardSize, col int, positions Positions, result *[]Positions) {
	if boardSize-1 < col { // col は 0 origin なので.
		// complete!!
		*result = append(*result, positions)
	} else {
		// 新しいqueenを1行ずつ試行.
		for row := 0; row < boardSize; row++ {
			newqueen := Pos{col, row}
			copyPositions := make(Positions, len(positions))
			copy(copyPositions, positions)
			if isSafePos(newqueen, copyPositions) {
				// 置けたら次の列へ
				queensRecursive(boardSize, col+1, append(positions, newqueen), result)
			}
		}
	}
}

func queens(boardSize int) []Positions {
	var positions Positions
	var result []Positions
	queensRecursive(boardSize, 0, positions, &result)
	return result
}
func main() {
	var boardSize int
	if len(os.Args) == 1 {
		boardSize = 8
	} else {
		boardSize = eAtoi(os.Args[1])
	}
	r := queens(boardSize)
	for _, positions := range r {
		fmt.Printf("%v\n", positions)
		// printBoard(boardSize, positions)
		// fmt.Printf("\n")
	}
}

func eAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		os.Exit(1)
	}
	return i
}

func printBoard(boardSize int, positions Positions) {
	board := make([][]bool, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = make([]bool, boardSize)
	}
	for _, pos := range positions {
		board[pos.y][pos.x] = true
	}
	var sep string
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			fmt.Printf(sep)
			sep = " "
			if board[row][col] {
				fmt.Printf("o")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Printf("\n")
		sep = ""
	}
}
