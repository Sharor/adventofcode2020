package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	ins := strings.Split(strings.TrimSpace(string(input)), "\n")

	ship, newStart := image.Point{0, 0}, image.Point{10, -1}
	fmt.Println(setSails(ins, &ship, &image.Point{1, 0}, &ship))
	fmt.Println(setSails(ins, &image.Point{0, 0}, &newStart, &newStart))
}

func setSails(ins []string, ship, coordinates, mov *image.Point) int {
	directionLookup := map[rune]image.Point{'N': {0, -1}, 'S': {0, 1}, 'E': {1, 0}, 'W': {-1, 0}, 'L': {-1, 1}, 'R': {1, -1}}
	for _, s := range ins {
		var direction rune
		var value int
		fmt.Sscanf(s, "%c%d", &direction, &value)

		switch direction {
		case 'N', 'S', 'E', 'W':
			*mov = mov.Add(directionLookup[direction].Mul(value))
		case 'L', 'R':
			coordinates.X, coordinates.Y = turnShip(value/90, directionLookup, direction, coordinates)
		case 'F':
			*ship = ship.Add(coordinates.Mul(value))
		}
	}
	return int(math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y)))
}

func turnShip(steps int, directionLookup map[rune]image.Point, instruction rune, coordinate *image.Point) (int, int) {
	for i := 0; i < steps; i++ {
		coordinate.X, coordinate.Y = directionLookup[instruction].Y*coordinate.Y, directionLookup[instruction].X*coordinate.X
	}
	return coordinate.X, coordinate.Y
}
