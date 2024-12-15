package day03

import (
	"bufio"
	"log"
	"os"
	"reflect"
)

func solvePartOne() int {
	lines, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	forward := []rune("XMAS")
	backward := []rune("SAMX")

	totalXmas := 0

	for li, l := range lines {
		for ci, c := range l {
			if c == forward[0] {
				totalXmas = totalXmas + calculatePossibilities(forward, ci, li, lines)
			} else if c == backward[0] {
				totalXmas = totalXmas + calculatePossibilities(backward, ci, li, lines)
			}
		}

	}
	return totalXmas
}

func solvePartTwo() int {
	lines, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	forward := []rune("MAS")
	backward := []rune("SAM")

	totalXmas := 0

	for li, line := range lines {
		for ci, c := range line {
			if ci > len(line)-3 {
				continue
			}

			firstWordForward := 0
			secondWordForward := 0
			firstWordbackward := 0
			secondWordbackward := 0
			if c == forward[0] {
				firstWordForward = findWordDiagonalRight(forward, ci, li, lines)
			}
			if c == backward[0] {
				firstWordbackward = findWordDiagonalRight(backward, ci, li, lines)
			}
			if line[ci+2] == forward[0] {
				secondWordForward = findWordDiagonalLeft(forward, ci+2, li, lines)
			}
			if line[ci+2] == backward[0] {
				secondWordbackward = findWordDiagonalLeft(backward, ci+2, li, lines)
			}

			score := firstWordForward + firstWordbackward + secondWordForward + secondWordbackward
			if score >= 2 {
				totalXmas++
			}
		}

	}
	return totalXmas
}

func calculatePossibilities(word []rune, firstCharIndex int, lineIndex int, lines [][]rune) int {

	wordForward := findWordForward(word, firstCharIndex, lines[lineIndex])
	wordDownward := findWordDownward(word, firstCharIndex, lineIndex, lines)
	wordDiagonalRight := findWordDiagonalRight(word, firstCharIndex, lineIndex, lines)
	wordDiagonalLeft := findWordDiagonalLeft(word, firstCharIndex, lineIndex, lines)

	return wordForward + wordDownward + wordDiagonalRight + wordDiagonalLeft
}

func findWordDiagonalRight(word []rune, firstCharIndex int, lineIndex int, lines [][]rune) int {
	endWordIndex := firstCharIndex + len(word)

	if endWordIndex > len(lines[lineIndex]) {
		return 0
	}

	endLineIndex := lineIndex + len(word)
	if endLineIndex > len(lines) {
		return 0
	}

	wordLines := lines[lineIndex:endLineIndex]
	var maybeWord []rune

	diagonalIndex := firstCharIndex
	for _, line := range wordLines {
		maybeWord = append(maybeWord, line[diagonalIndex])
		diagonalIndex++
	}

	if reflect.DeepEqual(word, maybeWord) {
		return 1
	}

	return 0
}

func findWordDiagonalLeft(word []rune, firstCharIndex int, lineIndex int, lines [][]rune) int {
	endWordIndex := firstCharIndex - len(word) + 1
	if endWordIndex < 0 {
		return 0
	}

	endLineIndex := lineIndex + len(word)
	if endLineIndex > len(lines) {
		return 0
	}

	wordLines := lines[lineIndex:endLineIndex]
	var maybeWord []rune

	diagonalIndex := firstCharIndex
	for _, line := range wordLines {
		maybeWord = append(maybeWord, line[diagonalIndex])
		diagonalIndex--

	}

	if reflect.DeepEqual(word, maybeWord) {
		return 1
	}

	return 0
}

func findWordDownward(word []rune, firstCharIndex int, lineIndex int, lines [][]rune) int {
	endLineIndex := lineIndex + len(word)

	if endLineIndex > len(lines) {
		return 0
	}

	wordLines := lines[lineIndex:endLineIndex]
	var maybeWord []rune

	for _, line := range wordLines {
		maybeWord = append(maybeWord, line[firstCharIndex])
	}

	if reflect.DeepEqual(word, maybeWord) {
		return 1
	}

	return 0
}

func findWordForward(word []rune, firstCharIndex int, line []rune) int {
	endCharIndex := firstCharIndex + len(word)

	if endCharIndex > len(line) {
		return 0
	}

	maybeWord := line[firstCharIndex:endCharIndex]

	if reflect.DeepEqual(word, maybeWord) {

		return 1
	}

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
