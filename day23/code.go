package day23

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func solvePartOne() int {
	connections, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	network := mapNetwork(connections)
	setsOfThree := getSetsOfConnections(network, 3)
	chiefHistorianConnections := 0

	for _, v := range setsOfThree {
		for _, c := range v {
			if strings.HasPrefix(c, "t") {
				chiefHistorianConnections++
				break
			}
		}
	}

	return chiefHistorianConnections
}

func solvePartTwo() string {
	connections, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	network := mapNetwork(connections)
	largestSetOfConnections := getLargestSetOfConnections(network)

	outputComaSeperated := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(largestSetOfConnections)), ","), "[]")

	return outputComaSeperated
}

func getLargestSetOfConnections(network map[string][]string) []string {
	var largestSet []string

	for k, v := range network {
		if len(v) == 0 {
			continue
		}

		combinations := generateCombinations(v, len(v)-1)
		for _, combo := range combinations {
			set := append([]string{k}, combo...)
			slices.Sort(set)

			if areConnectionsMutual(network, set) && len(set) > len(largestSet) {
				largestSet = set
			}
		}
	}

	return largestSet
}

func getSetsOfConnections(network map[string][]string, conSize int) map[string][]string {

	// conSize is the size of the subnetwork we are looking for
	// our network mapping is a map of computer to its connections
	// so we want to find computers that has conSize - 1 connections
	// Example: if conSize is 3, we want to find computers with at least 2 connections
	connectedComputerNum := conSize - 1

	sets := make(map[string][]string)

	for k, v := range network {
		if len(v) < connectedComputerNum {
			continue
		}

		combinations := generateCombinations(v, connectedComputerNum)
		for _, combo := range combinations {
			set := append([]string{k}, combo...)
			slices.Sort(set)
			key := fmt.Sprintf("%v", set)
			if _, ok := sets[key]; ok {
				continue
			}
			if !areConnectionsMutual(network, set) {
				continue
			}
			sets[key] = set
		}
	}

	return sets
}

func areConnectionsMutual(network map[string][]string, set []string) bool {
	for i := 0; i < len(set); i++ {
		for j := i + 1; j < len(set); j++ {
			if !slices.Contains(network[set[i]], set[j]) {
				return false
			}
		}
	}
	return true
}

func generateCombinations(elements []string, size int) [][]string {
	if size == 0 {
		return [][]string{{}}
	}
	if len(elements) == 0 {
		return nil
	}

	var combinations [][]string
	for i := 0; i <= len(elements)-size; i++ {
		for _, combo := range generateCombinations(elements[i+1:], size-1) {
			combinations = append(combinations, append([]string{elements[i]}, combo...))
		}
	}

	return combinations
}

func mapNetwork(connections [][]string) map[string][]string {
	network := make(map[string][]string)

	for _, connection := range connections {
		a, b := connection[0], connection[1]
		if !slices.Contains(network[a], b) {
			network[a] = append(network[a], b)
		}
		if !slices.Contains(network[b], a) {
			network[b] = append(network[b], a)
		}
	}
	return network
}

func readInput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a, b string
		line := scanner.Text()
		fmt.Sscanf(line, "%2s-%2s", &a, &b)
		data = append(data, []string{a, b})
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}
