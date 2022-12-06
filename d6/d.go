package d

import (
	"bufio"
	"fmt"
	"io"
)

type D struct {
	inputStream io.Reader
}

func NewD() *D {
	return &D{}
}

func (d *D) Input(input io.Reader) {
	d.inputStream = bufio.NewReader(input)
}

func (d *D) Run() int {
	runes, err := io.ReadAll(d.inputStream)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < len(runes); i++ {
		seq := make(map[byte]bool)
		for j := i; j < i+14; j++ {
			ok := seq[runes[j]]
			if ok {
				break
			}
			seq[runes[j]] = true
		}

		if len(seq) == 14 {
			return i + 14
		}

	}

	return 0
}

type stacks map[int][]string

func (d *D) RunStr() string {
	return ""
}
