package d

import (
	"bufio"
	"io"
	"sort"
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

type directory struct {
	name     string
	files    []file
	children []*directory
	parent   *directory
}

type file struct {
	name string
	size int
}

func (d *D) Run() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	allDirectories := make([]*directory, 0)
	root := directory{name: "/", files: make([]file, 0)}
	allDirectories = append(allDirectories, &root)
	currentDirectory := &root
	currentCommand := ""

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line[0] == '$' {
			currentCommand = line[1:]
			if currentCommand[1:3] == "ls" {
				continue
			}
			if currentCommand[1:3] == "cd" {
				expectedDirectory := currentCommand[4:]
				if expectedDirectory == ".." {
					currentDirectory = currentDirectory.parent
					continue
				}
				for _, directory := range currentDirectory.children {
					if directory.name == expectedDirectory {
						currentDirectory = directory
						break
					}
				}
				continue
			}
		}

		if line[0:3] == "dir" {
			dirName := line[4:]
			childDirectory := directory{name: dirName, files: make([]file, 0), parent: currentDirectory}
			currentDirectory.children = append(currentDirectory.children, &childDirectory)
			allDirectories = append(allDirectories, &childDirectory)
		} else {
			split := strings.Split(line, " ")
			fileSize, _ := strconv.Atoi(split[0])
			fileName := split[1]
			currentDirectory.files = append(currentDirectory.files, file{name: fileName, size: fileSize})
		}
	}

	var allSums []int
	for _, d := range allDirectories {
		sumOfDirectory := calcSumOfDirectory(d)
		allSums = append(allSums, sumOfDirectory)
	}

	sort.Ints(allSums)
	biggestSum := allSums[len(allSums)-1]
	availableSpace := 70_000_000 - biggestSum
	requiredSpace := 30_000_000 - availableSpace

	for i := 0; i < len(allSums); i++ {
		if allSums[i] >= requiredSpace {
			return allSums[i]
		}
	}
	return 0
}

func calcSumOfDirectory(d *directory) int {
	sumOfDirectory := 0
	for _, file := range d.files {
		sumOfDirectory += file.size
	}

	for _, child := range d.children {
		sumOfDirectory += calcSumOfDirectory(child)
	}

	return sumOfDirectory
}

func (d *D) RunStr() string {
	return ""
}
