package day18

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	strings "strings"
)

type coordinates struct {
	x, y int
}

type path []coordinates

var directions = []coordinates{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func solvePartOne() int {
	inputBytes, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	size := 70
	start, end := coordinates{0, 0}, coordinates{size, size}

	sliceSize := 1024

	bytesSlice := inputBytes[:sliceSize]
	fallenBytes := make(map[coordinates]struct{})
	for _, b := range bytesSlice {
		fallenBytes[b] = struct{}{}
	}

	return findMinStepsToReachTarget(start, end, fallenBytes, size)

}

func findMinStepsToReachTarget(start, end coordinates, fallenBytes map[coordinates]struct{}, size int) int {

	visited := make(map[coordinates]struct{})

	queue := []path{{start}}

	var currentPath path

	for len(queue) > 0 {
		currentPath, queue = queue[0], queue[1:]

		lastPos := currentPath[len(currentPath)-1]

		if lastPos == end {
			return len(currentPath) - 1
		}

		if _, ok := visited[lastPos]; ok {
			continue
		}

		visited[lastPos] = struct{}{}

		for _, d := range directions {
			newPosition := coordinates{lastPos.x + d.x, lastPos.y + d.y}
			if notOutOfRange(newPosition, size, size) {
				if _, ok := fallenBytes[newPosition]; !ok {
					newPath := append(slices.Clone(currentPath), newPosition)
					queue = append(queue, newPath)
				}
			}
		}
	}
	return -1
}

func solvePartTwo() coordinates {
	inputBytes, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	size := 70
	start, end := coordinates{0, 0}, coordinates{size, size}

	var blockadeByte coordinates

	for i := 1; i < len(inputBytes); i++ {

		fallenBytes := make(map[coordinates]struct{})

		for _, b := range inputBytes[:i] {
			fallenBytes[b] = struct{}{}
		}

		if findMinStepsToReachTarget(start, end, fallenBytes, size) == -1 {
			blockadeByte = inputBytes[i-1]
			break
		}

	}

	return blockadeByte
}

func readInput(fileName string) ([]coordinates, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []coordinates
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cs := strings.Split(line, ",")
		x, errX := strconv.Atoi(cs[0])
		y, errY := strconv.Atoi(cs[1])
		if errX != nil || errY != nil {
			return nil, err
		}
		data = append(data, coordinates{x, y})
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}

func notOutOfRange(pos coordinates, maxX, maxY int) bool {
	return pos.x >= 0 && pos.x <= maxX && pos.y >= 0 && pos.y <= maxY
}
