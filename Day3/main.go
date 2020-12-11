package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	landscape := readInput()
	firstRoute := findTreesOnRoute(landscape, 1, 1)
	secondRoute := findTreesOnRoute(landscape, 3, 1)
	thirdRoute := findTreesOnRoute(landscape, 5, 1)
	fourthRoute := findTreesOnRoute(landscape, 7, 1)
	fifthRoute := findTreesOnRoute(landscape, 1, 2)

	println(firstRoute * secondRoute * thirdRoute * fourthRoute * fifthRoute)
}

func findTreesOnRoute(landscape []string, speedRight int, speedDown int) int {

	horizontalCounter := speedRight
	treesHit := 0

	if []rune(landscape[0])[0] == '#' {
		treesHit++
	}

	for verticalCounter := speedDown; verticalCounter < len(landscape); verticalCounter += speedDown {
		currentLaneInLandscape := landscape[verticalCounter]
		currentLaneInLandscape = repeatLandscape(currentLaneInLandscape, horizontalCounter)

		if string([]rune(currentLaneInLandscape)[horizontalCounter]) == "#" {
			treesHit++
		}
		horizontalCounter += speedRight
	}
	return treesHit
}

func repeatLandscape(lineInLandscape string, horizontalCounter int) string {
	if len(lineInLandscape) <= horizontalCounter {
		lineInLandscape += lineInLandscape
		return repeatLandscape(lineInLandscape, horizontalCounter)
	}
	return lineInLandscape
}

func readInput() []string {
	file, err := os.Open("input")

	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	return text
}
