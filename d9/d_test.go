package d

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 36
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
