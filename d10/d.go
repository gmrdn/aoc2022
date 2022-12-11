package d

import (
	"bufio"
	"io"
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

type op struct {
	opcode  string
	arg     int
	Xbefore int
	Xafter  int
	cycles  []int
}

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	X := 1
	currentCycle := 1
	stackOfOps := []op{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "noop" {
			stackOfOps = append(stackOfOps, op{
				opcode:  "noop",
				arg:     0,
				Xbefore: X,
				Xafter:  X,
				cycles:  []int{currentCycle},
			})
			currentCycle++
		}
		if line[0:4] == "addx" {
			arg, err := strconv.Atoi(line[5:])
			if err != nil {
				panic(err)
			}
			stackOfOps = append(stackOfOps, op{
				opcode:  "addx",
				arg:     arg,
				Xbefore: X,
				Xafter:  X + arg,
				cycles:  []int{currentCycle, currentCycle + 1},
			})

			X += arg
			currentCycle += 2
		}
	}

	sum := 0

	cyclesToCheck := []int{20, 60, 100, 140, 180, 220}
	for _, op := range stackOfOps {
		for _, cycle := range cyclesToCheck {
			for _, opCycle := range op.cycles {
				if cycle == opCycle {
					sum += cycle * op.Xbefore
				}
			}
		}
	}

	return sum
}

func (d *D) RunStr() string {
	return ""
}
