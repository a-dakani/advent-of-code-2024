package day05

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func solvePartOne() int {
	rules, data, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	correctScore := 0
	for _, datum := range data {
		if isDatumValid(datum, rules) {
			middleNumber := datum[len(datum)/2]
			correctScore = correctScore + middleNumber
		}
	}

	return correctScore
}

func solvePartTwo() int {
	rules, data, err := readTextLines("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	correctedScore := 0
	var invalidData [][]int
	for _, datum := range data {
		if !isDatumValid(datum, rules) {
			invalidData = append(invalidData, datum)
		}
	}

	for _, invalidDatum := range invalidData {
		correctedDatum := correctDatum(invalidDatum, rules)
		middleNumber := correctedDatum[len(correctedDatum)/2]
		correctedScore = correctedScore + middleNumber
	}

	return correctedScore
}

func isDatumValid(datum []int, rules [][]int) bool {
	for k, num := range datum {
		numberOrderValid := IsNumValidInList(num, k, datum, rules)
		lastIteration := k == len(datum)-1
		if numberOrderValid && lastIteration {
			continue
		} else if numberOrderValid {
			continue
		} else {
			return false
		}
	}
	return true
}

func correctDatum(datum []int, rules [][]int) []int {
	for !isDatumValid(datum, rules) {
		for _, rule := range rules {
			indexOfFirstHand := slices.Index(datum, rule[0])
			indexOfSecondHand := slices.Index(datum, rule[1])
			if indexOfFirstHand < 0 || indexOfSecondHand < 0 {
				continue
			}
			if indexOfFirstHand > indexOfSecondHand {
				swap(datum, indexOfFirstHand, indexOfSecondHand)
			}
		}
	}
	return datum
}

func swap(list []int, i, j int) {
	if i < 0 || i >= len(list) || j < 0 || j >= len(list) {
		panic("index out of range")
	}
	list[i], list[j] = list[j], list[i]
}

func IsNumValidInList(num int, numIndex int, datum []int, rules [][]int) bool {
	for _, rule := range rules {
		if rule[0] == num {
			//	num must be before the second hand of rule
			for i := numIndex - 1; i >= 0; i-- {

				if datum[i] == rule[1] {
					return false
				}
			}
		}
		if rule[1] == num {
			//	num must be after the first hand of rule
			for i := numIndex + 1; i < len(datum); i++ {
				if datum[i] == rule[0] {
					return false
				}
			}
		}
	}
	return true
}

func readTextLines(fileName string) ([][]int, [][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var rules [][]int
	var data [][]int
	reachedEndOfRules := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			reachedEndOfRules = true
			continue
		}

		if !reachedEndOfRules {
			line := scanner.Text()
			rulesPair := strings.Split(line, "|")
			rules = append(rules, convertListOfStringToInt(rulesPair))
		} else {
			line := scanner.Text()
			datum := strings.Split(line, ",")
			data = append(data, convertListOfStringToInt(datum))
		}

	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, nil, scannerErr
	}

	return rules, data, nil
}

func convertListOfStringToInt(stringList []string) []int {
	var newList []int
	for _, s := range stringList {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic("found non-numeral value in list")
		}
		newList = append(newList, num)
	}
	return newList
}
