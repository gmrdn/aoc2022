package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 4
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
