package day16

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

var nextStepsCache = make(map[step][]step)

const (
	start = 'S'
	end   = 'E'
	wall  = '#'
)

type direction coordinates

var (
	up    = direction{0, -1}
	down  = direction{0, 1}
	right = direction{1, 0}
	left  = direction{-1, 0}
)

type coordinates struct {
	x, y int
}

type step struct {
	coordinates
	direction
}
type route struct {
	path  []step
	turns int
	score int
}

func solvePartOne() int {
	mapInput, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	startLocations := findRune(mapInput, start)
	if len(startLocations) != 1 {
		log.Fatalf("unexpected number of start coordinates in the map")
	}

	startCoordinates := startLocations[0]

	// we start by defining the first route that we know of

	firstStep := step{
		coordinates: startCoordinates,
		direction:   right,
	}
	firstRoute := route{
		path:  []step{firstStep},
		turns: 0,
		score: 0,
	}

	cheapestRoutes := findCheapestRoutes(mapInput, firstRoute)

	//	we need to sort the slice based on score and then return the first element

	return cheapestRoutes[0].score
}

func findCheapestRoutes(input [][]rune, initialRoute route) []route {

	scoreCache := make(map[step]int)
	var possibleRoutes []route
	queue := []route{initialRoute}

	for len(queue) > 0 {

		r := queue[0]
		queue = queue[1:]

		lastLocation := r.path[len(r.path)-1]

		if input[lastLocation.coordinates.y][lastLocation.coordinates.x] == end {
			//r.path = r.path[:len(r.path)-1]
			possibleRoutes = append(possibleRoutes, r)
			continue
		}

		// if we have already visited this location with a lower score we skip it
		if score, ok := scoreCache[lastLocation]; ok && score < r.score {
			continue
		}

		nextSteps := getPossibleNextSteps(input, lastLocation)
		if len(nextSteps) == 0 {
			continue
		}

		for _, next := range nextSteps {
			if isStepInPath(next, r.path) {
				continue
			}

			newPath := make([]step, len(r.path))
			copy(newPath, r.path)
			newPath = append(newPath, next)

			newRoute := route{
				newPath,
				r.turns,
				r.score,
			}

			if next.direction != lastLocation.direction {
				newRoute.turns++
				newRoute.score += 1001
			} else {
				newRoute.score++
			}

			queue = append(queue, newRoute)

			// we cache the score of the last location to avoid visiting it again with a higher score
			scoreCache[lastLocation] = newRoute.score
		}
	}

	var cheapestRoutes []route

	slices.SortFunc(possibleRoutes, func(a, b route) int {
		if a.score < b.score {
			return -1
		}
		if a.score > b.score {
			return 1
		}
		return 0
	})

	minScore := possibleRoutes[0].score
	for _, r := range possibleRoutes {
		if r.score == minScore {
			cheapestRoutes = append(cheapestRoutes, r)
		}
	}

	return cheapestRoutes
}

func isStepInPath(next step, path []step) bool {
	for _, s := range path {
		if s.coordinates == next.coordinates {
			return true
		}
	}
	return false
}

func getPossibleNextSteps(input [][]rune, s step) []step {
	if nextSteps, ok := nextStepsCache[s]; ok {
		return nextSteps
	}

	// possible steps are the same direction coming from or 90 degrees from the current step
	var nextSteps []step
	allDirections := []direction{up, down, left, right}

	comingDirection := s.direction

	// we remove the opisite direction we are coming from to avoid going back to the previous step
	// the opisite direction is identical in absolute value but with opposite sign
	// this would work for the first step because it is not initialized with a direction (0,0)

	for _, d := range allDirections {
		next := coordinates{s.coordinates.x + d.x, s.coordinates.y + d.y}
		nextStep := step{
			coordinates: next,
			direction:   d,
		}
		if !oppositeDirections(comingDirection, d) &&
			isOnMap(next, input) &&
			input[next.y][next.x] != wall {
			nextSteps = append(nextSteps, nextStep)
		}
	}

	nextStepsCache[s] = nextSteps

	return nextSteps
}

func oppositeDirections(d1, d2 direction) bool {
	return d1.x == -d2.x && d1.y == -d2.y

}

func isOnMap(p coordinates, input [][]rune) bool {
	if p.x < 0 || p.y < 0 {
		return false
	}
	if p.y >= len(input) || p.x >= len(input[p.y]) {
		return false
	}
	return true
}

func solvePartTwo() int {
	mapInput, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}

	startLocations := findRune(mapInput, start)
	if len(startLocations) != 1 {
		log.Fatalf("unexpected number of start coordinates in the map")
	}

	startCoordinates := startLocations[0]

	// we start by defining the first route that we know of

	firstStep := step{
		coordinates: startCoordinates,
		direction:   right,
	}
	firstRoute := route{
		path:  []step{firstStep},
		turns: 0,
		score: 0,
	}

	cheapestRoutes := findCheapestRoutes(mapInput, firstRoute)

	uniqueLocations := make(map[coordinates]struct{})
	for _, r := range cheapestRoutes {
		for _, s := range r.path {
			uniqueLocations[s.coordinates] = struct{}{}
		}
	}

	return len(uniqueLocations)
}

func findRune(data [][]rune, searchable rune) []coordinates {
	var found []coordinates
	for y, line := range data {
		for x, r := range line {
			if r == searchable {
				found = append(found, coordinates{x, y})
			}
		}
	}
	return found
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

func printMapWithRoute(input [][]rune, r route) {
	// create a deep copy of the input map
	inputCopy := make([][]rune, len(input))
	for i := range input {
		inputCopy[i] = make([]rune, len(input[i]))
		copy(inputCopy[i], input[i])
	}

	// replace every location on the map with an arrow depending on the direction of the step in the route
	for _, s := range r.path {
		inputCopy[s.coordinates.y][s.coordinates.x] = directionToArrow(s.direction, inputCopy[s.coordinates.y][s.coordinates.x])
	}
	for _, line := range inputCopy {
		fmt.Println(string(line))
	}

	fmt.Println("-----------------")
}

func directionToArrow(d direction, current rune) rune {

	// if the current rune is the start or the end we keep it as is
	switch d {
	case up:
		return '^'
	case down:
		return 'v'
	case left:
		return '<'
	case right:
		return '>'
	}

	if current == start || current == end {
		return current
	}

	return ' '
}
