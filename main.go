package main

import (
	"bytes"
	"fmt"
)

const (
	// Empty 空き
	Empty = iota

	// Piece 駒
	Piece

	// Control 利き
	Control
)

const (
	north = iota
	northeast
	east
	southeast
	south
	southwest
	west
	northwest
	numOfDir
)

// Pos マスの座標を表す
type Pos struct {
	X int
	Y int
}

// Stringer
func (o Pos) String() string {
	return fmt.Sprintf("(%v,%v)", o.X, o.Y)
}

// Board 盤面表現
type Board [8][8]int

// Stringer
func (p *Board) String() string {
	var buffer bytes.Buffer
	for _, row := range p {
		sep := ""
		for _, v := range row {
			buffer.WriteString(sep)
			sep = " "
			if v == Piece {
				buffer.WriteString("o")
			} else if v == Control {
				buffer.WriteString("x")
			} else {
				buffer.WriteString("-")
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
func (p *Board) piecesPos() (poses []Pos) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			pos := Pos{x, y}
			if p[pos.Y][pos.X] == Piece {
				poses = append(poses, pos)
			}
		}
	}
	return
}
func (p *Board) initialize() {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			pos := Pos{x, y}
			p[pos.Y][pos.X] = Empty
		}
	}
}
func (p *Board) inboard(pos Pos) bool {
	return 0 <= pos.X && pos.X <= 7 && 0 <= pos.Y && pos.Y <= 7
}
func (p *Board) isPutable(pos Pos) bool {
	return p.inboard(pos) && p[pos.Y][pos.X] == Empty
}
func getIncVal(dir int) (int, int) {
	switch dir {
	case north:
		return 0, -1
	case northeast:
		return 1, -1
	case east:
		return 1, 0
	case southeast:
		return 1, 1
	case south:
		return 0, 1
	case southwest:
		return -1, 1
	case west:
		return -1, 0
	case northwest:
		return -1, -1
	default:
		// must not be reach
		return 0, 0
	}
}
func (p *Board) put(center Pos) {
	p[center.Y][center.X] = Piece
	for d := north; d < numOfDir; d++ {
		pos := center
		incX, incY := getIncVal(d)
		pos.X += incX
		pos.Y += incY
		for p.inboard(pos) {
			p[pos.Y][pos.X] = Control
			pos.X += incX
			pos.Y += incY
		}
	}
}

func solve() (complete []Board) {
	var board Board
	board.initialize()
	solveRecursive(board, 0, &complete)
	return
}
func solveRecursive(board Board, currentY int, complete *[]Board) {
	for x := 0; x < 8; x++ {
		if board.isPutable(Pos{x, currentY}) {
			copyBoad := board
			copyBoad.put(Pos{x, currentY})
			if currentY == 7 {
				*complete = append(*complete, copyBoad)
			} else {
				solveRecursive(copyBoad, currentY+1, complete)
			}
		}
	}
}

func main() {
	complete := solve()
	for _, board := range complete {
		fmt.Println(board.piecesPos())
		fmt.Print(&board)
		fmt.Println("###############")
	}
	fmt.Println(len(complete))
}
