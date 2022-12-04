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
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	nbOverlapping := 0

	for fileScanner.Scan() {
		pairs := strings.Split(fileScanner.Text(), ",")
		leftSide := strings.Split(pairs[0], "-")
		left1, _ := strconv.Atoi(leftSide[0])
		right1, _ := strconv.Atoi(leftSide[1])

		rightSide := strings.Split(pairs[1], "-")
		left2, _ := strconv.Atoi(rightSide[0])
		right2, _ := strconv.Atoi(rightSide[1])

		// ..l1........r1..
		// .....l2....r2...
		if left1 <= left2 && right1 >= right2 {
			nbOverlapping++
			continue
		}
		// ...l1...r1....
		// ..l2......r2..
		if left1 >= left2 && right1 <= right2 {
			nbOverlapping++
			continue
		}

	}

	return nbOverlapping
}
