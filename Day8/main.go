package main

import (
	"bufio"
	"log"
	"os"
	sc "strconv"
	s "strings"
)

type instruction struct {
	instructionType string
	value           int
	positive        bool
}

func main() {
	println(fixBrokenInstruction(readInstructions(readInput())))
}

func runInstructions(instructions []instruction) (bool, int) {
	accellorationMeter := 0
	var observedSteps []int
	i := 0
	for true {
		if len(instructions) <= i {
			return true, accellorationMeter
		}
		if contains(observedSteps, i) {
			return false, accellorationMeter
		} else {
			observedSteps = appendIfMissing(observedSteps, i)
		}
		switch instructions[i].instructionType {
		case "nop":
			i++
		case "acc":
			if instructions[i].positive {
				accellorationMeter += instructions[i].value
				i++
			} else {
				accellorationMeter -= instructions[i].value
				i++
			}
		case "jmp":
			if instructions[i].positive {
				i += instructions[i].value
			} else {
				i -= instructions[i].value
			}
		default:
			i++
		}
	}

	return true, accellorationMeter
}

func fixBrokenInstruction(instructions []instruction) int {

	for i := 0; i < len(instructions); i++ {
		if instructions[i].instructionType == "nop" || instructions[i].instructionType == "jmp" {
			tmp := uncorruptOperation(instructions, i)
			correct, accValue := runInstructions(tmp)
			if correct {
				return accValue
			}
		}
	}
	return 0
}

func uncorruptOperation(instructions []instruction, i int) []instruction {
	tmp := make([]instruction, len(instructions))
	copy(tmp, instructions)
	if tmp[i].instructionType == "nop" {
		tmp[i].instructionType = "jmp"
	} else {
		tmp[i].instructionType = "nop"
	}
	return tmp
}

func appendIfMissing(slice []int, i int) []int {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readInstructions(input []string) []instruction {
	var instructions []instruction
	for i := 0; i < len(input); i++ {
		commands := s.Split(input[i], " ")
		currInstruction := createInstruction(commands)
		instructions = append(instructions, currInstruction)
	}
	return instructions
}

func createInstruction(commands []string) instruction {
	positive := false
	if s.Contains(commands[1], "+") {
		positive = true
	}

	findValue, _ := sc.Atoi(string([]rune(commands[1])[1:]))

	return instruction{instructionType: commands[0], value: findValue, positive: positive}
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
