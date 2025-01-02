package day20

import (
	"bufio"
	"log"
	"math"
	"os"
	"slices"
)

type coordinates struct {
	x, y int
}
type cheat struct {
	start, end coordinates
}

var directions = []coordinates{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func solvePartOne() int {
	track, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	startPoint := findRune(track, 'S')
	targetPoint := findRune(track, 'E')

	path := findPathBFS(startPoint, targetPoint, track)
	cheats := findPossibleCheats(path, 2, 100)

	return len(cheats)
}

func solvePartTwo() int {
	track, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	startPoint := findRune(track, 'S')
	targetPoint := findRune(track, 'E')

	path := findPathBFS(startPoint, targetPoint, track)
	cheats := findPossibleCheats(path, 20, 100)

	return len(cheats)
}

func findPathBFS(start, target coordinates, track [][]rune) []coordinates {
	queue := [][]coordinates{{start}}
	visited := map[coordinates]bool{start: true}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == target {
			return path
		}

		for _, d := range directions {
			next := coordinates{current.x + d.x, current.y + d.y}
			if isInRangeAndNotWall(next, track) && !visited[next] {
				newPath := append(slices.Clone(path), next)
				queue = append(queue, newPath)
				visited[next] = true
			}
		}
	}

	return nil
}

func findPossibleCheats(path []coordinates, numberOfCheats, minSavings int) map[cheat]int {

	cheats := make(map[cheat]int)

	for i := 0; i < len(path)-1; i++ {
		for j := i + 2; j < len(path); j++ {
			s := path[i]
			c := path[j]

			manhattanDistance := calculateManhattanDistance(s, c)
			distanceOnOriginalPath := j - i
			saving := distanceOnOriginalPath - manhattanDistance

			if manhattanDistance <= numberOfCheats && saving >= minSavings {
				cheats[cheat{s, c}] = saving
			}
		}
	}

	return cheats
}

// Taxicab geometry https://en.wikipedia.org/wiki/Taxicab_geometry
func calculateManhattanDistance(p1, p2 coordinates) int {
	return int(math.Abs(float64(p2.x-p1.x)) + math.Abs(float64(p2.y-p1.y)))
}

func isWall(point coordinates, track [][]rune) bool {
	return track[point.y][point.x] == '#'
}

func isInRange(point coordinates, track [][]rune) bool {
	return point.y >= 0 && point.y < len(track) &&
		point.x >= 0 && point.x < len(track[0])
}

func isInRangeAndNotWall(point coordinates, track [][]rune) bool {
	return isInRange(point, track) && !isWall(point, track)
}

func findRune(data [][]rune, searchable rune) coordinates {
	for y, line := range data {
		for x, r := range line {
			if r == searchable {
				return coordinates{x, y}
			}
		}
	}
	return coordinates{}
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
		chars := []rune(scanner.Text())
		data = append(data, chars)
	}
	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return data, nil
}

func diff(a, b int) int {
	return int(math.Abs(float64(a - b)))
}
