package main

import (
	"io/ioutil"
	"math"
	sc "strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	part1(lines)
	part2(lines)
}

func part2(lines []string) {
	cheatFirstCondition, step := 0, 1
	for i, bus := range strings.Split(lines[1], ",") {
		busID, err := sc.Atoi(bus) //ignores x, cos who cares about error
		if err != nil {
			continue
		}
		for (cheatFirstCondition+i)%busID != 0 {
			cheatFirstCondition += step
		}
		step *= busID

	}
	println(cheatFirstCondition)
}

func part1(lines []string) {
	lesser, _ := sc.Atoi(lines[0])

	currMin := math.MaxInt64
	minBus := 0
	for _, bus := range strings.Split(lines[1], ",") {
		busID, err := sc.Atoi(bus) //ignores x, cos who cares about error
		if err != nil {
			continue
		}
		minVal := highestVal(lesser, busID)
		if currMin > minVal {
			currMin = minVal
			minBus = busID
		}

	}
	println((currMin - lesser) * minBus)
}

func highestVal(earliest int, bustime int) int {
	return (earliest/bustime + 1) * bustime
}
