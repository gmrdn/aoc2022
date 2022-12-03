package d1

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

type D1 struct {
	inputStream io.Reader
}

func NewD1() *D1 {
	return &D1{}
}

func (d1 *D1) Input(input io.Reader) {
	d1.inputStream = bufio.NewReader(input)
}

type Elf struct {
	Carrying int
}

func NewElf() *Elf {
	return &Elf{}
}

func (e *Elf) AddCarriage(qty int) {
	e.Carrying += qty
}

type Elves []Elf

func (d1 *D1) Run() int {
	fileScanner := bufio.NewScanner(d1.inputStream)

	fileScanner.Split(bufio.ScanLines)
	var all Elves

	currentElf := NewElf()
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		if currentLine == "" {
			all = append(all, *currentElf)

			currentElf = NewElf()
			continue
		}

		qty, err := strconv.Atoi(currentLine)
		if err != nil {
			fmt.Println(err)
		}
		currentElf.AddCarriage(qty)
	}

	sort.Slice(all, func(i, j int) bool {
		return all[i].Carrying > all[j].Carrying
	})

	return all[0].Carrying + all[1].Carrying + all[2].Carrying
}
