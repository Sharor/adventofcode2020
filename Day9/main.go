package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

func main() {
	numbers := readInput()
	invalid := (checkSums(25, numbers))
	println(findContinousRangeSavingInvalid(invalid, numbers))
}

func findContinousRangeSavingInvalid(invalid int, numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		min := math.MaxUint32
		max := 0
		sum := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > invalid {
				break
			}
			min, max = setBothMinMax(min, max, numbers[i], numbers[j])
			if sum == invalid {
				return min + max
			}
		}
	}
	return 0
}

func setBothMinMax(min int, max int, numberI int, numberJ int) (int, int) {
	max = setMax(max, numberI)
	max = setMax(max, numberJ)
	min = setMin(min, numberI)
	min = setMin(min, numberJ)

	return min, max
}

func setMin(min int, number int) int {
	if min > number {
		min = number
	}
	return min
}

func setMax(max int, number int) int {
	if max < number {
		max = number
	}
	return max
}

func checkSums(preamble int, numbers []int) int {
	for i := preamble; i < len(numbers); i++ {
		found := false

		for j := 0; j < preamble; j++ {
			start := setPreamble(i, preamble)
			firstNumber := numbers[start+j]
			for k := 0; k < preamble; k++ {
				start := setPreamble(i, preamble)
				secondNumber := numbers[start+k+1]

				if firstNumber+secondNumber == numbers[i] && firstNumber != secondNumber {
					found = true
					break
				}
			}
			if found {
				break
			}
		}

		if !found {
			return numbers[i]
		}
	}
	return 0
}

func setPreamble(i int, preamble int) int {
	start := i - preamble
	if start <= 0 {
		start = 0
	}
	return start
}

func readInput() []int {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var perline int
	var nums []int
	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	return nums
}
