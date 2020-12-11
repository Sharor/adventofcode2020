package main

import (
	"bufio"
	"log"
	"os"
	sc "strconv"
	s "strings"
)

type password struct {
	requirements string
	letter       string
	password     string
}

func main() {
	strings := readInput()
	passwords := extractPasswords(strings)
	println(countValidPasswords(passwords))
}

func countValidPasswords(passwords []password) int {
	var validCount int
	for i := 0; i < len(passwords); i++ {
		min, _ := sc.Atoi(s.Split(passwords[i].requirements, "-")[0])
		max, _ := sc.Atoi(s.Split(passwords[i].requirements, "-")[1])
		letter := passwords[i].letter
		currPassword := passwords[i].password
		if passwordValid(min, max, letter, currPassword) {
			validCount++
		}
	}

	return validCount
}

func passwordValid(min int, max int, letter string, password string) bool {
	var countedLetter int

	for i, r := range password {
		if letter == string(r) && (i+1 == max || i+1 == min) {
			countedLetter++
		}
	}

	if countedLetter == 1 {
		return true
	}
	return false
}

func extractPasswords(allLines []string) []password {
	var processedIntoPasswords []password
	for i := 0; i < len(allLines); i++ {
		currentLine := allLines[i]
		password := processPassword(currentLine)
		processedIntoPasswords = append(processedIntoPasswords, password)
	}
	return processedIntoPasswords
}

func processPassword(line string) password {
	lineSplit := s.Split(line, ":")
	password := s.TrimSpace(lineSplit[1])
	requirements := s.TrimSpace(s.Split(lineSplit[0], " ")[0])
	letter := s.TrimSpace(s.Split(lineSplit[0], " ")[1])
	return createPassword(requirements, letter, password)
}

func createPassword(requirements string, letter string, passcode string) password {
	return password{requirements: requirements, letter: letter, password: passcode}
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
