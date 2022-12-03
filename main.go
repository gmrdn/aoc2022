package main

import (
	"bytes"
	"fmt"
	"os"

	"aoc2022/d1"
	"aoc2022/d2"
	"aoc2022/d3"
	"aoc2022/days"
)

func main() {
	days := []days.Day{
		d1.NewD1(),
		d2.NewD2(),
		d3.NewD3(),
	}

	for index, day := range days {
		input, err := os.ReadFile(fmt.Sprintf("input_d%d.txt", index+1))
		if err != nil {
			fmt.Println(err)
		}
		day.Input(bytes.NewBuffer(input))
		fmt.Println(day.Run())
	}
}
