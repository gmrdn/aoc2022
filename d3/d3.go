package d3

import (
	"bufio"
	"io"
)

type D3 struct {
	inputStream io.Reader
}

func NewD3() *D3 {
	return &D3{}
}

func (d3 *D3) Input(input io.Reader) {
	d3.inputStream = bufio.NewReader(input)
}

func (d3 *D3) Run() int {
	fileScanner := bufio.NewScanner(d3.inputStream)

	fileScanner.Split(bufio.ScanLines)

	sumOfPriorities := 0

	for fileScanner.Scan() {
		leftSde := fileScanner.Text()[:len(fileScanner.Text())/2]
		rightSide := fileScanner.Text()[len(fileScanner.Text())/2:]
		existingOnLeft := make(map[byte]bool)

		for i := 0; i < len(leftSde); i++ {
			existingOnLeft[leftSde[i]] = true
		}

		for i := 0; i < len(rightSide); i++ {
			if existingOnLeft[rightSide[i]] {
				if int(rightSide[i]) < 91 { // upper case
					sumOfPriorities += int(rightSide[i]) - 38
					break
				} else {
					sumOfPriorities += int(rightSide[i]) - 96
					break
				}
			}
		}
	}

	return sumOfPriorities
}
