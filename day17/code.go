package day17

import (
	"bufio"
	fmt "fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func solvePartOne() []int {
	a, b, c, program, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	output := calculateOutput(a, b, c, program)

	outputComaSeperated := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(output)), ","), "[]")

	fmt.Println("End of Program")
	fmt.Printf("Output comma seperated: %s\n", outputComaSeperated)

	return output
}

func calculateOutput(a, b, c int64, program []int) []int {

	combo := func(operand int) int64 {
		switch operand {
		case 0, 1, 2, 3:
			return int64(operand)
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		case 7:
			panic("invalid operand")
		default:
			panic("unknown operand")
		}
	}

	var output []int
	ip := 0

	for {

		if ip >= len(program) {
			break
		}

		opcode := program[ip]
		literalOperand := program[ip+1]

		switch opcode {
		case 0:
			a = dvProg(a, combo(literalOperand))
			ip += 2
		case 1:
			b = b ^ int64(literalOperand)
			ip += 2
		case 2:
			b = combo(literalOperand) % 8
			ip += 2
		case 3:
			if a > 0 {
				ip = literalOperand
			} else {
				ip += 2
			}
		case 4:
			b = b ^ c
			ip += 2
		case 5:
			output = append(output, int(combo(literalOperand)%8))
			ip += 2
		case 6:
			b = dvProg(a, combo(literalOperand))
			ip += 2
		case 7:
			c = dvProg(a, combo(literalOperand))
			ip += 2
		}

	}
	return output
}

func dvProg(register, combo int64) int64 {
	return int64(float64(register) / (math.Pow(2, float64(combo))))
}

func solvePartTwo() int64 {
	_, _, _, program, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	// credit to Wekoslav Stefanovski for explaining the problem
	// https://www.youtube.com/watch?v=QpvAyg1RIYI
	return calculateCorrectA(program, program, 0)

}

func calculateCorrectA(program []int, target []int, initialA int64) int64 {

	if len(target) == 0 {
		return initialA
	}
	for digit := int64(0); digit < 8; digit++ {
		// example initialA = 2 (010), digit = 5 (101)
		// target is a binary concatenation of initialA and digit
		// shifting initialA 3 positions returns (101000)
		// OR-ing (|) with digit (101) returns (101101)
		var a = (initialA << 3) | digit

		var b, c int64

		output := calculateOutput(a, b, c, program)

		if output[0] == target[len(target)-1] {
			var prev = calculateCorrectA(program, target[:len(target)-1], a)
			if prev > 0 {
				return prev
			}
		}
	}
	return 0
}

func readInput(fileName string) (int64, int64, int64, []int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, 0, 0, nil, err
	}
	defer file.Close()

	var a, b, c int64
	var program []int
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		switch lineNumber {
		case 0:
			fmt.Sscanf(line, "Register A: %d", &a)
		case 1:
			fmt.Sscanf(line, "Register B: %d", &b)
		case 2:
			fmt.Sscanf(line, "Register C: %d", &c)
		case 4:
			programString := strings.Split(line, " ")[1]
			programmNumsString := strings.Split(programString, ",")
			for _, numString := range programmNumsString {
				i, err := strconv.Atoi(numString)
				if err != nil {
					return 0, 0, 0, nil, err
				}
				program = append(program, i)
			}
		}
		lineNumber++

	}
	return a, b, c, program, nil
}
