package day01

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solvePartOne() int {
	firstList, secondList := readInputFile("input.txt")

	slices.Sort(firstList)
	slices.Sort(secondList)

	if len(firstList) != len(secondList) {
		log.Fatal("lists are of different sizes")
	}

	totalDistance := 0
	for i, s := range firstList {
		totalDistance += diff(s, secondList[i])
	}

	return totalDistance
}

func solvePartTwo() int {
	similarityScore := 0
	firstList, secondList := readInputFile("input.txt")

	for _, s := range firstList {
		recur := 0
		for _, k := range secondList {
			if s == k {
				recur++
			}
		}
		similarityScore = similarityScore + (recur * s)
	}
	return similarityScore
}

func readInputFile(fileName string) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	defer file.Close()

	var firstList, secondList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		if len(numbers) != 2 {
			log.Fatalf("input has unexpected format: %s", line)
		}
		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])
		if err1 != nil || err2 != nil {
			log.Fatalf("input has unexpected data types: %s", line)
		}
		firstList = append(firstList, num1)
		secondList = append(secondList, num2)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	return firstList, secondList
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
