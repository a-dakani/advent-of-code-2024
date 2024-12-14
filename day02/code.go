package day02

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func solvePartOne() int {
	reportList, err := readReportList("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	validReportCount := 0
	for _, report := range reportList {
		if isReportFullyValid(report) {
			validReportCount++
		}
	}
	return validReportCount
}

func isReportFullyValid(report []int) bool {
	if len(report) <= 2 {
		return true
	}
	reportAscending := report[1] > report[0]

	for i := 1; i < len(report); i++ {
		diff := calcDiff(report[i], report[i-1])
		if diff < 1 || diff > 3 {
			return false
		}

		if reportAscending && report[i] < report[i-1] {
			return false
		}
		if !reportAscending && report[i] > report[i-1] {
			return false
		}
	}
	return true
}

func solvePartTwo() int {
	reportList, err := readReportList("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	partiallyValidReportCount := 0
	for _, report := range reportList {
		if isReportPartiallyValid(report) {
			partiallyValidReportCount++
		}
	}
	return partiallyValidReportCount
}

func isReportPartiallyValid(report []int) bool {
	if isReportFullyValid(report) {
		return true
	}

	for i := range report {
		var reportWithoutIndex []int
		for k := range report {
			if k == i {
				continue
			}
			reportWithoutIndex = append(reportWithoutIndex, report[k])
		}

		if isReportFullyValid(reportWithoutIndex) {
			return true
		}
	}

	return false
}

func readReportList(fileName string) ([][]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reportList [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var report []int
		line := scanner.Text()
		numbers := strings.Fields(line)

		for _, n := range numbers {
			num, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			report = append(report, num)
		}
		reportList = append(reportList, report)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return reportList, nil
}

func calcDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
