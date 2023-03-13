package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Read data from data file
func ReadDataFromFile(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileContent = []string{}

	for scanner.Scan() {

		rowInFile := scanner.Text()

		fileContent = append(fileContent, rowInFile)

	}

	return fileContent
}

/* Sample input:

"2-4,6-8"
"2-3,4-5"

*/
func extractExtremes(pairAssigment string) ([]string, []string) {

	assignmentPair := strings.Split(pairAssigment, ",")

	firstPair := strings.Split(assignmentPair[0], "-")
	secondPair := strings.Split(assignmentPair[1], "-")

	return firstPair, secondPair
}

/*

Sample input
"2-4", "6-8"

Output
Determines if one pair in other or not
Determines if the pairs are disjointed or not

*/
func compareExtremes(firstPair []string, secondPair []string) (bool, bool, bool) {

	pair1Left, err1l := strconv.Atoi(firstPair[0])
	pair1Right, err1r := strconv.Atoi(firstPair[1])
	pair2Left, err2l := strconv.Atoi(secondPair[0])
	pair2Right, err2r := strconv.Atoi(secondPair[1])

	var pair1in2, pair2in1, pairDisjointed bool

	if err1l == nil && err1r == nil && err2l == nil && err2r == nil {

		//Pair 1 in Pair 2
		if pair1Left >= pair2Left && pair1Right <= pair2Right {
			pair1in2 = true
		}

		//Pair 2 in pair 1
		if pair2Left >= pair1Left && pair2Right <= pair1Right {
			pair2in1 = true
		}

		//Pairs disjointed
		if pair1Right < pair2Left || pair2Right < pair1Left {
			pairDisjointed = true
		}

	}

	return pair1in2, pair2in1, pairDisjointed

}

func loopThroughData(sectionAssignments []string) (int, int) {

	var subsetPairCount int
	var disjointedPairCount int

	for _, assignments := range sectionAssignments {
		pair1in2, pair2in1, pairDisjointed := compareExtremes(extractExtremes(assignments))
		if pair1in2 || pair2in1 {
			subsetPairCount++
		}

		if pairDisjointed {
			disjointedPairCount++
		}
	}

	return subsetPairCount, disjointedPairCount

}

func main() {

	sectionAssignments := ReadDataFromFile("./input/data.txt")

	subsetPairCount, disjointedPairCount := loopThroughData(sectionAssignments)
	fmt.Println("Assignment pairs that fully contain ther other: ", subsetPairCount)

	fmt.Println("Assignment pairs that overlap: ", len(sectionAssignments)-disjointedPairCount)

}
