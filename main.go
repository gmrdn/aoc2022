package main

import (
	"bytes"
	"fmt"
	"os"

	d1 "aoc2022/d1"
	d2 "aoc2022/d2"
	d3 "aoc2022/d3"
	d4 "aoc2022/d4"
	d5 "aoc2022/d5"
	d6 "aoc2022/d6"
	d7 "aoc2022/d7"
	d8 "aoc2022/d8"
	"aoc2022/days"
)

func main() {
	days := []days.Day{
		d1.NewD(),
		d2.NewD(),
		d3.NewD(),
		d4.NewD(),
		d5.NewD(),
		d6.NewD(),
		d7.NewD(),
		d8.NewD(),
	}

	for index, day := range days {
		input, err := os.ReadFile(fmt.Sprintf("input_d%d.txt", index+1))
		if err != nil {
			fmt.Println(err)
		}
		day.Input(bytes.NewBuffer(input))

		result := day.Run()
		if result == 0 {
			fmt.Printf("Day %d: %s\n", index+1, day.RunStr())
		} else {
			fmt.Printf("Day %d: %d\n", index+1, result)
		}
	}
}
