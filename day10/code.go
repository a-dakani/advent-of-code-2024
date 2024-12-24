package day10

import (
	"bufio"
	"log"
	"os"
	"slices"
	"unicode"
)

type location struct {
	x, y int
}
type path []location

func solvePartOne() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	score := 0

	for y, line := range input {
		for x, p := range line {
			if p == 0 {
				startPoint := location{x, y}
				score += len(findUniqueEndPositions(input, startPoint, []location{}))
			}
		}
	}

	return score
}

func findUniqueEndPositions(input [][]int, startPoint location, knownEnds []location) []location {

	if input[startPoint.y][startPoint.x] == 9 {
		// if the new end is unique add it to the list of known ends
		if !slices.Contains(knownEnds, startPoint) {
			knownEnds = append(knownEnds, startPoint)
		}
		return knownEnds
	}

	nextSteps := getPossibleNextSteps(input, startPoint)
	if len(nextSteps) == 0 {
		return knownEnds
	}

	for _, next := range nextSteps {
		knownEnds = findUniqueEndPositions(input, next, knownEnds)
	}

	return knownEnds
}

func findPossiblePaths(input [][]int, paths []path) []path {
	var newPaths []path

	for _, p := range paths {
		lastLocation := p[len(p)-1]

		if input[lastLocation.y][lastLocation.x] == 9 {
			newPaths = append(newPaths, p)
			continue
		}

		nextSteps := getPossibleNextSteps(input, lastLocation)
		if len(nextSteps) == 0 {
			//there are no more steps to take so this path is invalid
			continue
		}

		for _, next := range nextSteps {
			//	here is the recursive call to follow the path
			newPath := make(path, len(p))
			copy(newPath, p)
			newPath = append(newPath, next)
			newPaths = append(newPaths, findPossiblePaths(input, []path{newPath})...)
		}
	}
	return newPaths
}

func getPossibleNextSteps(input [][]int, l location) []location {
	var nextSteps []location
	directions := []location{
		{0, -1}, //up
		{1, 0},  // right
		{0, 1},  // down
		{-1, 0}, // left
	}

	for _, dir := range directions {
		next := location{l.x + dir.x, l.y + dir.y}
		if isOnMap(next, input) && input[next.y][next.x] == input[l.y][l.x]+1 {
			nextSteps = append(nextSteps, next)
		}
	}
	return nextSteps
}

func solvePartTwo() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	var possiblePaths []path

	for y, line := range input {
		for x, p := range line {
			if p == 0 {
				// found starting point now follow the path
				pathStart := []path{{{x, y}}}
				possiblePaths = append(possiblePaths, findPossiblePaths(input, pathStart)...)
			}
		}
	}

	return len(possiblePaths)
}

func isOnMap(p location, input [][]int) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	if p.y >= len(input) || p.x >= len(input[p.y]) {
		return false
	}
	return true
}

func readInput(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var datum []int
		text := scanner.Text()
		for _, char := range text {
			if unicode.IsDigit(char) {
				datum = append(datum, int(char-'0'))
			}
		}
		data = append(data, datum)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
