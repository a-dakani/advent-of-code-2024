package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type coordinates struct {
	x, y int
}

type robot struct {
	position, velocity coordinates
}

func solvePartOne() int {
	robots, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	gridWidth := 101
	gridHeight := 103
	timeInSeconds := 100

	for range timeInSeconds {
		var newRobots []robot
		for _, r := range robots {
			newRobot := moveRobotOneSecond(r, gridWidth, gridHeight)
			newRobots = append(newRobots, newRobot)
		}
		robots = newRobots
	}

	return calculateSafetyFactor(robots, gridWidth, gridHeight)
}

func moveRobotOneSecond(r robot, width int, height int) robot {
	r.position.x = r.position.x + r.velocity.x
	r.position.y = r.position.y + r.velocity.y

	if r.position.x < 0 {
		r.position.x += width
	}
	if r.position.x >= width {
		r.position.x -= width
	}
	if r.position.y < 0 {
		r.position.y += height
	}
	if r.position.y >= height {
		r.position.y -= height
	}
	return r
}

func calculateSafetyFactor(robots []robot, width int, height int) int {
	northEastCount := 0
	northWestCount := 0
	southEastCount := 0
	southWestCount := 0
	for _, r := range robots {
		if r.position.x == width/2 || r.position.y == height/2 {
			continue
		}

		switch {
		case r.position.x < width/2 && r.position.y < height/2:
			northWestCount++
		case r.position.x > width/2 && r.position.y < height/2:
			northEastCount++
		case r.position.x < width/2 && r.position.y > height/2:
			southWestCount++
		case r.position.x > width/2 && r.position.y > height/2:
			southEastCount++
		default:
			log.Fatalf("robot is in an unexpected position: %v", r)
		}
	}

	return northEastCount * northWestCount * southEastCount * southWestCount
}

func solvePartTwo() int {
	robots, err := readInput("input.txt")
	if err != nil {
		log.Fatalf("couldn't read input file: %v", err)
	}
	gridWidth := 101
	gridHeight := 103
	timeInSeconds := 10000
	var christmasSecond int

	for s := range timeInSeconds {
		var newRobots []robot
		for _, r := range robots {
			newRobot := moveRobotOneSecond(r, gridWidth, gridHeight)
			newRobots = append(newRobots, newRobot)
		}
		robots = newRobots

		if treeStructureFound(robots, gridWidth, gridHeight) {
			christmasSecond = s + 1
			break
		}
	}

	return christmasSecond
}

func treeStructureFound(robots []robot, width, height int) bool {
	// if there is a tree there will be at least one line with 10% of the width of the grid full of robots
	// for visual inspection we replace all positions of robots with # and empty spaces with .
	// if we find a row with 10 # we print the grid

	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}

	for _, r := range robots {
		grid[r.position.y][r.position.x] += 1
	}

	gridAsString := ""

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] > 0 {
				gridAsString += "#"
			} else {
				gridAsString += "."
			}
		}
		gridAsString += "\n"
	}

	if strings.Contains(gridAsString, "##########") {
		fmt.Println(gridAsString)
		return true
	}

	return false
}

func readInput(fileName string) ([]robot, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var robots []robot

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var r robot
		_, err := fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &r.position.x, &r.position.y, &r.velocity.x, &r.velocity.y)
		if err != nil {
			return nil, err
		}
		robots = append(robots, r)
	}

	if scannerErr := scanner.Err(); scannerErr != nil {
		return nil, scannerErr
	}

	return robots, nil
}
