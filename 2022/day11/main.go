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
}

var monkeyList = make(map[int]*Monkey)

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
		monkeyList[currentMonkeyId] = &currentMonkey

	}
}


func multiply(a int, b int) int {
	return a * b
}

func add(a int, b int) int {
	return a + b
}

func keepAway(monkey *Monkey) {

	//identify operation
	operation := strings.Split(monkey.operation, " ")
	worryFactor, _ := strconv.Atoi(operation[2])

	for _, worryLevel := range monkey.startingItems {

		var currentWorryLevel, reducedWorryLevel int

		switch operation[1] {
		case "*":
			currentWorryLevel = multiply(worryLevel, worryFactor)

		case "+":
			currentWorryLevel = add(worryLevel, worryFactor)
		}

		reducedWorryLevel = currentWorryLevel / 3

		if reducedWorryLevel%monkey.divisibilityTest == 0 {
			receivingMonkey := monkeyList[monkey.action[0]]
			receivingMonkey.holdingItems = append(receivingMonkey.holdingItems, currentWorryLevel)

		} else {
			receivingMonkey := monkeyList[monkey.action[1]]
			receivingMonkey.holdingItems = append(receivingMonkey.holdingItems, currentWorryLevel)
		}

	}

}




func main() {

	rawData := readDataFromFile("./input/data_test.txt")
	processRawData(rawData)

	
	
	

}
