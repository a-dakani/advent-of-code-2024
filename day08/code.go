package day08

import (
	"bufio"
	"log"
	"os"
	"slices"
	"unicode"
)

type frequency struct {
	antennas []location
	freq     rune
}
type location struct {
	x, y int
}

func solvePartOne() int {
	input, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	frequencies := findFrequencies(input)

	antinodes := findAntinodeLocations(input, frequencies)
	uniqueAntinodes := findUniqueAntinodeLocations(antinodes)

	return len(uniqueAntinodes)
}

func findFrequencies(input [][]rune) []frequency {
	antennas := make(map[rune][]location)
	for y, line := range input {

		for x, position := range line {
			if unicode.IsDigit(position) || unicode.IsLetter(position) {
				if _, ok := antennas[position]; !ok {
					antennas[position] = []location{{x, y}}
				} else {
					antennas[position] = append(antennas[position], location{x, y})
				}
			}
		}
	}
	var frequencies []frequency

	for freq, locs := range antennas {
		frequencies = append(frequencies, frequency{antennas: locs, freq: freq})
	}
	return frequencies
}

func findUniqueAntinodeLocations(antinodes []location) []location {
	var unique []location
	for _, anti := range antinodes {
		if !slices.Contains(unique, anti) {
			unique = append(unique, anti)
		}
	}
	return unique
}

func findAntinodeLocations(input [][]rune, frequencies []frequency) []location {
	var antinodes []location
	for _, freq := range frequencies {
		if len(freq.antennas) < 2 {
			continue
		}
		for i, firstAntenna := range freq.antennas {
			for k := i + 1; k < len(freq.antennas); k++ {
				secondAntenna := freq.antennas[k]
				antinodes = append(antinodes, calculateValidAntinode(firstAntenna, secondAntenna, input)...)
			}
		}
	}
	return antinodes
}

func calculateValidAntinode(ant1 location, ant2 location, input [][]rune) []location {
	var validAntinodes []location

	xDiff := ant2.x - ant1.x
	yDiff := ant2.y - ant1.y

	firstAntinode := location{x: ant1.x - xDiff, y: ant1.y - yDiff}
	secondAntinode := location{x: ant2.x + xDiff, y: ant2.y + yDiff}

	if isOnMap(firstAntinode, input) {
		validAntinodes = append(validAntinodes, firstAntinode)
	}
	if isOnMap(secondAntinode, input) {
		validAntinodes = append(validAntinodes, secondAntinode)
	}
	return validAntinodes
}

func isOnMap(antinode location, input [][]rune) bool {
	if antinode.x < 0 || antinode.y < 0 {
		return false
	}
	if antinode.y >= len(input) || antinode.x >= len(input[antinode.y]) {
		return false
	}
	return true
}

func solvePartTwo() int {
	input, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	frequencies := findFrequencies(input)

	antinodes := findResonantAntinodeLocations(input, frequencies)
	uniqueAntinodes := findUniqueAntinodeLocations(antinodes)

	return len(uniqueAntinodes)
}

func findResonantAntinodeLocations(input [][]rune, frequencies []frequency) []location {
	var antinodes []location
	for _, freq := range frequencies {
		if len(freq.antennas) < 2 {
			continue
		}
		for i, firstAntenna := range freq.antennas {
			for k := i + 1; k < len(freq.antennas); k++ {
				secondAntenna := freq.antennas[k]
				antinodes = append(antinodes, calculateValidResonantAntinode(firstAntenna, secondAntenna, input)...)
			}
		}
	}
	return antinodes
}

func calculateValidResonantAntinode(ant1 location, ant2 location, input [][]rune) []location {
	var validAntinodes []location

	xDiff := ant2.x - ant1.x
	yDiff := ant2.y - ant1.y

	firstAntinode := ant1
	secondAntinode := ant2

	for isOnMap(firstAntinode, input) || isOnMap(secondAntinode, input) {
		if isOnMap(firstAntinode, input) {
			validAntinodes = append(validAntinodes, firstAntinode)
		}
		if isOnMap(secondAntinode, input) {
			validAntinodes = append(validAntinodes, secondAntinode)
		}
		firstAntinode = location{x: firstAntinode.x - xDiff, y: firstAntinode.y - yDiff}
		secondAntinode = location{x: secondAntinode.x + xDiff, y: secondAntinode.y + yDiff}
	}

	return validAntinodes
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
		data = append(data, []rune(scanner.Text()))
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
