package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solvePartOne() int64 {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	iterations := 2000

	var secretSum int64
	for _, secretNumber := range input {
		secretNumber, _ = generateNextSecretNumber(secretNumber, []int64{secretNumber % 10}, iterations)
		secretSum += secretNumber
	}

	return secretSum
}

func solvePartTwo() int64 {
	input, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	iterations := 2000

	var allBananaCount [][]int64
	for _, secretNumber := range input {
		_, bananaCount := generateNextSecretNumber(secretNumber, []int64{secretNumber % 10}, iterations)
		allBananaCount = append(allBananaCount, bananaCount)
	}

	quartetChangesMaps := generateQuartetChangesMaps(allBananaCount)

	combinedQuartetMap := combineQuartetChangesMaps(quartetChangesMaps)

	var maxBananaCount int64
	for _, bananaCount := range combinedQuartetMap {
		if bananaCount > maxBananaCount {
			maxBananaCount = bananaCount
		}
	}
	return maxBananaCount
}
func combineQuartetChangesMaps(fourChangesMaps []map[string]int64) map[string]int64 {
	combinedMap := make(map[string]int64)
	for _, fourChangesMap := range fourChangesMaps {
		for key, val := range fourChangesMap {
			combinedMap[key] += val
		}
	}
	return combinedMap
}

func generateQuartetChangesMaps(allBananaCount [][]int64) []map[string]int64 {
	var quartetChangesMaps []map[string]int64
	for _, bananaCount := range allBananaCount {
		changeMap := make(map[string]int64)
		for i := 4; i < len(bananaCount); i++ {
			key := fmt.Sprintf("%d,%d,%d,%d",
				bananaCount[i-3]-bananaCount[i-4],
				bananaCount[i-2]-bananaCount[i-3],
				bananaCount[i-1]-bananaCount[i-2],
				bananaCount[i]-bananaCount[i-1])

			if _, ok := changeMap[key]; !ok {
				changeMap[key] = bananaCount[i]
			}
		}
		quartetChangesMaps = append(quartetChangesMaps, changeMap)
	}
	return quartetChangesMaps
}

func generateNextSecretNumber(secretNumber int64, bananaCount []int64, depth int) (int64, []int64) {
	for depth > 0 {
		// step 1
		secretNumber = pruneNumber(mixNumbers(secretNumber, secretNumber<<6))
		// step 2
		secretNumber = pruneNumber(mixNumbers(secretNumber, secretNumber>>5))
		// step 3
		secretNumber = pruneNumber(mixNumbers(secretNumber, secretNumber<<11))
		// add banana count (last digit of secret number)
		bananaCount = append(bananaCount, secretNumber%10)
		depth--
	}
	return secretNumber, bananaCount
}

func pruneNumber(a int64) int64 {
	return a % 16777216
}

func mixNumbers(a, b int64) int64 {
	return a ^ b
}

func readInput(fileName string) ([]int64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		parsedInt, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			return nil, err
		}
		data = append(data, parsedInt)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
