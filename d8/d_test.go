package d

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `30373
25512
65332
33549
35390
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 8
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
