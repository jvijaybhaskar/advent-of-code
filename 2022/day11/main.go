package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	monkeyId         int
	startingItems    []int
	holdingItems     []int
	operation        string
	divisibilityTest int
	action           []int
	inspectedItems   int
}

var monkeyList []*Monkey

// Read data from data file
func readDataFromFile(filepath string) []string {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("Error while reading file", err)
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

func processRawData(rawData []string) {

	var currentMonkeyId int
	var monkeyDetails []string

	for i := 0; i < len(rawData); i = i + 7 {

		var currentMonkey Monkey

		//create a new Monkey and add to map
		dataRow := rawData[i]
		monkeyDetails = strings.Split(dataRow, " ")
		currentMonkeyId, _ = strconv.Atoi(strings.TrimRight(monkeyDetails[1], ":"))

		currentMonkey.monkeyId = currentMonkeyId

		fmt.Println(dataRow)
		//Populate monkey details

		//Populate starting items data
		dataRow = rawData[i+1]
		startingItems := strings.Split(strings.Split(dataRow, ": ")[1], ", ")
		var items []int
		for _, item := range startingItems {
			intItem, _ := strconv.Atoi(item)
			items = append(items, intItem)
		}

		currentMonkey.startingItems = items

		//Populating operations information
		dataRow = rawData[i+2]
		operation := strings.Split(dataRow, " = ")[1]
		currentMonkey.operation = operation

		//Populating Divisibility test details
		dataRow = rawData[i+3]
		divisibilityTest := strings.SplitAfter(dataRow, "Test: divisible by ")
		divisibleBy, _ := strconv.Atoi(divisibilityTest[1])
		currentMonkey.divisibilityTest = divisibleBy

		//Populating true action
		dataRow = rawData[i+4]
		trueAction := strings.SplitAfter(dataRow, "    If true: throw to monkey ")
		iTrueAction, _ := strconv.Atoi(trueAction[1])
		currentMonkey.action = append(currentMonkey.action, iTrueAction)

		//Populating false action
		dataRow = rawData[i+5]
		falseAction := strings.SplitAfter(dataRow, "    If false: throw to monkey ")
		iFalseAction, _ := strconv.Atoi(falseAction[1])
		currentMonkey.action = append(currentMonkey.action, iFalseAction)

		fmt.Println(currentMonkey)
		monkeyList = append(monkeyList, &currentMonkey)

	}
}

func multiply(a int, b int) int {
	return a * b
}

func add(a int, b int) int {
	return a + b
}

func keepAway(monkey *Monkey, items []int, puzzlePart int) {

	var worryFactor int
	var reductionFactor int

	if puzzlePart == 1 {
		reductionFactor = 3
	} else {
		reductionFactor = computeReductionFactor()
	}

	//identify operation
	operation := strings.Split(monkey.operation, " ")

	for _, worryLevel := range items {

		//fmt.Println("Monkey and Worry level:::: ", monkey.monkeyId, worryLevel)

		var currentWorryLevel, reducedWorryLevel int

		if operation[2] == "old" {
			worryFactor = worryLevel
		} else {
			worryFactor, _ = strconv.Atoi(operation[2])
		}

		switch operation[1] {
		case "*":
			currentWorryLevel = multiply(worryLevel, worryFactor)

		case "+":
			currentWorryLevel = add(worryLevel, worryFactor)
		}

		if puzzlePart == 1 {
			//this is needed for Part 1
			reducedWorryLevel = currentWorryLevel / reductionFactor
		} else {
			reducedWorryLevel = currentWorryLevel % reductionFactor

		}

		//Move the item to the holding list of another monkey
		passItem(reducedWorryLevel, monkey)

	}

}

func passItem(reducedWorryLevel int, monkey *Monkey) {

	if reducedWorryLevel%monkey.divisibilityTest == 0 {
		receivingMonkey := monkeyList[monkey.action[0]]
		receivingMonkey.holdingItems = append(receivingMonkey.holdingItems, reducedWorryLevel)

	} else {
		receivingMonkey := monkeyList[monkey.action[1]]
		receivingMonkey.holdingItems = append(receivingMonkey.holdingItems, reducedWorryLevel)
	}

}

func computeReductionFactor() int {
	reductionFator := 1
	for _, monkey := range monkeyList {
		reductionFator *= monkey.divisibilityTest
	}

	return reductionFator
}

func runRounds(rounds int, puzzlePart int) {

	for i := 0; i < rounds; i++ {

		//Move items as per instructions in data set
		for _, monkey := range monkeyList {

			var items []int

			items = monkey.startingItems
			items = append(items, monkey.holdingItems...)

			monkey.inspectedItems += len(items)

			//empty all holding items as they are going to be thrown
			monkey.startingItems = make([]int, 0)
			monkey.holdingItems = make([]int, 0)

			keepAway(monkey, items, puzzlePart)

		}

		//Rerunning the loop to find the output after the keep away round
		fmt.Println("Output after round:", i+1)

		for _, monkey := range monkeyList {

			fmt.Printf("MonkeyID:%v | InspectedItems:%d \n", monkey.monkeyId, monkey.inspectedItems)

		}

	}

}

func main() {

	rawData := readDataFromFile("./input/data_test.txt")
	processRawData(rawData)

	//Code to run part 1

	fmt.Println("-----------------------PART1-------------------------")
	puzzlePart := 1
	runRounds(20, puzzlePart)

	fmt.Println("-----------------------PART2-------------------------")
	//Code to run part 2
	puzzlePart = 2
	runRounds(10000, puzzlePart)

}
