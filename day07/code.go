package day07

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Operator int

const (
	ADD Operator = iota
	MUL
	CON
)

func solvePartOne() int {

	input, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	totalCalculations := 0

	for _, equation := range input {
		wantedResult := equation[0]
		chain := equation[1:]
		if isReachable(wantedResult, chain, []Operator{ADD, MUL}) {
			totalCalculations += wantedResult
		}
	}
	return totalCalculations
}

func solvePartTwo() int {
	input, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	totalCalculations := 0

	for _, equation := range input {
		wantedResult := equation[0]
		chain := equation[1:]
		if isReachable(wantedResult, chain, []Operator{ADD, MUL, CON}) {
			totalCalculations += wantedResult
		}
	}
	return totalCalculations
}

func isReachable(wantedResult int, chain []int, ops []Operator) bool {
	if len(chain) == 1 {
		return wantedResult == chain[0]
	}

	if chain[0] > wantedResult {
		return false
	}

	for _, op := range ops {
		next := op.Apply(chain[0], chain[1])
		newChain := append([]int{next}, chain[2:]...)
		if isReachable(wantedResult, newChain, ops) {
			return true
		}
	}

	return false
}

func (op Operator) Apply(left, right int) int {
	switch op {
	case ADD:
		return left + right
	case MUL:
		return left * right
	case CON:
		concat, err := strconv.Atoi(strconv.Itoa(left) + strconv.Itoa(right))
		if err != nil {
			panic(fmt.Errorf("concating error %d , %d", left, right))
		}
		return concat
	default:
		panic(fmt.Errorf("invalid operator: %v", op))
	}
}

func readTextLines(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		equationInts := []int{}
		equationString := strings.Split(scanner.Text(), ": ")
		leftSide, err := strconv.Atoi(equationString[0])
		if err != nil {
			return nil, err
		}
		equationInts = append(equationInts, leftSide)

		rightSide := strings.Split(equationString[1], " ")
		for _, num := range rightSide {
			numInt, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			equationInts = append(equationInts, numInt)
		}

		data = append(data, equationInts)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
