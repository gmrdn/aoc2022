package main

import (
	"fmt"
	"os"

	"aoc2022/d1"
	"aoc2022/d2"
)

func main() {
	runD1()
	runD2()
}

func runD1() {
	readFile, err := os.Open("input_d1.txt")
	if err != nil {
		fmt.Println(err)
	}

	d1 := d1.NewD1()
	d1.Input(readFile)
	fmt.Printf("Day 1: The top three Elves carrying the most Calories are carrying %d Calories\n", d1.Run())
}

func runD2() {
	readFile, err := os.Open("input_d2.txt")
	if err != nil {
		fmt.Println(err)
	}

	d2 := d2.NewD2()
	d2.Input(readFile)
	fmt.Printf("Day 2: My total score if everything goes exactly according to your strategy guide would be %d points\n", d2.Run())
}
