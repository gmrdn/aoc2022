package d

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`
	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 29
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestParsing(t *testing.T) {
	input := `Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	mat, startingPositions, dest := d.ParseInput()
	want := [][]byte{
		{97, 97, 98, 113, 112, 111, 110, 109},
		{97, 98, 99, 114, 121, 120, 120, 108},
		{97, 99, 99, 115, 122, 122, 120, 107},
		{97, 99, 99, 116, 117, 118, 119, 106},
		{97, 98, 100, 101, 102, 103, 104, 105},
	}

	if len(mat) != len(want) {
		t.Errorf("got %v want %v", mat, want)
	}

	for i, row := range mat {
		for j, col := range row {
			if col != want[i][j] {
				t.Errorf("got %v want %v", col, want[i][j])
			}
		}
	}

	expectedStartingPositions := []position{
		{0, 0},
		{1, 0},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
	}
	for i, pos := range startingPositions {
		if pos != expectedStartingPositions[i] {
			t.Errorf("got %v want %v", pos, expectedStartingPositions[i])
		}
	}

	expectedDest := position{5, 2}
	if dest != expectedDest {
		t.Errorf("got %v want %v", dest, expectedDest)
	}
}
