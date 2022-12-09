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
					adjustDown(nodes[i], nodes[i+1])
					adjustRight(nodes[i], nodes[i+1])
					adjustLeft(nodes[i], nodes[i+1])
				}
			case "D":
				nodes[0].y--
				for i := 0; i < 9; i++ {
					adjustUp(nodes[i], nodes[i+1])
					adjustDown(nodes[i], nodes[i+1])
					adjustRight(nodes[i], nodes[i+1])
					adjustLeft(nodes[i], nodes[i+1])
				}
			case "L":
				nodes[0].x--
				for i := 0; i < 9; i++ {
					adjustUp(nodes[i], nodes[i+1])
					adjustDown(nodes[i], nodes[i+1])
					adjustRight(nodes[i], nodes[i+1])
					adjustLeft(nodes[i], nodes[i+1])
				}
			case "R":
				nodes[0].x++
				for i := 0; i < 9; i++ {
					adjustUp(nodes[i], nodes[i+1])
					adjustDown(nodes[i], nodes[i+1])
					adjustRight(nodes[i], nodes[i+1])
					adjustLeft(nodes[i], nodes[i+1])
				}

			}
			visited[*nodes[9]] = true
			// visualizeNodes(nodes)
		}
	}
	return len(visited)
}

func visualizeNodes(nodes map[int]*node) {
	grid := make([][]string, 21)
	for i := 0; i < 21; i++ {
		grid[i] = make([]string, 26)
	}

	for i := 0; i < 10; i++ {
		grid[nodes[i].y+5][nodes[i].x+11] = strconv.Itoa(i)
	}

	for i := 0; i < 21; i++ {
		for j := 0; j < 26; j++ {
			if grid[i][j] == "" {
				grid[i][j] = "."
			}
		}
	}

	for i := 0; i < 21; i++ {
		for j := 0; j < 26; j++ {
			print(grid[i][j])
		}
		println()
	}

	fmt.Println("=====================================")
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
