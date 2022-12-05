package d

import (
	"bufio"
	"io"
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

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

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
	}

	return score
}

func (d *D) RunStr() string {
	return ""
}
