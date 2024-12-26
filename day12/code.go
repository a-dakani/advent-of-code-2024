package day12

import (
	"bufio"
	"log"
	"os"
)

type location struct {
	x, y int
}

type blockSize struct {
	size, perimeter, corners int
}
type discountedBlockSize struct {
	size, sides int
}

var directions = []location{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

var diagonalDirections = []location{
	{-1, -1}, // up left
	{1, -1},  // up right
	{1, 1},   // down right
	{-1, 1},  // down left
}

func solvePartOne() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	visited := map[location]struct{}{}
	cost := 0

	for y, line := range input {
		for x, _ := range line {
			l := location{x, y}
			if _, ok := visited[l]; ok {
				continue
			}

			bs := getBlockSize(input, l, visited)
			//fmt.Printf("tree type %c has a size of %d and a peremiter of %d\n", input[y][x], bs.size, bs.perimeter)
			cost += bs.perimeter * bs.size
		}
	}

	return cost
}

func getBlockSize(input [][]rune, l location, visited map[location]struct{}) blockSize {
	var result blockSize

	//mark the start location as visited
	visited[l] = struct{}{}

	queue := []location{l}

	// helper location of the reference tree
	var referenceTree location

	for len(queue) > 0 {

		// get the actual tree and remove it from queue
		referenceTree, queue = queue[0], queue[1:]
		referenceTreeType := input[referenceTree.y][referenceTree.x]

		//we have a tree, so size increased
		result.size++

		//we iterate in all directions and search for the same type of tree
		for _, d := range directions {

			treeInDirection := location{referenceTree.x + d.x, referenceTree.y + d.y}
			//if the tree in direction is out of range this means that we have reached an edge.
			// we add a perimeter and continue
			if !isOnMap(treeInDirection, input) {
				result.perimeter++
				continue
			}
			treeInDirectionType := input[treeInDirection.y][treeInDirection.x]

			// if the tree in direction is the same type and has not been visited
			// then it is a new tree and need to be queued
			if referenceTreeType == treeInDirectionType {
				if _, ok := visited[treeInDirection]; !ok {
					queue = append(queue, treeInDirection)
					visited[treeInDirection] = struct{}{}
				}
			} else {
				// if the tree is already visited then we need to close the perimeter with a new fence
				result.perimeter++
			}
		}

		// check for corners
		for _, d := range diagonalDirections {

			// we try to get the corner and sides of the corner.
			// if out of bound then we define it as empty space for simplicity
			// WARNING: This only works if the map does not have empty spaces
			corner, cornerX, cornerY := getDiagonalCornerAndSides(input, referenceTree, d)

			// if x and y side of the corner are different from the reference tree. then we have a convex corner
			if referenceTreeType != cornerX &&
				referenceTreeType != cornerY {
				result.corners++
			}

			// if x and y side of the corner are the same as the reference tree. then we have a concave corner
			if referenceTreeType == cornerX &&
				referenceTreeType == cornerY &&
				referenceTreeType != corner {
				result.corners++
			}
		}

	}

	return result
}

func solvePartTwo() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	visited := map[location]struct{}{}
	cost := 0

	for y, line := range input {
		for x, _ := range line {
			l := location{x, y}
			if _, ok := visited[l]; ok {
				continue
			}

			bs := getBlockSize(input, l, visited)
			//getting corners to count the sides
			//fmt.Printf("tree type %c has a size of %d and corners of %d\n", input[y][x], bs.size, bs.corners)
			cost += bs.corners * bs.size
		}
	}

	return cost
}

func getDiagonalCornerAndSides(input [][]rune, refTree location, d location) (rune, rune, rune) {
	var corner, cornerX, cornerY rune
	if isOnMap(location{refTree.x + d.x, refTree.y + d.y}, input) {
		corner = input[refTree.y+d.y][refTree.x+d.x]
	} else {
		corner = ' '
	}
	if isOnMap(location{refTree.x + d.x, refTree.y}, input) {
		cornerX = input[refTree.y][refTree.x+d.x]
	} else {
		cornerX = ' '
	}
	if isOnMap(location{refTree.x, refTree.y + d.y}, input) {
		cornerY = input[refTree.y+d.y][refTree.x]
	} else {
		cornerY = ' '
	}
	return corner, cornerX, cornerY
}

func isOnMap(p location, input [][]rune) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	if p.y >= len(input) || p.x >= len(input[p.y]) {
		return false
	}
	return true
}

func readInput(fileName string) ([][]rune, error) {
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
