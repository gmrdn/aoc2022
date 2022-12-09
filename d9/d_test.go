package d

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 13
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `U 4
R 4
D 4
L 4
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 13
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
