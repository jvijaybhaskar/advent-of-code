package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cycle []int
var registerX int
var instructions []string

var signalStrength = make(map[int]int)

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

func processInstructions() {
	registerX = 1

	//Initializing to starting cycle from 1
	cycle = append(cycle, registerX) //Represents 0 - Dummy
	//fmt.Println(registerX, "=", "Dummy", "|", "0")
	cycle = append(cycle, registerX) // Represents 1st cycle
	//fmt.Println(registerX, "=", "Init", "|", "1")

	cycleTracker := 1

	for _, instruction := range instructions {

		if strings.Index(instruction, "noop") >= 0 {
			//simulating a clock cycle
			cycle = append(cycle, registerX)

			cycleTracker++
			//fmt.Println(registerX, " ", instruction, "|", cycleTracker)
		} else {
			//simulating two cycle to complete
			cycle = append(cycle, registerX)
			cycleTracker++
			//fmt.Println(registerX, " ", instruction, "|", cycleTracker)

			/*
				cycle = append(cycle, registerX)
				cycleTracker++
				fmt.Println(registerX, " ", instruction, "|", cycleTracker)
			*/

			parsedInstruction := strings.Split(instruction, " ")
			valueToAdd, _ := strconv.Atoi(parsedInstruction[1])

			registerX += valueToAdd

			//simulating a the second clock cycle
			cycle = append(cycle, registerX)

			cycleTracker++
			//fmt.Println(registerX, "=", instruction, "|", cycleTracker)
		}
	}

}

func computeSignalStrength() int {

	cycleCount := 20
	cycleInterval := 40

	var totalSignalStrength int

	for cycleCount <= len(cycle) {
		totalSignalStrength += cycleCount * cycle[cycleCount]

		fmt.Println(cycleCount, "*", cycle[cycleCount], "=", cycleCount*cycle[cycleCount])
		cycleCount += cycleInterval
	}

	/*
		for _, v := range signalStrength {
			totalSignalStrength += v
		}
	*/

	return totalSignalStrength

}

func displayOnCRT() {

	crtColLength := 40
	rowOffset := 1
	var crtRow string

	for i, cycleValue := range cycle {

		if i == 0 {
			continue
		}

		if rowOffset == cycleValue || rowOffset == cycleValue+1 || rowOffset == cycleValue+2 {
			crtRow += "#"
		} else {
			crtRow += "."
		}

		if rowOffset == crtColLength {
			rowOffset = 1
			fmt.Println(crtRow, "|", i)
			crtRow = ""
		} else {
			rowOffset++
		}

	}

}

func main() {

	instructions = readDataFromFile("./input/data.txt")
	processInstructions()

	fmt.Println("")
	fmt.Println("Sum of signal strengths for PART 1 = ", computeSignalStrength())
	fmt.Println("")

	displayOnCRT()

}
