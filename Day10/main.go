package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input")
	split := strings.Split(strings.TrimSpace(string(input)), "\n")

	jolts := make([]int, len(split))
	for i, s := range split {
		jolts[i], _ = strconv.Atoi(s)
	}
	sort.Ints(jolts)
	jolts = sortInput(jolts)

	mapOfOptions := map[int]int{0: 1}

	for _, v := range jolts[1:] {
		mapOfOptions[v] = mapOfOptions[v-1] + mapOfOptions[v-2] + mapOfOptions[v-3]
	}
	//fmt.Println(diff[1] * diff[3])
	fmt.Println(mapOfOptions[jolts[len(jolts)-1]])
}

func day1() {
	println(findJump((sortInput(readInput()))))
}

var countedPaths map[int]int

func day2() {
	joltages := sortInput(readInput())
	road := make(map[int]bool)

	for _, joltage := range joltages {
		road[joltage] = true
	}

	road[joltages[len(joltages)-1]+3] = true

	println(countPathsFrom(0, road))
}

func countPathsFrom(start int, road map[int]bool) int {
	count, visited := countedPaths[start]

	if visited {
		return count
	}

	res := 0
	noCandidate := true

	for i := 1; i <= 3; i++ {
		candidate := start + i
		_, ok := road[candidate]
		if ok {
			res += countPathsFrom(candidate, road)
			noCandidate = false
		}
	}

	if noCandidate {
		res++
	}

	countedPaths[start] = res

	return res
}

func sortInput(input []int) []int {
	var newInput []int
	newInput = append(newInput, 0)
	sort.Ints(input[:])
	newInput = append(newInput, input...)
	newInput = append(newInput, newInput[len(newInput)-1]+3)
	return newInput
}

func findJumpOptions(adapters []int, currentVoltage, routeCounter int) int {
	if len(adapters) == 0 {
		routeCounter++
		return routeCounter
	}
	options, optionsAdapters := calcOptions(currentVoltage, adapters)
	for j := 0; j < len(options); j++ {
		routeCounter = findJumpOptions(optionsAdapters[j], options[j], routeCounter)
	}
	return routeCounter
}

func calcOptions(currentVoltage int, adapters []int) ([]int, [][]int) {
	var options []int
	var optionsAdapters [][]int
	for j := 0; j < len(adapters); j++ {
		if currentVoltage == adapters[j] {
			continue
		}
		jumpSize := jumpSize(currentVoltage, adapters[j])
		if jumpSize > 3 {
			break
		}
		options = append(options, adapters[j])
		optionsAdapters = append(optionsAdapters, adapters[j+1:(len(adapters))])
	}
	return options, optionsAdapters
}

func jumpSize(optionOne int, optionTwo int) int {
	return optionTwo - optionOne
}

func findJump(adapters []int) int {
	var jumps []int

	for i := 0; i < len(adapters)-1; i++ {
		jumpSize := adapters[i+1] - adapters[i]
		jumps = append(jumps, jumpSize)
	}
	return findPower(jumps)
}

func findPower(jumps []int) int {
	countOnes := 0
	countThrees := 0

	for i := 0; i < len(jumps); i++ {
		if jumps[i] == 1 {
			countOnes++
		}
		if jumps[i] == 3 {
			countThrees++
		}
	}
	return countOnes * countThrees
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
