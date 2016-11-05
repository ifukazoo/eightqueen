package main

import (
	"fmt"
	"testing"
)

func TestStringer(t *testing.T) {
	var board Board
	board.initialize()
	board.put(Pos{3, 4})
	fmt.Println(&board)
	board.put(Pos{1, 0})
	fmt.Println(&board)
	fmt.Println("")
	board.put(Pos{7, 2})
	fmt.Println(&board)
}
func TestPutable(t *testing.T) {
	var board Board
	board.initialize()
	if board.IsPutable(Pos{8, 0}) {
		t.Error(`board.IsPutable(Pos{8,0}) == false`)
	}
	if board.IsPutable(Pos{0, 8}) {
		t.Error(`board.IsPutable(Pos{0,8}) == false`)
	}

	board.put(Pos{0, 0})
	if board.IsPutable(Pos{0, 0}) {
		t.Error(`board.IsPutable(Pos{0,0}) == true`)
	}
	if board.IsPutable(Pos{1, 0}) {
		t.Error(`board.IsPutable(Pos{1,0}) == true`)
	}
	if board.IsPutable(Pos{0, 1}) {
		t.Error(`board.IsPutable(Pos{0,1}) == true`)
	}
	if !board.IsPutable(Pos{2, 1}) {
		t.Error(`board.IsPutable(Pos{2,1}) == false`)
	}
	if !board.IsPutable(Pos{1, 2}) {
		t.Error(`board.IsPutable(Pos{1,2}) == false`)
	}
}

func TestMain(t *testing.T) {
	complete := solve()
	if len(complete) != 92 {
		t.Error(`len(complete) != 92`)
	}
}
