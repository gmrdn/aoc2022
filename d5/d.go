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
	return 0
}

type stacks map[int][]string

func (d *D) RunStr() string {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	s := make(stacks)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, " 1   2   3 ") {
			break
		}

		for i := 1; i < len(line)+2; i++ {
			if i%4 == 0 {
				if s[i/4] == nil {
					s[i/4] = make([]string, 0)
				}

				value := string(line[i-3])
				if value == " " {
					continue
				}

				s[i/4] = append([]string{value}, s[i/4]...)
			}
		}
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			continue
		}
		nb, from, to := parseMove(line)

		for i := 0; i < nb; i++ {
			value := s[from][len(s[from])-1]
			s[to] = append(s[to], value)
			s[from] = s[from][:len(s[from])-1]
		}
	}

	result := ""
	for _, stack := range s {
		result += stack[len(stack)-1]
	}
	return result
}

func parseMove(line string) (int, int, int) {
	split := strings.Split(line, " ")
	nb, _ := strconv.Atoi(split[1])
	from, _ := strconv.Atoi(split[3])
	to, _ := strconv.Atoi(split[5])

	return nb, from, to
}
