package main

import (
	"bufio"
	"log"
	"os"
	s "strings"
)

func main() {
	goThroughQuestions(readInput())
}

func goThroughQuestions(answers []string) {
	sum := 0
	var currSheet []string
	for i := 0; i < len(answers); i++ {
		if isNew(answers[i]) {
			sum += uniqueAnswers(currSheet)
			currSheet = currSheet[:0]
		} else {
			currSheet = append(currSheet, answers[i])
		}
	}
	sum += uniqueAnswers(currSheet)
	println(sum)
}

func uniqueAnswers(groupAnswers []string) int {
	var seen []string
	for i := 0; i < len(groupAnswers); i++ {
		for j := 0; j < len(groupAnswers[i]); j++ {
			seen = appendIfMissing(seen, string([]rune(groupAnswers[i])[j]))
		}
	}
	//fmt.Printf("%v", seen)
	allCorrect := 0
	for i := 0; i < len(seen); i++ {
		found := 0
		for j := 0; j < len(groupAnswers); j++ {
			if s.Contains(groupAnswers[j], seen[i]) {
				found++
			}
		}
		if (found) == len(groupAnswers) {
			allCorrect++
		}
	}

	return allCorrect
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func isNew(lineBeingRead string) bool {
	if lineBeingRead == "" {
		return true
	}
	return false
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
