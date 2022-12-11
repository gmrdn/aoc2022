package d

import (
	"bufio"
	"fmt"
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

func (d *D) Run() int {
	allLines, err := io.ReadAll(d.inputStream)
	if err != nil {
		panic(err)
	}

	g := Game{
		monkeys: make(map[int]Monkey),
	}

	inputMonkeys := strings.Split(string(allLines), "\n\n")
	for _, im := range inputMonkeys {
		lines := strings.Split(im, "\n")
		monkeyNumber := parseMonkeyNumber(lines[0])
		monkey := Monkey{
			startingItems: parseStartingItems(lines[1]),
			operation:     parseOperation(lines[2]),
			test:          parseTest(lines[3]),
			ifTrue:        parseIfTrue(lines[4]),
			ifFalse:       parseIfFalse(lines[5]),
		}
		g.monkeys[monkeyNumber] = monkey

	}

	fmt.Println(g.monkeys)

	return 0
}

func (d *D) RunStr() string {
	return ""
}

type Game struct {
	monkeys map[int]Monkey
}

type Monkey struct {
	startingItems []int
	operation     func(int) int
	test          func(int) bool
	ifTrue        int
	ifFalse       int
}

func parseMonkeyNumber(s string) int {
	nb, err := strconv.Atoi(s[7:8])
	if err != nil {
		panic(err)
	}
	return nb
}

func parseStartingItems(s string) []int {
	items := []int{}
	parsedItems := strings.Split(s[18:], ", ")
	for _, pi := range parsedItems {
		nb, err := strconv.Atoi(pi)
		if err != nil {
			panic(err)
		}
		items = append(items, nb)
	}
	return items
}

func parseOperation(s string) func(int) int {
	formula := s[23:]
	elements := strings.Split(formula, " ")

	if elements[1] == "old" {
		if elements[0] == "*" {
			return func(i int) int {
				return i * i
			}
		}
		if elements[0] == "+" {
			return func(i int) int {
				return i + i
			}
		}
	} else {
		variable, err := strconv.Atoi(elements[1])
		if err != nil {
			panic(err)
		}

		if elements[0] == "*" {
			return func(i int) int {
				return i * variable
			}
		}
		if elements[0] == "+" {
			return func(i int) int {
				return i + variable
			}
		}
	}
	return nil
}

func parseTest(s string) func(int) bool {
	divisor, err := strconv.Atoi(s[21:])
	if err != nil {
		panic(err)
	}
	return func(i int) bool {
		return i%divisor == 0
	}
}

func parseIfTrue(s string) int {
	nb, err := strconv.Atoi(s[29:])
	if err != nil {
		panic(err)
	}
	return nb
}

func parseIfFalse(s string) int {
	nb, err := strconv.Atoi(s[30:])
	if err != nil {
		panic(err)
	}
	return nb
}
