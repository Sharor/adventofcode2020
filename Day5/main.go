package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	println(traverseRowsColumns(readInput()))
}

func traverseRowsColumns(instructions []string) int {
	rowLower := 0
	rowUpper := 127
	colLower := 0
	colUpper := 7

	var seatIds []int

	highest := 0
	for _, instruction := range instructions {
		for i := 0; i < len(instruction); i++ {
			currentInstruction := string([]rune(instruction)[i])
			if currentInstruction == "F" {
				rowLower, rowUpper = keepLowerHalf(rowLower, rowUpper)
			}

			if currentInstruction == "B" {
				rowLower, rowUpper = keepUpperHalf(rowLower, rowUpper)
			}

			if currentInstruction == "L" {
				colLower, colUpper = keepLowerHalf(colLower, colUpper)
			}

			if currentInstruction == "R" {
				colLower, colUpper = keepUpperHalf(colLower, colUpper)
			}
		}

		curr := rowLower*8 + colLower
		seatIds = append(seatIds, curr)
		rowLower = 0
		rowUpper = 127
		colLower = 0
		colUpper = 7
		if curr > highest {
			highest = curr
		}
	}

	for i := 0; i < 128*8; i++ {
		if contains(seatIds, i+1) && contains(seatIds, i-1) && !contains(seatIds, i) {
			return i
		}
	}

	return highest
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func keepLowerHalf(lower int, upper int) (int, int) {
	upper = (lower + upper) / 2
	return lower, upper
}

func keepUpperHalf(lower int, upper int) (int, int) {
	quotient, remainder := (lower+upper)/2, upper%2
	lower = quotient + remainder
	return lower, upper
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
