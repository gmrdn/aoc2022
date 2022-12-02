package d2

import (
	"bytes"
	"testing"
)

func TestD2(t *testing.T) {
	input := `A Y
B X
C Z
`

	d2 := NewD2()
	d2.Input(bytes.NewBufferString(input))
	got := d2.Run()
	want := 12
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
