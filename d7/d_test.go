package d

import (
	"bytes"
	"testing"
)

func TestD1(t *testing.T) {
	input := `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

	d := NewD()
	d.Input(bytes.NewBufferString(input))
	got := d.Run()
	want := 95437
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
