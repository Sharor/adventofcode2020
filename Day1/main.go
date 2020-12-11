package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	nums := readInput()
	fmt.Println(scanFor2020(nums))
}

func scanFor2020(expenseReport []int) int {
	for i := 0; i < len(expenseReport); i++ {
		for j := 0; j < len(expenseReport); j++ {
			for k := 0; k < len(expenseReport); k++ {
				if i == j {
					j++
				}
				if i == k {
					k++
				}
				if j == k && len(expenseReport) != k {
					k++
				}
				if k == 200 {
					k = 199
				}
				if expenseReport[i]+expenseReport[j]+expenseReport[k] == 2020 {
					return expenseReport[i] * expenseReport[j] * expenseReport[k]
				}
			}
		}
	}
	return 0
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
