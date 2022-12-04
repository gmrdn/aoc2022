package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 70
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
