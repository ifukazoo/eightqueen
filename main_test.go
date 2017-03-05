package main

import (
	"reflect"
	"testing"
)

type testIsSafePosInput struct {
	p         Pos
	positions Positions
}

func Test_isSafePos(t *testing.T) {
	tests := []struct {
		input testIsSafePosInput
		want  bool
	}{
		{testIsSafePosInput{Pos{4, 5}, Positions{{3, 2}, {2, 0}, {1, 3}, {0, 1}}}, false},
		{testIsSafePosInput{Pos{4, 4}, Positions{{3, 2}, {2, 0}, {1, 3}, {0, 1}}}, true},
	}
	for _, test := range tests {
		if got := isSafePos(test.input.p, test.input.positions); got != test.want {
			t.Errorf("test(%v) return[%v], want[%v]", test.input, got, test.want)
		}
	}
}
func Test_getControls(t *testing.T) {
	tests := []struct {
		input Pos
		want  Positions
	}{
		{Pos{0, 0}, Positions{}},
		{Pos{1, 1}, Positions{{0, 1}, {0, 2}, {0, 0}}},
		{Pos{2, 2}, Positions{{1, 2}, {1, 3}, {1, 1}, {0, 2}, {0, 4}, {0, 0}}},
	}
	for _, test := range tests {
		if got := getControls(test.input); !reflect.DeepEqual(got, test.want) {
			t.Errorf("test(%v) return[%v], want[%v]", test.input, got, test.want)
		}
	}
}
