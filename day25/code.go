package day25

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func solvePartOne() int {
	fileContent, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	keys, locks := parseKeysAndLocks(fileContent)

	fittingPairs := findUniqueFittingPairs(keys, locks)

	return len(fittingPairs)
}

func findUniqueFittingPairs(keys [][]int, locks [][]int) map[string]struct{} {
	fittingPairs := make(map[string]struct{})

	for _, key := range keys {
		for _, lock := range locks {
			if fits(key, lock) {
				fittingPairs[fmt.Sprintf("%v-%v", key, lock)] = struct{}{}
			}
		}
	}

	return fittingPairs
}

func fits(key []int, lock []int) bool {
	for i, keyHeight := range key {
		if keyHeight+lock[i] > 5 {
			return false
		}
	}
	return true
}

func parseKeysAndLocks(fileContent []string) ([][]int, [][]int) {
	keys := make([][]int, 0)
	locks := make([][]int, 0)

	for i := 0; i < len(fileContent); i += 8 {

		schematic := fileContent[i : i+7]

		if schematic[0] == "#####" && schematic[6] == "....." {
			locks = append(locks, convertToHeights(schematic, true))
		}
		if schematic[0] == "....." && schematic[6] == "#####" {
			keys = append(keys, convertToHeights(schematic, false))
		}
	}

	return keys, locks
}

func convertToHeights(schematic []string, isLock bool) []int {
	heights := make([]int, len(schematic[0]))

	for col := 0; col < len(schematic[0]); col++ {
		height := -1
		if isLock {
			for row := 0; row < len(schematic); row++ {
				if schematic[row][col] == '#' {
					height++
				}
			}
		} else {
			for row := len(schematic) - 1; row >= 0; row-- {
				if schematic[row][col] == '#' {
					height++
				}
			}
		}
		heights[col] = height
	}

	return heights
}

func solvePartTwo() int {
	_, err := readFile("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	return 0
}

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var fileContent []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return fileContent, nil
}
