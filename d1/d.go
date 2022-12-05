package d

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
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

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

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

func (d *D) RunStr() string {
	return ""
}
