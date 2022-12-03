package d3

import (
	"bytes"
	"testing"
)

func TestD2(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

	d3 := NewD3()
	d3.Input(bytes.NewBufferString(input))
	got := d3.Run()
	want := 70
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
