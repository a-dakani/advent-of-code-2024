package day06

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

type direction rune

const (
	up    direction = '^'
	down  direction = 'v'
	left  direction = '<'
	right direction = '>'
)
const obstacle = '#'
const tracePosition = 'X'

func solvePartOne() int {

	guardMap, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	x, y, err := findGuard(guardMap)
	if err != nil {
		panic(err)
	}

	finalMap := guardMap
	for {
		newX, newY, newMap, existed := guardWalk(guardMap, x, y)

		if existed {
			break
		}
		x, y, guardMap = newX, newY, newMap
	}

	return calculateDistinctPositions(finalMap)
}

func calculateDistinctPositions(finalMap [][]rune) int {
	//	Calculate the number of distinct positions
	distinctPositions := 0
	for _, row := range finalMap {
		for _, elem := range row {
			if elem == tracePosition {
				distinctPositions++
			}
		}
	}
	//add one to account for the last position before exiting
	return distinctPositions + 1
}

func printGuardMap(guardMap [][]rune) {
	for _, row := range guardMap {
		for _, elem := range row {
			fmt.Printf("%c ", elem)
		}
		fmt.Println()
	}
	fmt.Println("----------------------")
}

func guardWalk(guardMap [][]rune, x int, y int) (int, int, [][]rune, bool) {
	guard := guardMap[y][x]
	if direction(guard) == up {
		return walkUp(guardMap, x, y)
	}
	if direction(guard) == down {
		return walkDown(guardMap, x, y)
	}
	if direction(guard) == left {
		return walkLeft(guardMap, x, y)
	}
	if direction(guard) == right {
		return walkRight(guardMap, x, y)
	}
	return -1, -1, nil, false
}

func walkRight(guardMap [][]rune, x int, y int) (int, int, [][]rune, bool) {
	for ix := x; ix < len(guardMap[y]); ix++ {
		if ix == len(guardMap[y])-1 {
			moveGuard(x, y, ix, y, guardMap)
			return ix, y, guardMap, true
		}
		if guardMap[y][ix+1] == obstacle {
			rotateGuardNinetyDegrees(ix, y, guardMap)
			return ix, y, guardMap, false
		}
		traceGuardPosition(ix, y, ix+1, y, guardMap)
	}
	return -1, -1, nil, false

}

func walkLeft(guardMap [][]rune, x int, y int) (int, int, [][]rune, bool) {
	for ix := x; ix >= 0; ix-- {
		if ix == 0 {
			moveGuard(x, y, ix, y, guardMap)
			return ix, y, guardMap, true
		}
		if guardMap[y][ix-1] == obstacle {
			rotateGuardNinetyDegrees(ix, y, guardMap)
			return ix, y, guardMap, false
		}
		traceGuardPosition(ix, y, ix-1, y, guardMap)
	}
	return -1, -1, nil, false
}

func walkDown(guardMap [][]rune, x int, y int) (int, int, [][]rune, bool) {
	for iy := y; iy < len(guardMap); iy++ {
		if iy == len(guardMap)-1 {
			moveGuard(x, y, x, iy, guardMap)
			return x, iy, guardMap, true
		}
		if guardMap[iy+1][x] == obstacle {
			rotateGuardNinetyDegrees(x, iy, guardMap)
			return x, iy, guardMap, false
		}
		traceGuardPosition(x, iy, x, iy+1, guardMap)
	}
	return -1, -1, nil, false

}

func walkUp(guardMap [][]rune, x int, y int) (int, int, [][]rune, bool) {
	for iy := y; iy >= 0; iy-- {
		if iy == 0 {
			moveGuard(x, y, x, iy, guardMap)
			return x, iy, guardMap, true
		}
		if guardMap[iy-1][x] == obstacle {
			rotateGuardNinetyDegrees(x, iy, guardMap)
			return x, iy, guardMap, false
		}
		traceGuardPosition(x, iy, x, iy-1, guardMap)
	}
	return -1, -1, nil, false
}

func traceGuardPosition(fromX int, fromY, toX, toY int, guardMap [][]rune) {
	guardMap[toY][toX] = guardMap[fromY][fromX]
	guardMap[fromY][fromX] = tracePosition
}

func rotateGuardNinetyDegrees(x int, y int, guardMap [][]rune) {
	switch guardMap[y][x] {
	case rune(up):
		guardMap[y][x] = rune(right)
	case rune(down):
		guardMap[y][x] = rune(left)
	case rune(left):
		guardMap[y][x] = rune(up)
	case rune(right):
		guardMap[y][x] = rune(down)
	default:
		panic("invalid guard direction")
	}
}

func moveGuard(oldX, oldY, newX, newY int, guardMap [][]rune) {
	guardMap[oldY][oldX], guardMap[newY][newX] = guardMap[newY][newX], guardMap[oldY][oldX]
}

func findGuard(guardMap [][]rune) (int, int, error) {
	for i, line := range guardMap {
		for k, position := range line {
			if position == rune(up) || position == rune(down) || position == rune(left) || position == rune(right) {
				return k, i, nil
			}
		}
	}
	return -1, -1, errors.New("guard not found")

}

func solvePartTwo() int {
	return 0
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
