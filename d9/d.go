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

type node struct {
	x, y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	nodes := make(map[int]*node)
	for i := 0; i < 10; i++ {
		nodes[i] = &node{0, 0}
	}
	visited := make(map[node]bool)

	visited[*nodes[9]] = true

	for fileScanner.Scan() {
		line := fileScanner.Text()
		move := strings.Split(line, " ")
		direction := move[0]
		steps, _ := strconv.Atoi(move[1])

		for i := 0; i < steps; i++ {
			switch direction {
			case "U":
				nodes[0].y++
				for i := 0; i < 9; i++ {
					adjustUp(nodes[i], nodes[i+1])
				}
			case "D":
				nodes[0].y--
				for i := 0; i < 9; i++ {
					adjustDown(nodes[i], nodes[i+1])
				}
			case "L":
				nodes[0].x--
				for i := 0; i < 9; i++ {
					adjustLeft(nodes[i], nodes[i+1])
				}
			case "R":
				nodes[0].x++
				for i := 0; i < 9; i++ {
					adjustRight(nodes[i], nodes[i+1])
				}

			}
			visited[*nodes[9]] = true
		}
	}
	return len(visited)
}

func adjustUp(h, t *node) {
	if h.y-t.y > 1 {
		t.y++
		if h.x-t.x > 0 {
			t.x++
		} else if h.x-t.x < 0 {
			t.x--
		}
	}
}

func adjustDown(h, t *node) {
	if t.y-h.y > 1 {
		t.y--
		if h.x-t.x > 0 {
			t.x++
		} else if h.x-t.x < 0 {
			t.x--
		}
	}
}

func adjustRight(h, t *node) {
	if h.x-t.x > 1 {
		t.x++
		if h.y-t.y > 0 {
			t.y++
		} else if h.y-t.y < 0 {
			t.y--
		}
	}
}

func adjustLeft(h, t *node) {
	if t.x-h.x > 1 {
		t.x--
		if h.y-t.y > 0 {
			t.y++
		} else if h.y-t.y < 0 {
			t.y--
		}
	}
}

func (d *D) RunStr() string {
	return ""
}
