package main

import (
	"fmt"
	"os"

	"aoc2022/d2"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	d2 := d2.NewD2()
	d2.Input(readFile)
	fmt.Println(d2.Run())
}
