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
	want := 31
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
