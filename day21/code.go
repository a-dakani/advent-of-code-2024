package day21

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
)

type coordinates struct {
	x, y int
}

type move struct {
	from, to rune
}

var numPad = map[rune]coordinates{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'0': {1, 3},
	'A': {2, 3},
}

var dirPad = map[rune]coordinates{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
}

// Credit to William Y. Feng for the Cost Calculation Idea.
// https://www.youtube.com/watch?v=q5I6ZvJmHEo
// I tried DFS to generate all possible sequences but in Part 2 it was too slow.
func solvePartOne() int64 {
	codes, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	possibleNumPadMoves := calculateAllPossibleMoves(numPad)
	possibleDirPadMoves := calculateAllPossibleMoves(dirPad)
	cheapestDirMovesAfterDepth := generateAllMovesCostsAfterDepth(possibleDirPadMoves, 2)

	var complexity int64
	for _, code := range codes {
		minCost := int64(math.MaxInt64)
		firstSet := findPossibleSequencesForInput(code, possibleNumPadMoves)

		for _, f := range firstSet {
			cost := calculateMinCost(f, cheapestDirMovesAfterDepth)
			if cost < minCost {
				minCost = cost
			}
		}
		if minCost == math.MaxInt64 {
			log.Fatalf("couldn't find a valid sequence for code %v", string(code))
		}
		var strippedCode int64
		fmt.Sscanf(string(code), "%dA", &strippedCode)
		complexity += minCost * strippedCode
	}

	return complexity
}

func solvePartTwo() int64 {
	codes, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	possibleNumPadMoves := calculateAllPossibleMoves(numPad)
	possibleDirPadMoves := calculateAllPossibleMoves(dirPad)
	cheapestDirMovesAfterDepth := generateAllMovesCostsAfterDepth(possibleDirPadMoves, 25)

	var complexity int64
	for _, code := range codes {
		minCost := int64(math.MaxInt64)
		firstSet := findPossibleSequencesForInput(code, possibleNumPadMoves)

		for _, f := range firstSet {
			cost := calculateMinCost(f, cheapestDirMovesAfterDepth)
			if cost < minCost {
				minCost = cost
			}
		}
		if minCost == math.MaxInt64 {
			log.Fatalf("couldn't find a valid sequence for code %v", string(code))
		}
		var strippedCode int64
		fmt.Sscanf(string(code), "%dA", &strippedCode)
		complexity += minCost * strippedCode
	}

	return complexity
}

func calculateMinCost(code []rune, costs map[move]int64) int64 {
	var cost int64
	code = append([]rune{'A'}, code...)

	for i := 0; i < len(code)-1; i++ {
		cost += costs[move{code[i], code[i+1]}]
	}

	return cost
}

func generateAllMovesCostsAfterDepth(possibleMoves map[move][][]rune, depth int) map[move]int64 {
	cache := make(map[string]int64)
	costs := make(map[move]int64)
	for k := range possibleMoves {
		costs[k] = getCost(k, possibleMoves, cache, depth)
	}
	return costs
}

func getCost(k move, moves map[move][][]rune, cache map[string]int64, depth int) int64 {
	cacheKey := fmt.Sprintf("%c%c%d", k.from, k.to, depth)
	if val, found := cache[cacheKey]; found {
		return val
	}
	if depth == 1 {
		return int64(len(moves[k][0]))
	}

	bestCost := int64(math.MaxInt64)
	for _, seq := range moves[k] {
		seq = append([]rune{'A'}, seq...)
		var cost int64
		for i := 0; i < len(seq)-1; i++ {
			cost += getCost(move{seq[i], seq[i+1]}, moves, cache, depth-1)
		}
		if cost < bestCost {
			bestCost = cost
		}
	}

	cache[cacheKey] = bestCost
	return bestCost
}

func findPossibleSequencesForInput(code []rune, moves map[move][][]rune) [][]rune {
	code = append([]rune{'A'}, code...)
	var movePairs []move
	for i := 0; i < len(code)-1; i++ {
		movePairs = append(movePairs, move{code[i], code[i+1]})
	}
	return generateSequences(movePairs, moves, []rune{})
}

func generateSequences(movePairs []move, moves map[move][][]rune, currentSequence []rune) [][]rune {
	if len(movePairs) == 0 {
		return [][]rune{currentSequence}
	}

	var allSequences [][]rune
	currentMove := movePairs[0]
	remainingMoves := movePairs[1:]

	for _, moveSequence := range moves[currentMove] {
		newSequence := append(slices.Clone(currentSequence), moveSequence...)
		allSequences = append(allSequences, generateSequences(remainingMoves, moves, newSequence)...)
	}

	return allSequences
}

func calculateAllPossibleMoves(pad map[rune]coordinates) map[move][][]rune {
	possibleMoves := make(map[move][][]rune)
	for k1, v1 := range pad {
		for k2, v2 := range pad {
			m := move{k1, k2}
			sequences := findShortestSequencesBetweenTwoPoints(v1, v2, pad)
			possibleMoves[m] = sequences
		}
	}
	return possibleMoves
}

func findShortestSequencesBetweenTwoPoints(start, target coordinates, pad map[rune]coordinates) [][]rune {
	queue := [][]coordinates{{start}}
	visited := map[coordinates]int{start: 0}
	var allSequences [][]rune

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == target {
			allSequences = append(allSequences, pathToDirections(path))
			continue
		}

		for _, next := range findNeighbors(current, pad) {
			newDist := visited[current] + 1
			if dist, ok := visited[next]; !ok || newDist <= dist {
				visited[next] = newDist
				newPath := append(slices.Clone(path), next)
				queue = append(queue, newPath)
			}
		}
	}

	return allSequences
}

func findNeighbors(current coordinates, pad map[rune]coordinates) []coordinates {
	var neighbors []coordinates
	for _, neighbor := range pad {
		if neighbor == current {
			continue
		}
		if neighbor.x == current.x && math.Abs(float64(neighbor.y-current.y)) == 1 {
			neighbors = append(neighbors, neighbor)
		}
		if neighbor.y == current.y && math.Abs(float64(neighbor.x-current.x)) == 1 {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func pathToDirections(path []coordinates) []rune {
	var directions []rune
	for i := 0; i < len(path)-1; i++ {
		directions = append(directions, coordinatesToDirection(path[i], path[i+1]))
	}
	directions = append(directions, 'A')
	return directions
}

func coordinatesToDirection(from, to coordinates) rune {
	if math.Abs(float64(from.x+from.y)-math.Abs(float64(to.x+to.y))) != 1 {
		fmt.Printf("more than one step between coordinates %v and %v\n", from, to)
		panic("more than one step between coordinates")
	}
	if from.x == to.x {
		if from.y > to.y {
			return '^'
		} else {
			return 'v'
		}
	} else {
		if from.x > to.x {
			return '<'
		} else {
			return '>'
		}
	}
}

func readInput(fileName string) ([][]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var lineData []rune
		lineData = append(lineData, []rune(scanner.Text())...)
		data = append(data, lineData)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
