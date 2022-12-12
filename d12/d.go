package d

import (
	"bufio"
	"io"
	"math"
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
	matrix, startingPositions, dest := d.ParseInput()

	minDistance := math.MaxInt32

	for _, startingPosition := range startingPositions {
		pathLength := d.ShortestPath(matrix, startingPosition, dest, minDistance)
		if pathLength < minDistance {
			minDistance = pathLength
		}
	}

	return minDistance
}

func (d *D) RunStr() string {
	return ""
}

func (d *D) ParseInput() ([][]byte, []position, position) {
	result := [][]byte{}
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	currentRow := 0

	startingPositions := []position{}
	dest := position{0, 0}
	for fileScanner.Scan() {
		line := fileScanner.Text()

		currentLine := []byte{}
		for i, char := range line {

			if char == 'a' {
				startingPositions = append(startingPositions, position{i, currentRow})
			}
			if char == 'S' {
				startingPositions = append(startingPositions, position{i, currentRow})
				char = 'a'
			}
			if char == 'E' {
				dest = position{i, currentRow}
				char = 'z'
			}

			currentLine = append(currentLine, byte(char))
		}

		result = append(result, currentLine)

		currentRow++
	}

	return result, startingPositions, dest
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

func (d *D) ShortestPath(mat [][]byte, src position, dest position, currentMinDistance int) int {
	queue := []node{
		{
			x:    src.x,
			y:    src.y,
			dist: 0,
		},
	}

	visited := make(map[position]bool)
	visited[src] = true

	minDistance := math.MaxInt32

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currentAltitude := int(mat[current.y][current.x])

		dist := current.dist
		if dist > currentMinDistance {
			return math.MaxInt32
		}

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
