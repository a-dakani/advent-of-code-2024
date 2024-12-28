package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
)

type coordinates struct {
	x, y int
}

var (
	left  = coordinates{-1, 0}
	right = coordinates{1, 0}
	up    = coordinates{0, -1}
	down  = coordinates{0, 1}
)

const (
	crateLeftSide  = '['
	crateRightSide = ']'
	crate          = 'O'
	emptySpace     = '.'
	robot          = '@'
	wall           = '#'
)

func solvePartOne() int {
	mapData, directions, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	robotsCoordinates := findRune(mapData, robot)
	if len(robotsCoordinates) != 1 {
		log.Fatalf("unexpected number of robots in the map")
	}
	robotCoordinates := robotsCoordinates[0]

	for _, d := range directions {
		newMap, newRobotCoordinates := moveRobot(mapData, robotCoordinates, d)
		mapData = newMap
		robotCoordinates = newRobotCoordinates
	}
	cratesCoordinates := findRune(mapData, crate)
	score := 0

	for _, c := range cratesCoordinates {
		score += 100*c.y + c.x
	}

	return score
}

func printMap(data [][]rune) {
	fmt.Println("-------------------------------------------------")
	for _, line := range data {
		fmt.Println(string(line))
	}
	fmt.Println("-------------------------------------------------")
}

func moveRobot(mapData [][]rune, robotCoordinates coordinates, direction coordinates) ([][]rune, coordinates) {
	movablesInWay, movable := isObjectMovable(mapData, robotCoordinates, direction)

	if !movable {
		return mapData, robotCoordinates
	}

	// get a unique list of movables
	var uniqueMovables []coordinates
	for _, m := range movablesInWay {
		if !slices.Contains(uniqueMovables, m) {
			uniqueMovables = append(uniqueMovables, m)
		}
	}

	shiftObjectsInDirection(mapData, uniqueMovables, direction)

	//we check if the robot can move and have an empty space in the direction then we move the robot
	if mapData[robotCoordinates.y+direction.y][robotCoordinates.x+direction.x] == emptySpace {
		switchObjects(mapData, robotCoordinates, coordinates{robotCoordinates.x + direction.x, robotCoordinates.y + direction.y})
	} else {
		log.Fatalf("unexpected object in the direction of the robot")
	}

	robotCoordinates.x += direction.x
	robotCoordinates.y += direction.y

	return mapData, robotCoordinates
}

func switchObjects(mapData [][]rune, objectA coordinates, objectB coordinates) {
	mapData[objectA.y][objectA.x], mapData[objectB.y][objectB.x] = mapData[objectB.y][objectB.x], mapData[objectA.y][objectA.x]
}

func shiftObjectsInDirection(mapData [][]rune, movables []coordinates, direction coordinates) {

	// We sort the movables so that we move them in the correct order
	// Example: If we shift left "x = -1" we want to move the leftmost object first
	sort.Slice(movables, func(i, j int) bool {
		switch direction {
		case left:
			return movables[i].x < movables[j].x
		case right:
			return movables[i].x > movables[j].x
		case up:
			return movables[i].y < movables[j].y
		case down:
			return movables[i].y > movables[j].y
		}
		return false
	})

	// Move each object in the direction
	for _, obj := range movables {
		switchObjects(mapData, obj, coordinates{obj.x + direction.x, obj.y + direction.y})
	}
}

func isObjectMovable(mapData [][]rune, objectCoordinates coordinates, direction coordinates) ([]coordinates, bool) {
	var movablesCoordinates []coordinates
	movable := false

	nextX := objectCoordinates.x + direction.x
	nextY := objectCoordinates.y + direction.y

	for {
		// not really necessary to check if the next position is out of bounds, because the map is surrounded by walls
		if nextX < 0 || nextX >= len(mapData[0]) || nextY < 0 || nextY >= len(mapData) {
			break
		}
		if mapData[nextY][nextX] == emptySpace {
			movable = true
			break
		}
		if slices.Contains([]rune{crate, crateRightSide, crateLeftSide}, mapData[nextY][nextX]) {
			movablesCoordinates = append(movablesCoordinates, coordinates{nextX, nextY})
		}

		if mapData[nextY][nextX] == wall {
			movable = false
			break
		}

		// if we are moving up or down we need to calculate the other side of the crate and all crates in the way
		if slices.Contains([]coordinates{up, down}, direction) {
			if mapData[nextY][nextX] == crateLeftSide {
				// we use recursive call to check if the crate can be moved
				rightSideCoordinates := coordinates{nextX + 1, nextY}
				movablesCoordinates = append(movablesCoordinates, rightSideCoordinates)

				rightSideObjects, rightSideMovable := isObjectMovable(mapData, rightSideCoordinates, direction)
				movablesCoordinates = append(movablesCoordinates, rightSideObjects...)
				if !rightSideMovable {
					movable = false
					break
				}
			}

			if mapData[nextY][nextX] == crateRightSide {
				// we use recursive call to check if the crate can be moved
				leftSideCoordinates := coordinates{nextX - 1, nextY}
				movablesCoordinates = append(movablesCoordinates, leftSideCoordinates)

				leftSideObjects, leftSideMovable := isObjectMovable(mapData, leftSideCoordinates, direction)
				movablesCoordinates = append(movablesCoordinates, leftSideObjects...)
				if !leftSideMovable {
					movable = false
					break
				}
			}
		}

		nextX = nextX + direction.x
		nextY = nextY + direction.y
	}
	return movablesCoordinates, movable
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

func widenTheMap(mapData [][]rune) {
	for y := range mapData {
		newLine := make([]rune, 0, len(mapData[y])*2)
		for x := range mapData[y] {
			switch mapData[y][x] {
			case wall:
				newLine = append(newLine, wall, wall)
			case crate:
				newLine = append(newLine, crateLeftSide, crateRightSide)
			case emptySpace:
				newLine = append(newLine, emptySpace, emptySpace)
			case robot:
				newLine = append(newLine, robot, emptySpace)
			}
		}
		mapData[y] = newLine
	}
}

func solvePartTwo() int {
	mapData, directions, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	widenTheMap(mapData)

	robotsCoordinates := findRune(mapData, robot)
	if len(robotsCoordinates) != 1 {
		log.Fatalf("unexpected number of robots in the map")
	}
	robotCoordinates := robotsCoordinates[0]

	for _, d := range directions {
		newMap, newRobotCoordinates := moveRobot(mapData, robotCoordinates, d)
		mapData = newMap
		robotCoordinates = newRobotCoordinates

	}
	cratesCoordinates := findRune(mapData, crateLeftSide)
	score := 0

	for _, c := range cratesCoordinates {
		score += 100*c.y + c.x
	}

	return score
}

func readInput(fileName string) ([][]rune, []coordinates, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var mapData [][]rune
	var directions []coordinates
	readingMap := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// If the line is empty, we have finished reading the map
		if line == "" {
			readingMap = false
			continue
		}

		runesInLine := []rune(line)

		// If we haven't finished reading the map, add the line to the map data
		if readingMap {
			mapData = append(mapData, runesInLine)
			continue
		}

		// If we have finished reading the map, add the line to the directions
		// we map the directions to coordinates for ease of use
		for _, r := range runesInLine {
			switch r {
			case '^':
				directions = append(directions, up)
			case 'v':
				directions = append(directions, down)
			case '>':
				directions = append(directions, right)
			case '<':
				directions = append(directions, left)
			default:
				log.Fatalf("unexpected character in input: %c", r)
			}
		}

	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, nil, scannerErr
	}

	return mapData, directions, nil
}
