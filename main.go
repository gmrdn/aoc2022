package main

import (
	"bytes"
	"fmt"
	"os"

	d1 "aoc2022/d1"
	d2 "aoc2022/d2"
	d3 "aoc2022/d3"
	"aoc2022/days"
)

func main() {
	days := []days.Day{
		d1.NewD(),
		d2.NewD(),
		d3.NewD(),
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
