package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type KnotPosition struct {
	x                int
	y                int
	previousPosition map[string]Position
}

type Position struct {
	x int
	y int
}

type Movement struct {
	direction string
	steps     int
}

var headMovements []Movement

func makeMove(move Movement, head *KnotPosition, tail *KnotPosition) {

	switch move.direction {
	case "U":
		head.y += move.steps
	case "D":
		head.y -= move.steps
	case "L":
		head.x -= move.steps
	case "R":
		head.x += move.steps
	}

}

func makeSmallMove(move Movement, head *KnotPosition, tail *KnotPosition) {

	for i := move.steps; i > 0; i-- {

		//fmt.Println(head.x, head.y)

		switch move.direction {
		case "U":
			head.y++
		case "D":
			head.y--
		case "L":
			head.x--
		case "R":
			head.x++
		}

		fmt.Println("H>", head.x, head.y)

		// tests if the tail is on the sample plane as head or not
		if head.x == tail.x || head.y == tail.y {
			offsetLaterally(head, tail)
		} else {
			//Tests if the tail will align with head if moved one step diagonally, if not move one step diagonally
			if head.x != tail.x+1 && head.x != tail.x-1 || head.y != tail.y-1 && head.y != tail.y+1 {
				offsetDiaognally(head, tail)
			}
		}

	}
}

func offsetLaterally(head *KnotPosition, tail *KnotPosition) {

	// same vertical line
	if tail.y == head.y {
		if tail.x > head.x {
			tail.x = head.x + 1
		}

		if tail.x < head.x {
			tail.x = head.x - 1
		}
	}

	// same horizontal line
	if tail.x == head.x {
		if tail.y > head.y {
			tail.y = head.y + 1
		}

		if tail.y < head.y {
			tail.y = head.y - 1
		}
	}

	fmt.Println("T>", tail.x, tail.y)
	stringPosition := strconv.Itoa(tail.x) + " " + strconv.Itoa(tail.y)
	tail.previousPosition[stringPosition] = Position{tail.x, tail.y}

}

func offsetDiaognally(head *KnotPosition, tail *KnotPosition) {
	if tail.x > head.x {
		if tail.y > head.y {
			//3rd quad
			tail.x--
			tail.y--

		} else {
			//2nd quad
			tail.x--
			tail.y++
		}

	} else {
		if tail.y > head.y {
			tail.x++
			tail.y--
		} else {
			tail.x++
			tail.y++
		}
	}

	fmt.Println("T>", tail.x, tail.y)

	stringPosition := strconv.Itoa(tail.x) + " " + strconv.Itoa(tail.y)
	tail.previousPosition[stringPosition] = Position{tail.x, tail.y}
}

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

/*
Function to process raw data and convert to two dimentional grid
*/
func processRawData(rawFileData []string) {

	for _, row := range rawFileData {

		splitRow := strings.Split(row, " ")
		move, _ := strconv.Atoi(splitRow[1])

		movement := Movement{splitRow[0], move}
		headMovements = append(headMovements, movement)

	}

}

//Function to visualize the results to compare with the output in the puzzle
func visualizeGrid(tailPosition KnotPosition) {

	numRows := 5
	numColumns := 6

	// Initialize a ten length slice of empty slices
	grid := make([][]string, numRows)

	// Initialize those 10 empty slices
	for i := 0; i < numRows; i++ {
		grid[i] = make([]string, numColumns)
	}

	for _, position := range tailPosition.previousPosition {
		grid[position.y][position.x] = "#" // rows represent y and columns represent x

	}

	// grid is a 2d slice of strings
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == "#" {
				fmt.Print(grid[i][j])
			} else {
				grid[i][j] = "."
				fmt.Print(grid[i][j])
			}
		}
		fmt.Println("")

	}

}

func main() {

	rawFileContent := readDataFromFile("./input/data_test.txt")
	processRawData(rawFileContent)

	startPosition := make(map[string]Position)

	startPosition["0 0"] = Position{0, 0}

	//Initialize head position
	var headPosition = KnotPosition{
		x:                0,
		y:                0,
		previousPosition: startPosition,
	}

	//Initialize tail position
	var tailPosition = KnotPosition{
		x:                0,
		y:                0,
		previousPosition: startPosition,
	}

	for _, move := range headMovements {
		//fmt.Println("Before move: ", headPosition.x, headPosition.y)
		makeSmallMove(move, &headPosition, &tailPosition) // move head as per instruction
		//fmt.Println("After move: ", headPosition.x, headPosition.y)
	}

	fmt.Println("---------")
	fmt.Println("Numbert of unique posittions as tail end of rope moved is: ", len(tailPosition.previousPosition))
	fmt.Println("---------")

	visualizeGrid(tailPosition)

	//initializeGrid()

}
