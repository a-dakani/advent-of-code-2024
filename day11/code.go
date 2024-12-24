package day11

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePartOne() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	for range 25 {
		var newStones []int
		for _, stone := range input {
			if stone == 0 {
				newStones = append(newStones, 1)
				continue
			}

			if len(strconv.Itoa(stone))%2 == 0 {
				stoneStr := strconv.Itoa(stone)
				half := len(stoneStr) / 2
				left, _ := strconv.Atoi(stoneStr[:half])
				right, _ := strconv.Atoi(stoneStr[half:])
				newStones = append(newStones, left, right)
				continue
			}
			newStones = append(newStones, stone*2024)
		}
		input = newStones
	}

	return len(input)
}

func solvePartTwo() int {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	stones := make(map[int]int)
	for _, stone := range input {
		stones[stone]++
	}

	for range 75 {
		newStones := make(map[int]int)
		for stone, count := range stones {
			if stone == 0 {
				newStones[1] += count
			} else if len(strconv.Itoa(stone))%2 == 0 {
				stoneStr := strconv.Itoa(stone)
				half := len(stoneStr) / 2
				left, _ := strconv.Atoi(stoneStr[:half])
				right, _ := strconv.Atoi(stoneStr[half:])
				newStones[left] += count
				newStones[right] += count
			} else {
				newStones[stone*2024] += count
			}
		}
		stones = newStones
	}

	totalCount := 0
	for _, count := range stones {
		totalCount += count
	}

	return totalCount
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
		numbers := strings.Split(text, " ")
		for _, number := range numbers {
			num, err := strconv.Atoi(number)
			if err != nil {
				return nil, err
			}
			data = append(data, num)
		}
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
