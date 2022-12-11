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

func (d *D) Run() int {
	allLines, err := io.ReadAll(d.inputStream)
	if err != nil {
		panic(err)
	}

	g := Game{
		monkeys:           make(map[int]Monkey),
		tranquilizingShot: 1,
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
			nbInspections: 0,
		}
		g.monkeys[monkeyNumber] = monkey
		g.tranquilizingShot = g.tranquilizingShot * monkey.test
	}

	for r := 0; r < 10000; r++ {
		g.RunRound()
	}

	max1 := 0
	max2 := 0
	for _, m := range g.monkeys {
		if m.nbInspections > max1 {
			max2 = max1
			max1 = m.nbInspections
		} else if m.nbInspections > max2 {
			max2 = m.nbInspections
		}
	}

	return max1 * max2
}

func (d *D) RunStr() string {
	return ""
}

type Game struct {
	monkeys           map[int]Monkey
	tranquilizingShot int
}

func (g *Game) RunRound() {
	for i := 0; i < len(g.monkeys); i++ {
		for _, item := range g.monkeys[i].startingItems {
			worryLevel := item
			worryLevel = g.monkeys[i].operation(worryLevel)
			worryLevel = worryLevel % g.tranquilizingShot
			var targetMonkey int
			if worryLevel%g.monkeys[i].test == 0 {
				targetMonkey = g.monkeys[i].ifTrue
			} else {
				targetMonkey = g.monkeys[i].ifFalse
			}
			receiver := g.monkeys[targetMonkey]
			receiver.startingItems = append(receiver.startingItems, worryLevel)
			g.monkeys[targetMonkey] = receiver
			sender := g.monkeys[i]
			sender.startingItems = sender.startingItems[1:]
			sender.nbInspections++
			g.monkeys[i] = sender
		}
	}
}

type Monkey struct {
	startingItems []int
	operation     func(int) int
	test          int
	ifTrue        int
	ifFalse       int
	nbInspections int
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

func parseTest(s string) int {
	divisor, err := strconv.Atoi(s[21:])
	if err != nil {
		panic(err)
	}
	return divisor
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
