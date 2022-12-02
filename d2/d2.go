package d2

import (
	"bufio"
	"fmt"
	"io"
)

type D2 struct {
	inputStream io.Reader
}

func NewD2() *D2 {
	return &D2{}
}

func (d2 *D2) Input(input io.Reader) {
	d2.inputStream = bufio.NewReader(input)
}

var Score = map[byte]int{
	'A': 1,
	'B': 2,
	'C': 3,
}

var Beats = map[byte]byte{
	'A': 'B',
	'B': 'C',
	'C': 'A',
}

var BeattenBy = map[byte]byte{
	'A': 'C',
	'B': 'A',
	'C': 'B',
}

var Draw = map[byte]byte{
	'A': 'A',
	'B': 'B',
	'C': 'C',
}

func (d2 *D2) Run() int {
	fileScanner := bufio.NewScanner(d2.inputStream)

	fileScanner.Split(bufio.ScanLines)

	score := 0

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()

		opponentsPlay := currentLine[0]
		shouldEndBy := currentLine[2]

		var myPlay byte

		if shouldEndBy == 'X' {
			myPlay = BeattenBy[opponentsPlay]
		}

		if shouldEndBy == 'Y' {
			myPlay = Draw[opponentsPlay]
			score += 3
		}

		if shouldEndBy == 'Z' {
			myPlay = Beats[opponentsPlay]
			score += 6
		}

		score += Score[myPlay]
		fmt.Println(string(myPlay))

	}

	return score
}
