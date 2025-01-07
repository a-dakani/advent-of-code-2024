package day24

import (
	"bufio"
	"fmt"
	"log"
	"maps"
	"os"
	"sort"
	"strings"
)

type OP int

// Operation codes
const (
	AND OP = iota
	OR
	XOR
)

type Operation struct {
	Left   string
	Right  string
	Result string
	OpCode OP
}

func solvePartOne() int {
	wires, operations, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	sortedOperations := sortOperations(wires, operations)

	for _, op := range sortedOperations {
		left := wires[op.Left]
		right := wires[op.Right]
		wires[op.Result] = doOp(op, left, right)
	}
	sortedZVariables := getSortedZValues(wires)

	return generateBinary(sortedZVariables)
}

func generateBinary(values []bool) int {
	var binaryStr string
	for i := len(values) - 1; i >= 0; i-- {
		v := values[i]
		if v {
			binaryStr += "1"
		} else {
			binaryStr += "0"
		}
	}

	var result int
	fmt.Sscanf(binaryStr, "%b", &result)
	return result

}

func getSortedZValues(variables map[string]bool) []bool {
	zVariables := make(map[string]bool)
	for k, v := range variables {
		if strings.HasPrefix(k, "z") {
			zVariables[k] = v
		}
	}

	keys := make([]string, 0, len(zVariables))
	for key := range zVariables {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	sortedValues := make([]bool, 0, len(keys))
	for _, key := range keys {
		sortedValues = append(sortedValues, zVariables[key])
	}

	return sortedValues
}

func sortOperations(knownVariables map[string]bool, operations []Operation) []Operation {
	knownTemp := maps.Clone(knownVariables)
	sortedOperations := make([]Operation, 0, len(operations))

	notSorted := true

	for notSorted {
		foundNotKnown := false
		for _, op := range operations {
			if _, okLeft := knownTemp[op.Left]; okLeft {
				if _, okRight := knownTemp[op.Right]; okRight {
					sortedOperations = append(sortedOperations, op)
					knownTemp[op.Result] = true
				} else {
					foundNotKnown = true
				}
			} else {
				foundNotKnown = true
			}
		}

		if !foundNotKnown {
			notSorted = false
		}
	}

	return sortedOperations

}

func doOp(op Operation, left, right bool) bool {
	switch op.OpCode {
	case AND:
		return left && right
	case OR:
		return left || right
	case XOR:
		return left != right
	default:
		log.Fatalf("unknown operation: %v", op.OpCode)
	}

	return false
}

func solvePartTwo() string {
	_, _, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	return "not solved yet"
}

func readInput(fileName string) (map[string]bool, []Operation, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	variables := make(map[string]bool)
	var operations []Operation
	readingVariables := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readingVariables = false
			continue
		}

		if readingVariables {
			var name string
			var value bool
			parts := strings.Split(line, ": ")
			fmt.Sscanf(parts[0], "%s", &name)
			fmt.Sscanf(parts[1], "%t", &value)

			variables[name] = value

		} else {
			var op Operation
			var opStr string
			var left, right, result string
			fmt.Sscanf(line, "%s %s %s -> %s", &left, &opStr, &right, &result)

			switch opStr {
			case "AND":
				op.OpCode = AND
			case "OR":
				op.OpCode = OR
			case "XOR":
				op.OpCode = XOR
			default:
				log.Fatalf("unknown operation: %s", opStr)
			}

			op.Left = left
			op.Right = right
			op.Result = result
			operations = append(operations, op)
		}
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, nil, scannerErr
	}

	return variables, operations, nil
}
