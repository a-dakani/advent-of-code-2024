package day13

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coordinates struct {
	x, y int
}

type pressedButtons struct {
	a, b int
}

type clawMachine struct {
	buttonA, buttonB, prize coordinates
}

func solvePartOne() int {
	clawMachines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	totalCost := 0

	for _, machine := range clawMachines {
		minCost, reachable := getMinCostToReachTarget(machine.buttonA, machine.buttonB, machine.prize, 100)
		if reachable {
			totalCost += minCost
			continue
		}
	}

	return totalCost
}

func getMinCostToReachTarget(a, b, target coordinates, maxPresses int) (int, bool) {
	var possibleCombinations []pressedButtons

	for pressedA := range maxPresses {
		for pressedB := range maxPresses {
			if a.x*pressedA+b.x*pressedB == target.x && a.y*pressedA+b.y*pressedB == target.y {
				possibleCombinations = append(possibleCombinations, pressedButtons{pressedA, pressedB})
			}
		}
	}

	if len(possibleCombinations) == 0 {
		return 0, false
	}

	tokensForA := 3
	tokensForB := 1

	minCost := math.MaxInt

	for _, comb := range possibleCombinations {
		cost := comb.a*tokensForA + comb.b*tokensForB
		if cost < minCost {
			minCost = cost
		}
	}

	return minCost, true
}

func solvePartTwo() int {
	clawMachines, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	totalCost := 0

	for _, machine := range clawMachines {
		machine.prize.x += 10000000000000
		machine.prize.y += 10000000000000
		minCost, reachable := getMinCostToReachTargetWithAlgebra(machine.buttonA, machine.buttonB, machine.prize)
		if reachable {
			totalCost += minCost
			continue
		}
	}

	return totalCost
}

func getMinCostToReachTargetWithAlgebra(a, b, prize coordinates) (int, bool) {
	// we can not brute force this problem, we need to use math to solve it

	// the Problem can be represented as algebraic equations
	// a.x * pressedA + b.x * pressedB = prize.x
	// a.y * pressedA + b.y * pressedB = prize.y

	// we can represent these equations as a matrix
	// | a.x b.x | | pressedA | = | prize.x |
	// | a.y b.y | | pressedB | = | prize.y |

	//pressedButtons=buttonMatrix^(-1) * prizeValues

	// Calculate the determinant of the button matrix
	det := a.x*b.y - a.y*b.x

	// if det == 0, the system of equations has no solution
	if det == 0 {
		return 0, false
	}

	// we only want integer solutions
	if (prize.x*b.y-prize.y*b.x)%det != 0 || (prize.y*a.x-prize.x*a.y)%det != 0 {
		return 0, false
	}

	// Calculate the values of pressedA and pressedB
	pressedA := (prize.x*b.y - prize.y*b.x) / det
	pressedB := (prize.y*a.x - prize.x*a.y) / det

	// we only want all positive solutions
	if pressedA < 0 || pressedB < 0 {
		return 0, false
	}

	// Calculate the total cost
	tokensForA := 3
	tokensForB := 1
	totalCost := pressedA*tokensForA + pressedB*tokensForB

	return totalCost, true
}

func readInput(fileName string) ([]clawMachine, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var machines []clawMachine
	var machine clawMachine
	lineNumber := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch lineNumber % 4 {
		case 0:
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &machine.buttonA.x, &machine.buttonA.y)
		case 1:
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &machine.buttonB.x, &machine.buttonB.y)
		case 2:
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &machine.prize.x, &machine.prize.y)
			machines = append(machines, machine)

		}
		lineNumber++
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return machines, nil
}
