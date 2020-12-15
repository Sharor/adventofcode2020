package main

import (
	"fmt"
	"io/ioutil"
	sc "strconv"
	s "strings"
)

func main() {
	day1 := speakNumbers(2020)
	day2 := speakNumbers(30000000)
	fmt.Printf("Part 1: %v\n", day1)
	fmt.Printf("Part 2: %v", day2)
}

func speakNumbers(speakThisMany int) int {
	input, _ := ioutil.ReadFile("input")
	seenNumbers := map[int]int{}
	number := 0

	for i, inputNumber := range s.Split(s.TrimSpace(string(input)), ",") {
		number, _ = sc.Atoi(inputNumber)
		seenNumbers[number] = i + 1
	}

	for i := len(seenNumbers); i < speakThisMany; i++ {
		if newNumber, exist := seenNumbers[number]; exist {
			seenNumbers[number], number = i, i-newNumber
		} else {
			seenNumbers[number], number = i, 0
		}
	}
	return number
}
