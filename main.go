package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	maxCalories := 0

	currentElf := NewElf()
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		if currentLine == "" {
			if currentElf.Carrying > maxCalories {
				maxCalories = currentElf.Carrying
			}
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

	fmt.Println(maxCalories)
}
