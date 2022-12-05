package d

import (
	"bytes"
	"testing"
)

func TestD(t *testing.T) {
	input := `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.RunStr()
	want := "MCD"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
