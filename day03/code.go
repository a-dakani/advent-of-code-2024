package day03

import (
	"bufio"
	"log"
	"os"
	regexp "regexp"
	"strconv"
	"strings"
)

func solvePartOne() int {
	data, err := readCorruptedMemory("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	return calcMulOpsInString(data)
}

func calcMulOpsInString(data string) int {
	r := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := r.FindAllString(data, -1)

	var ops [][]int
	numRegex := regexp.MustCompile(`\d{1,3}`)

	for _, match := range matches {
		nums := numRegex.FindAllString(match, -1)
		if len(nums) == 2 {
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			ops = append(ops, []int{num1, num2})
		}
	}
	totalSumOfMulOps := 0
	for _, op := range ops {
		result := op[0] * op[1]
		totalSumOfMulOps = totalSumOfMulOps + result
	}

	return totalSumOfMulOps
}

func solvePartTwo() int {
	data, err := readCorruptedMemory("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	dontStatements := strings.Split(data, `don't()`)

	totalSumOfMulOps := calcMulOpsInString(dontStatements[0])

	for i, s := range dontStatements {
		if i == 0 {
			continue
		}
		if !strings.Contains(s, `do()`) {
			continue
		}

		doStatement := strings.SplitN(s, `do()`, 2)[1]
		totalSumOfMulOps = totalSumOfMulOps + calcMulOpsInString(doStatement)
	}
	return totalSumOfMulOps
}

func readCorruptedMemory(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = data + scanner.Text()
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return "", scannerErr
	}

	return data, nil
}
