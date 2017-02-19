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

	// Control Queenの利き
	Control
)

// 方角
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

const (
	width  = 8
	height = 8
)

// Board 盤面表現
type Board [width][height]int

// 盤面の初期化
func (p *Board) initialize() {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pos := Pos{x, y}
			p[pos.Y][pos.X] = Empty
		}
	}
}

// 盤面の範囲内か?
func (p *Board) isInsidePos(pos Pos) bool {
	return 0 <= pos.X && pos.X < width && 0 <= pos.Y && pos.Y < height
}
func (p *Board) canPut(pos Pos) bool {
	return p.isInsidePos(pos) && p[pos.Y][pos.X] == Empty
}

func getIncrementNum(dir int) (int, int) {
	incrementMap := map[int]struct{ x, y int }{
		north:     {0, -1},
		northeast: {1, -1},
		east:      {1, 0},
		southeast: {1, 1},
		south:     {0, 1},
		southwest: {-1, 1},
		west:      {-1, 0},
		northwest: {-1, -1},
	}
	val := incrementMap[dir]
	return val.x, val.y
}

// 駒を置いて,利きを足す
func (p *Board) putAndAddControl(here Pos) {
	// その場に置く
	p[here.Y][here.X] = Piece

	// 8方位に展開する
	for d := north; d < numOfDir; d++ {
		pos := here
		incX, incY := getIncrementNum(d)
		pos.X += incX
		pos.Y += incY

		// 利きを伸ばす
		for p.isInsidePos(pos) {
			p[pos.Y][pos.X] = Control
			pos.X += incX
			pos.Y += incY
		}
	}
}

func solve() []Board {
	var (
		board     Board
		completed []Board
	)
	board.initialize()
	solveRecursive(board, 0, &completed)
	return completed
}
func solveRecursive(board Board, y int, completed *[]Board) {
	if y == width {
		// もうこれ以上置けないので完成
		*completed = append(*completed, board)
		return
	}
	for x := 0; x < width; x++ {
		if board.canPut(Pos{x, y}) {
			copyBoad := board
			copyBoad.putAndAddControl(Pos{x, y})
			// 次の行を調べる
			solveRecursive(copyBoad, y+1, completed)
		}
	}
	return
}

func main() {
	completed := solve()
	for _, board := range completed {
		fmt.Println(board.collectPosOfPiece())
		fmt.Print(&board)
		fmt.Println("###############")
	}
	fmt.Println(len(completed))
}

/*
 * 主にデバッグ用
 */
// Stringer
func (o Pos) String() string {
	return fmt.Sprintf("(%v,%v)", o.X, o.Y)
}

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

func (p *Board) collectPosOfPiece() (poses []Pos) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pos := Pos{x, y}
			if p[pos.Y][pos.X] == Piece {
				poses = append(poses, pos)
			}
		}
	}
	return
}
