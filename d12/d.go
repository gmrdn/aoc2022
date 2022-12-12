package d

import (
	"bufio"
	"fmt"
	"io"
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
	matrix, src, dest := d.ParseInput()
	pathLength := d.ShortestPath(matrix, src, dest)

	return pathLength
}

func (d *D) RunStr() string {
	return ""
}

func (d *D) ParseInput() ([][]byte, position, position) {
	result := [][]byte{}
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	currentRow := 0

	src := position{0, 0}
	dest := position{0, 0}
	for fileScanner.Scan() {
		line := fileScanner.Text()

		srcIndex := strings.Index(line, "S")
		if srcIndex > -1 {
			src = position{srcIndex, currentRow}
			line = strings.Replace(line, "S", "a", 1)
		}

		destIndex := strings.Index(line, "E")
		if destIndex > -1 {
			dest = position{destIndex, currentRow}
			line = strings.Replace(line, "E", "z", 1)
		}

		result = append(result, []byte(line))
		currentRow++
	}

	return result, src, dest
}

type position struct {
	x int
	y int
}

type node struct {
	x    int
	y    int
	dist int
}

func (d *D) ShortestPath(mat [][]byte, src position, dest position) int {
	queue := []node{
		{
			x:    src.x,
			y:    src.y,
			dist: 0,
		},
	}

	visited := make(map[position]bool)
	visited[src] = true

	minDistance := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentAltitude := int(mat[current.y][current.x])
		fmt.Println(currentAltitude)

		dist := current.dist

		if current.x == dest.x && current.y == dest.y {
			minDistance = dist
			break
		}

		for _, next := range d.GetNextPositions(mat, current) {
			if isValid(mat, visited, position{next.x, next.y}, currentAltitude) {
				visited[position{next.x, next.y}] = true
				queue = append(queue, next)
			}
		}

	}

	return minDistance
}

func (d *D) GetNextPositions(mat [][]byte, current node) []node {
	positions := []node{}
	positions = append(positions, node{
		x:    current.x,
		y:    current.y - 1,
		dist: current.dist + 1,
	})

	positions = append(positions, node{
		x:    current.x,
		y:    current.y + 1,
		dist: current.dist + 1,
	})

	positions = append(positions, node{
		x:    current.x - 1,
		y:    current.y,
		dist: current.dist + 1,
	})

	positions = append(positions, node{
		x:    current.x + 1,
		y:    current.y,
		dist: current.dist + 1,
	})
	return positions
}

func isValid(mat [][]byte, visited map[position]bool, pos position, currentAltitude int) bool {
	if pos.x < 0 || pos.x >= len(mat[0]) {
		return false
	}

	if pos.y < 0 || pos.y >= len(mat) {
		return false
	}

	if visited[pos] {
		return false
	}

	altitude := int(mat[pos.y][pos.x])

	return altitude <= currentAltitude+1
}
