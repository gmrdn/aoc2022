package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

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

	readFile.Close()

	sort.Slice(all, func(i, j int) bool {
		return all[i].Carrying > all[j].Carrying
	})
	fmt.Println("All", all)

	fmt.Println(all[0].Carrying + all[1].Carrying + all[2].Carrying)
}
