package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
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

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 45000
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
