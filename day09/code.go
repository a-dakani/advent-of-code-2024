package day09

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

var emptySpace = '.'

func solvePartOne() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	memory := createMemoryRepresentation(input)
	rearrangedMemory := rearrangeMemory(memory)
	return calculateChecksum(rearrangedMemory)
}

func calculateChecksum(memory []rune) int {
	checksum := 0
	for i, elem := range memory {
		if elem == emptySpace {
			continue
		}
		checksum += i * int(elem-'0')
	}
	return checksum
}

func rearrangeMemory(memory []rune) []rune {
	for i := len(memory) - 1; i >= 0; i-- {
		if isMemoryRearranged(memory) {
			break
		}
		if memory[i] == emptySpace {
			continue
		}
		for j, elm := range memory {
			if elm != emptySpace {
				continue
			}
			memory[i], memory[j] = memory[j], memory[i]
			break
		}
	}
	return memory
}

func isMemoryRearranged(memory []rune) bool {
	foundEmptySpace := false
	for _, elem := range memory {
		if elem == emptySpace {
			foundEmptySpace = true
		} else if foundEmptySpace {
			return false
		}
	}
	return true
}

func createMemoryRepresentation(input []int) []rune {
	var memoryMap []rune
	lastId := 0
	for i, num := range input {
		if i%2 == 0 {
			for j := 0; j < num; j++ {
				memoryMap = append(memoryMap, rune(lastId+'0'))
			}
			lastId++
		} else {
			for j := 0; j < num; j++ {
				memoryMap = append(memoryMap, emptySpace)
			}
		}
	}

	return memoryMap
}

func solvePartTwo() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	memory := createMemoryRepresentation(input)
	rearrangedMemory := rearrangeMemoryFiles(memory)
	return calculateChecksum(rearrangedMemory)
}

func rearrangeMemoryFiles(memory []rune) []rune {
	//	rearange complete blocks of memory (same rune) if it fits in the next possible empty space if not then skip
	for i := len(memory) - 1; i >= 0; i-- {
		//	find block of memory
		if memory[i] == emptySpace {
			continue
		}
		blockStart := i
		blockEnd := i
		for j := i - 1; j >= 0; j-- {
			if memory[j] == memory[i] {
				blockStart = j
			} else {
				break
			}
		}
		//	find adequate empty space before block starting from the first empty space
		for j := 0; j < blockStart; j++ {
			if memory[j] != emptySpace {
				continue
			}
			//	check if block fits in empty space
			emptySpaceStart := j
			emptySpaceEnd := j
			for k := j + 1; k < len(memory); k++ {
				if memory[k] == emptySpace {
					emptySpaceEnd = k
				} else {
					break
				}
			}
			if emptySpaceEnd-emptySpaceStart < blockEnd-blockStart {
				continue
			} else {
				//	rearrange memory
				for k := blockStart; k <= blockEnd; k++ {
					memory[k], memory[emptySpaceStart] = memory[emptySpaceStart], memory[k]
					emptySpaceStart++
				}
				break
			}
		}
		//	set i to the start of the block
		i = blockStart

	}
	return memory
}

func readInput(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		for _, char := range text {
			if unicode.IsDigit(char) {
				data = append(data, int(char-'0'))
			}
		}
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
