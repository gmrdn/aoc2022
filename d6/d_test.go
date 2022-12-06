package d

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `bvwbjplbgvbhsrlpgdmjqwftvncz`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 5
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD2(t *testing.T) {
	input := `nppdvjthqldpwncqszvftbrmjlhg`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 6
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD3(t *testing.T) {
	input := `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 10
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestD4(t *testing.T) {
	input := `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 11
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
