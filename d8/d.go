package d

import (
	"bufio"
	"io"
	"strconv"
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

type tree struct {
	height  int
	canSeeL int
	canSeeR int
	canSeeU int
	canSeeD int
}

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	var trees [][]tree

	currentRow := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()

		currentRowOfTrees := make([]tree, 0)
		for i := 0; i < len(line); i++ {
			h, _ := strconv.Atoi(string(line[i]))

			tree := tree{
				height: h,
			}

			currentRowOfTrees = append(currentRowOfTrees, tree)

		}

		trees = append(trees, currentRowOfTrees)
		currentRow++
	}

	maxScenicScore := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees[i]); j++ {

			if i == 0 {
				trees[i][j].canSeeU = 0
			}

			if i == len(trees)-1 {
				trees[i][j].canSeeD = 0
			}

			if j == 0 {
				trees[i][j].canSeeL = 0
			}

			if j == len(trees[i])-1 {
				trees[i][j].canSeeR = 0
			}

			for left := j - 1; left >= 0; left-- {
				trees[i][j].canSeeL++
				if trees[i][left].height >= trees[i][j].height {
					break
				}
			}

			for right := j + 1; right < len(trees[i]); right++ {
				trees[i][j].canSeeR++
				if trees[i][right].height >= trees[i][j].height {
					break
				}
			}

			for up := i - 1; up >= 0; up-- {
				trees[i][j].canSeeU++
				if trees[up][j].height >= trees[i][j].height {
					break
				}
			}

			for down := i + 1; down < len(trees); down++ {
				trees[i][j].canSeeD++
				if trees[down][j].height >= trees[i][j].height {
					break
				}
			}

			scenicScore := trees[i][j].canSeeL * trees[i][j].canSeeR * trees[i][j].canSeeU * trees[i][j].canSeeD

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}

		}
	}

	return maxScenicScore
}

func (d *D) RunStr() string {
	return ""
}
