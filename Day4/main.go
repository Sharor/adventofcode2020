package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	sc "strconv"
	s "strings"
)

type passport struct {
	birthYear      string
	issueYear      string
	expirationYear string
	height         string
	hairColor      string
	eyeColor       string
	passportID     string
	countryID      string
}

func main() {
	passports := splitIntoPassports(readInput())
	println(countValidPassports(passports))
}

func countValidPassports(passports []passport) int {
	validPassports := 0

	for _, passport := range passports {
		if passportIsValid(passport) {
			validPassports++
		}
	}

	return validPassports
}

func passportIsValid(passport passport) bool {
	if passport.birthYear == "" || passport.issueYear == "" || passport.expirationYear == "" || passport.height == "" {
		return false
	}
	if passport.hairColor == "" || passport.eyeColor == "" || passport.passportID == "" {
		return false
	}

	yearsValid := false
	if yearValid(passport.birthYear, 1920, 2002) && yearValid(passport.issueYear, 2010, 2020) && yearValid(passport.expirationYear, 2020, 2030) {
		yearsValid = true
	}

	//validate
	if yearsValid && eyeColorValid(passport.eyeColor) && heightValid(passport.height) && validHairColor(passport.hairColor) && passportIDValid(passport.passportID) {
		return true
	}
	return false
}

func heightValid(height string) bool {
	if s.Contains(height, "cm") {
		newHeight := s.ReplaceAll(height, "cm", "")
		value, err := sc.Atoi(newHeight)
		if err != nil {
			return false
		}
		if 150 <= value && value <= 193 {
			return true
		}
	}

	if s.Contains(height, "in") {
		newHeight := s.ReplaceAll(height, "in", "")
		value, err := sc.Atoi(newHeight)
		if err != nil {
			return false
		}
		if 59 <= value && value <= 76 {
			return true
		}
	}
	return false
}

func validHairColor(hairColor string) bool {
	matchFirst, _ := regexp.MatchString("#([0-9]|[a-f]){6}", hairColor)
	//matchSecond, _ := regexp.MatchString("#[0-9]{6}", hairColor)
	if matchFirst {
		return true
	}
	return false
}

func passportIDValid(passportNumber string) bool {
	return len(passportNumber) == 9
}

func yearValid(year string, min int, max int) bool {
	value, err := sc.Atoi(year)
	if err != nil {
		return false
	}

	if min <= value && value <= max {
		return true
	}
	return false
}

func eyeColorValid(identifier string) bool {
	switch identifier {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

func splitIntoPassports(input []string) []passport {
	var passports []passport
	var currPassport passport

	for i := 0; i < len(input); i++ {
		if isNewPassport(input[i]) {
			passports = append(passports, currPassport)
			currPassport = passport{}
		} else {
			currPassport = readPassport(input[i], currPassport)
		}
	}
	return passports
}

func isNewPassport(lineBeingRead string) bool {
	if lineBeingRead == "" {
		return true
	}
	return false
}

func readPassport(line string, passport passport) passport {
	passportEntries := s.Split(line, " ")

	for i := 0; i < len(passportEntries); i++ {
		currentEntry := s.Split(passportEntries[i], ":")
		passport = saveValue(passport, currentEntry[0], currentEntry[1])
	}
	return passport
}

func saveValue(passport passport, identifier string, value string) passport {
	switch identifier {
	case "byr":
		passport.birthYear = value
	case "iyr":
		passport.issueYear = value
	case "eyr":
		passport.expirationYear = value
	case "hgt":
		passport.height = value
	case "hcl":
		passport.hairColor = value
	case "ecl":
		passport.eyeColor = value
	case "pid":
		passport.passportID = value
	case "cid":
		passport.countryID = value
	default:
		println("KABOOM. Santa forgot about a value in the input file: " + identifier)
	}
	return passport
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
