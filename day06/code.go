package day06

import (
	"bufio"
	"log"
	"os"
)

type direction coordinates

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	right = direction{1, 0}
	left  = direction{-1, 0}
)

type coordinates struct {
	x, y int
}

type step struct {
	coordinates
	direction
}

type path []step

const obstacle = '#'
const emptySpave = '.'
const guard = '^'

func solvePartOne() int {

	guardMap, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	startPosition := getGuardStartPosition(guardMap)

	route, _ := walkTheMapUntilOutOfRange(guardMap, startPosition)

	uniqueCoordinates := make(map[coordinates]struct{})
	for _, s := range route {
		uniqueCoordinates[s.coordinates] = struct{}{}
	}

	return len(uniqueCoordinates)
}

func solvePartTwo() int {

	guardMap, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	guardCoordinates := findRune(guardMap, guard)
	if len(guardCoordinates) != 1 {
		log.Fatalf("unexpected number of start coordinates in the map")
	}

	startPoint := guardCoordinates[0]
	startPosition := step{
		coordinates: startPoint,
		direction:   up,
	}

	route, _ := walkTheMapUntilOutOfRange(guardMap, startPosition)

	possibleLoopCoordinates := findPossibleLoopCoordinates(guardMap, route)

	uniqueCoordinates := make(map[coordinates]struct{})
	for _, s := range possibleLoopCoordinates {
		uniqueCoordinates[s] = struct{}{}
	}

	return len(uniqueCoordinates)
}

func walkTheMapUntilOutOfRange(guardMap [][]rune, startPosition step) (path, bool) {

	inLoop := false
	var route path
	visited := make(map[step]struct{})

	route = append(route, startPosition)

	for {
		lastStep := route[len(route)-1]
		nextStep, outOfRange := getNextStepOrOutOfRange(guardMap, lastStep)
		if outOfRange {
			break
		}
		if _, ok := visited[nextStep]; ok {
			inLoop = true
			break
		}
		visited[nextStep] = struct{}{}
		route = append(route, nextStep)
	}

	return route, inLoop
}

func getNextStepOrOutOfRange(guardMap [][]rune, lastStep step) (step, bool) {
	nextX := lastStep.coordinates.x + lastStep.direction.x
	nextY := lastStep.coordinates.y + lastStep.direction.y

	if nextX < 0 || nextX >= len(guardMap[0]) || nextY < 0 || nextY >= len(guardMap) {
		return step{}, true
	}

	if guardMap[nextY][nextX] == obstacle {
		switch lastStep.direction {
		case up:
			lastStep.direction = right
		case right:
			lastStep.direction = down
		case down:
			lastStep.direction = left
		case left:
			lastStep.direction = up
		}
		return getNextStepOrOutOfRange(guardMap, lastStep)
	}

	return step{
		coordinates: coordinates{nextX, nextY},
		direction:   lastStep.direction,
	}, false
}

func findPossibleLoopCoordinates(guardMap [][]rune, route path) []coordinates {

	startPosition := getGuardStartPosition(guardMap)

	var possibleLoopCoordinates []coordinates
	for i, s := range route {
		if i == 0 {
			continue
		}

		guardMap[s.coordinates.y][s.coordinates.x] = obstacle

		_, isInLoop := walkTheMapUntilOutOfRange(guardMap, startPosition)

		if isInLoop {
			possibleLoopCoordinates = append(possibleLoopCoordinates, s.coordinates)
		}

		guardMap[s.coordinates.y][s.coordinates.x] = emptySpave
	}

	return possibleLoopCoordinates
}

func findRune(data [][]rune, searchable rune) []coordinates {
	var found []coordinates
	for y, line := range data {
		for x, r := range line {
			if r == searchable {
				found = append(found, coordinates{x, y})
			}
		}
	}
	return found
}

func getGuardStartPosition(guardMap [][]rune) step {
	guardCoordinates := findRune(guardMap, guard)
	if len(guardCoordinates) != 1 {
		log.Fatalf("unexpected number of start coordinates in the map")
	}
	startPoint := guardCoordinates[0]

	return step{
		coordinates: startPoint,
		direction:   up,
	}
}

func readTextLines(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := []rune(scanner.Text())
		data = append(data, chars)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
