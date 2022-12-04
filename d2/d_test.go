package d

import (
	"bytes"
	"testing"
)

func TestD2(t *testing.T) {
	input := `A Y
B X
C Z
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 12
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
