package main

import (
	"bufio"
	"log"
	"os"
	sc "strconv"
	s "strings"
)

type bagRule struct {
	owner         string
	subcontainers []string
}

type bagQuantity struct {
	bagType  string
	quantity int
	children []bagQuantity
}

func main() {
	input := readInput()
	bagRules := processBagRules(input)
	shinyGold := bagQuantity{bagType: "shiny gold", quantity: 1}
	theBag := createCollectionOfBags(bagRules, shinyGold)
	println(calculateBags(theBag) - 1)
}

func part1(bagRules []bagRule) {
	var options []string
	options = parentFittingBag(bagRules, "shiny gold", options)
	println(len(options))
}

func parentFittingBag(bagRules []bagRule, bagtype string, options []string) []string {
	for i := 0; i < len(bagRules); i++ {
		if contains(bagRules[i].subcontainers, bagtype) {
			if !contains(options, bagRules[i].owner) {
				options = appendIfMissing(options, bagRules[i].owner)
				options = parentFittingBag(bagRules, bagRules[i].owner, options)
			}
		}
	}
	return options

}

//The bag trademarked trademarked copyright something
func calculateBags(theBagTM bagQuantity) int {
	sum := 0
	for i := 0; i < len(theBagTM.children); i++ {
		if len(theBagTM.children[i].children) > 0 {
			sum += calculateBags(theBagTM.children[i])
		} else {
			return theBagTM.quantity
		}
	}
	return sum*theBagTM.quantity + theBagTM.quantity
}

func createCollectionOfBags(bagRules []bagRule, currentBag bagQuantity) bagQuantity {
	var childrenBags []bagQuantity
	for i := 0; i < len(bagRules); i++ {
		if bagRules[i].owner == currentBag.bagType {
			for j := 0; j < len(bagRules[i].subcontainers); j++ {
				childBag := createBagQuantities(bagRules[i].subcontainers[j])
				updatedChildBag := createCollectionOfBags(bagRules, childBag)
				childrenBags = append(childrenBags, updatedChildBag)
			}
		}
	}
	currentBag.children = childrenBags
	return currentBag
}

func createBagQuantities(containers string) bagQuantity {
	if containers == "no other" {
		return bagQuantity{}
	}

	splitted := s.Split(containers, " ")
	if splitted[0] == "" {
		return createBagQuantities(s.TrimSpace(containers))
	}

	rest := delChar([]rune(containers), 0)
	quantity, _ := sc.Atoi(splitted[0])
	return bagQuantity{bagType: s.TrimSpace(string(rest)), quantity: quantity}
}
func delChar(s []rune, index int) []rune {
	return append(s[0:index], s[index+1:]...)
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}

func contains(stuff []string, e string) bool {
	for _, a := range stuff {
		if s.Contains(a, e) {
			return true
		}
	}
	return false
}

func processBagRules(inputRules []string) []bagRule {
	var bagRules []bagRule
	for _, rule := range inputRules {
		ruleSplit := s.Split(rule, " contain ")
		newBagRule := bagRule{owner: keepColorsOnly(ruleSplit[0]), subcontainers: toContainers(ruleSplit[1])}
		bagRules = append(bagRules, newBagRule)
	}
	return bagRules
}

func toContainers(ruleSplit string) []string {
	containers := s.Split(ruleSplit, ",")
	for i := 0; i < len(containers); i++ {
		containers[i] = keepColorsOnly(containers[i])
	}
	return containers
}

func keepColorsOnly(colorCode string) string {
	colorCode = s.ReplaceAll(colorCode, " bags", "")
	colorCode = s.ReplaceAll(colorCode, " bag", "")
	return s.ReplaceAll(colorCode, ".", "")
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
