package d

import (
	"bufio"
	"io"
	"strconv"
	"strings"
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
	arg   int
	X     int
	cycle int
}

func (d *D) Run() int {
	return 0
}

func (d *D) RunStr() string {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	X := 1
	currentCycle := 1
	stackOfOps := []op{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "noop" {
			stackOfOps = append(stackOfOps, op{
				arg:   0,
				X:     X,
				cycle: currentCycle,
			})
			currentCycle++
		}
		if line[0:4] == "addx" {
			arg, err := strconv.Atoi(line[5:])
			if err != nil {
				panic(err)
			}
			stackOfOps = append(stackOfOps, op{
				arg:   arg,
				X:     X,
				cycle: currentCycle,
			})

			X += arg
			currentCycle++

			stackOfOps = append(stackOfOps, op{
				arg:   arg,
				X:     X,
				cycle: currentCycle,
			})
			currentCycle++
		}
	}

	crtStack := []string{}
	crtStack = append(crtStack, "#")

	for i := 1; i < 40; i++ {
		spriteMiddle := stackOfOps[i-1].X
		spriteLeft := spriteMiddle - 1
		spriteRight := spriteMiddle + 1

		if spriteLeft == i || spriteRight == i || spriteMiddle == i {
			crtStack = append(crtStack, "#")
		} else {
			crtStack = append(crtStack, ".")
		}
	}

	for i := 40; i < 80; i++ {
		spriteMiddle := stackOfOps[i-1].X + 40
		spriteLeft := spriteMiddle - 1
		spriteRight := spriteMiddle + 1

		if spriteLeft == i || spriteRight == i || spriteMiddle == i {
			crtStack = append(crtStack, "#")
		} else {
			crtStack = append(crtStack, ".")
		}
	}

	for i := 80; i < 120; i++ {
		spriteMiddle := stackOfOps[i-1].X + 80
		spriteLeft := spriteMiddle - 1
		spriteRight := spriteMiddle + 1

		if spriteLeft == i || spriteRight == i || spriteMiddle == i {
			crtStack = append(crtStack, "#")
		} else {
			crtStack = append(crtStack, ".")
		}
	}

	for i := 120; i < 160; i++ {
		spriteMiddle := stackOfOps[i-1].X + 120
		spriteLeft := spriteMiddle - 1
		spriteRight := spriteMiddle + 1

		if spriteLeft == i || spriteRight == i || spriteMiddle == i {
			crtStack = append(crtStack, "#")
		} else {
			crtStack = append(crtStack, ".")
		}
	}

	for i := 160; i < 200; i++ {
		spriteMiddle := stackOfOps[i-1].X + 160
		spriteLeft := spriteMiddle - 1
		spriteRight := spriteMiddle + 1

		if spriteLeft == i || spriteRight == i || spriteMiddle == i {
			crtStack = append(crtStack, "#")
		} else {
			crtStack = append(crtStack, ".")
		}
	}

	for i := 200; i < 240; i++ {
		spriteMiddle := stackOfOps[i-1].X + 200
		spriteLeft := spriteMiddle - 1
		spriteRight := spriteMiddle + 1

		if spriteLeft == i || spriteRight == i || spriteMiddle == i {
			crtStack = append(crtStack, "#")
		} else {
			crtStack = append(crtStack, ".")
		}
	}

	display := "\n"
	display += strings.Join(crtStack[0:40], "")
	display += "\n"
	display += strings.Join(crtStack[40:80], "")
	display += "\n"
	display += strings.Join(crtStack[80:120], "")
	display += "\n"
	display += strings.Join(crtStack[120:160], "")
	display += "\n"
	display += strings.Join(crtStack[160:200], "")
	display += "\n"
	display += strings.Join(crtStack[200:240], "")
	display += "\n"

	return display
}
