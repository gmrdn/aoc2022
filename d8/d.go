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

	visibleTrees := 0
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
				if trees[i][left].height < trees[i][j].height {
					trees[i][j].canSeeL++
				}
			}

			for right := j + 1; right < len(trees[i]); right++ {
				if trees[i][right].height < trees[i][j].height {
					trees[i][j].canSeeR++
				}
			}

			for up := i - 1; up >= 0; up-- {
				if trees[up][j].height < trees[i][j].height {
					trees[i][j].canSeeU++
				}
			}

			for down := i + 1; down < len(trees); down++ {
				if trees[down][j].height < trees[i][j].height {
					trees[i][j].canSeeD++
				}
			}

			if trees[i][j].canSeeL == j || trees[i][j].canSeeR == len(trees[i])-j-1 || trees[i][j].canSeeU == i || trees[i][j].canSeeD == len(trees)-i-1 {
				visibleTrees++
			}

		}
	}

	return visibleTrees
}

func (d *D) RunStr() string {
	return ""
}
