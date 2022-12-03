package d1

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000

`

	d1 := NewD1()
	d1.Input(bytes.NewBufferString(input))
	got := d1.Run()
	want := 45000
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
