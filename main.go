package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var results []int

func main() {
	var numTest int
	fmt.Scanln(&numTest)

	doTest(numTest)
	printResults(results, 0)
}

func doTest(numTest int) {
	// get number of test
	var numInput int
	fmt.Scanln(&numInput)

	// get inputs
	var inputs string
	inputs, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("failed to read input %s \n", err)
	}
	// trim last \n
	inputs = strings.TrimSuffix(inputs, "\n")

	// split string then calculate the sum of squares
	list := strings.Split(inputs, " ")
	result := calSquare(list, 0, 0)

	// add result to results
	results = append(results, result)

	// run next test
	if numTest > 1 {
		doTest(numTest - 1)
	}
}

func calSquare(list []string, index int, result int) int {
	if index < len(list) {
		// convert string int to int
		val, err := strconv.Atoi(list[index])
		if err != nil {
			fmt.Println(err)
		}

		// add square of int if it's possitive
		if val > 0 {
			result += val * val
		}

		return calSquare(list, index+1, result)
	} else {
		return result
	}
}

func printResults(results []int, index int) {
	if index < len(results) {
		fmt.Println(results[index])
		printResults(results, index+1)
	}
}
