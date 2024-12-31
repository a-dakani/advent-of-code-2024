package day19

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func solvePartOne() int {
	towels, designs, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	sharedCache := make(map[string]int)
	doableDesigns := 0

	for _, design := range designs {
		if possiblePatterns := countPossiblePatterns(towels, design, sharedCache); possiblePatterns > 0 {
			doableDesigns++
		}
	}

	return doableDesigns
}

func solvePartTwo() int {
	towels, designs, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	sharedCache := make(map[string]int)
	allPossibleDesigns := 0

	for _, design := range designs {
		allPossibleDesigns += countPossiblePatterns(towels, design, sharedCache)
	}

	return allPossibleDesigns
}

func countPossiblePatterns(towels []string, design string, cache map[string]int) int {
	if knownComp, ok := cache[design]; ok {
		return knownComp
	}

	if len(design) == 0 {
		return 1
	}

	var possibleComps int
	for _, towel := range towels {
		if remainingDesign, found := strings.CutPrefix(design, towel); found {
			possibleComps += countPossiblePatterns(towels, remainingDesign, cache)
		}
	}

	cache[design] = possibleComps
	return possibleComps
}

func readInput(fileName string) ([]string, []string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var towels []string
	var designs []string
	readingTowels := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			readingTowels = false
			continue
		}
		if readingTowels {
			towels = strings.Split(text, ", ")
			continue
		}
		designs = append(designs, text)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, nil, scannerErr
	}

	return towels, designs, nil
}
