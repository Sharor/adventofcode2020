package main

import (
	"fmt"
	i "image"
	"io/ioutil"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	seatingArrangement := map[i.Point]rune{}
	for yCoord, inputLines := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		for xCoord, val := range inputLines {
			seatingArrangement[i.Point{xCoord, yCoord}] = val
		}
	}

	fmt.Println(shiftSeating(seatingArrangement, 4, func(point, diff i.Point) i.Point { return point.Add(diff) }))
	fmt.Println(shiftSeating(seatingArrangement, 5, func(point, diff i.Point) i.Point {
		for seatingArrangement[point.Add(diff)] == '.' {
			point = point.Add(diff)
		}
		return point.Add(diff)
	}))
}

func shiftSeating(seats map[i.Point]rune, allowedSeatsAdjacent int, recursiveSight func(p, d i.Point) i.Point) (counter int) {
	moveInMatrix := []i.Point{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for diff := true; diff; {
		counter, diff = 0, false

		next := map[i.Point]rune{}
		for point, seat := range seats {
			sum := 0
			for _, d := range moveInMatrix {
				if seats[recursiveSight(point, d)] == '#' {
					sum++
				}
			}
			seat, counter = setSeat(seat, counter, sum, allowedSeatsAdjacent)
			next[point] = seat
			diff = isDifferent(diff, next, seats, point)
		}
		seats = next
	}
	return counter
}

func setSeat(seat rune, counter int, sum int, allowedSeatsAdjacent int) (rune, int) {
	if seat == '#' && sum >= allowedSeatsAdjacent {
		seat = 'L'
	} else if seat == 'L' && sum == 0 || seat == '#' {
		seat = '#'
		counter++
	}
	return seat, counter
}

func isDifferent(diff bool, next map[i.Point]rune, seats map[i.Point]rune, point i.Point) bool {
	return diff || next[point] != seats[point]
}
