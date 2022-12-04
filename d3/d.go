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

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	sumOfPriorities := 0

	elfNb := 0
	foundInFirstElf := make(map[byte]bool)
	foundInSecondElf := make(map[byte]bool)

	for fileScanner.Scan() {
		rumsack := fileScanner.Text()

		if elfNb == 0 || elfNb == 3 {
			elfNb = 0
			foundInFirstElf = make(map[byte]bool)
			foundInSecondElf = make(map[byte]bool)
			for i := 0; i < len(rumsack); i++ {
				foundInFirstElf[rumsack[i]] = true
			}
			elfNb++
			continue
		}

		if elfNb == 1 {
			for i := 0; i < len(rumsack); i++ {
				if foundInFirstElf[rumsack[i]] {
					foundInSecondElf[rumsack[i]] = true
				}
			}
			elfNb++
			continue
		}

		if elfNb == 2 {
			for i := 0; i < len(rumsack); i++ {
				if foundInSecondElf[rumsack[i]] {
					if int(rumsack[i]) < 91 { // upper case
						sumOfPriorities += int(rumsack[i]) - 38
						break
					} else {
						sumOfPriorities += int(rumsack[i]) - 96
						break
					}
				}
			}
			elfNb++
			continue
		}
	}

	return sumOfPriorities
}
