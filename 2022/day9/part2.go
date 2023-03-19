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
	x       int
	y       int
	knotNum string
}

type Movement struct {
	direction string
	steps     int
}

var headMovements []Movement

var knots []*KnotPosition

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

func makeSmallMove(move Movement, knots []*KnotPosition) {

	for i := move.steps; i > 0; i-- {

		//fmt.Println(knots[0].x, knots[0].y)

		switch move.direction {
		case "U":
			knots[0].y++
		case "D":
			knots[0].y--
		case "L":
			knots[0].x--
		case "R":
			knots[0].x++
		}

		//fmt.Println("H>", knots[0].x, knots[0].y)

		for i := 0; i < len(knots)-1; i++ {
			determineOffsettype(knots[i], knots[i+1], i+1)
		}

	}
}

func determineOffsettype(head *KnotPosition, tail *KnotPosition, knotNum int) {
	// tests if the tail is on the sample plane as head or not

	var stringPosition string
	if head.x == tail.x || head.y == tail.y {
		x, y := offsetLaterally(head, tail, knotNum)
		stringPosition = strconv.Itoa(x) + " " + strconv.Itoa(y)
	} else {
		//Tests if the tail will align with head if moved one step diagonally, if not move one step diagonally
		if head.x != tail.x+1 && head.x != tail.x-1 || head.y != tail.y-1 && head.y != tail.y+1 {
			x, y := offsetDiaognally(head, tail, knotNum)
			stringPosition = strconv.Itoa(x) + " " + strconv.Itoa(y)

		} else {
			stringPosition = ""
		}
	}

	if stringPosition != "" {

		tail.previousPosition[stringPosition] = Position{tail.x, tail.y, strconv.Itoa(knotNum)}

	}

}

func offsetLaterally(head *KnotPosition, tail *KnotPosition, knotNum int) (int, int) {

	//fmt.Println("Moving laterally")

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

	//fmt.Println("T>", tail.x, tail.y)
	return tail.x, tail.y

}

func offsetDiaognally(head *KnotPosition, tail *KnotPosition, knotNum int) (int, int) {

	//fmt.Println("Moving Diagonally")

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

	//fmt.Println("T>", tail.x, tail.y)

	return tail.x, tail.y

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
//CAUTION: works for test data only
func visualizeGrid(tailPosition KnotPosition) {

	numRows := 25
	numColumns := 30

	// Initialize a ten length slice of empty slices
	grid := make([][]string, numRows)

	// Initialize those 10 empty slices
	for i := 0; i < numRows; i++ {
		grid[i] = make([]string, numColumns)
	}

	for _, position := range tailPosition.previousPosition {
		fmt.Println(position.x, position.y)
		grid[position.y+6][position.x+12] = position.knotNum // rows represent y and columns represent x

	}

	// grid is a 2d slice of strings
	for i := len(grid) - 1; i >= 0; i-- {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "" {
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

	rawFileContent := readDataFromFile("./input/data.txt")
	processRawData(rawFileContent)

	//startPosition := make(map[string]Position)
	//startPosition["0 0"] = Position{0, 0, "S"}

	//initilizing 9 knots
	for i := 0; i < 10; i++ {
		startPosition := make(map[string]Position)
		startPosition["0 0"] = Position{0, 0, "S"}

		var knot = KnotPosition{
			x:                0,
			y:                0,
			previousPosition: startPosition,
		}

		knots = append(knots, &knot)
	}

	for _, move := range headMovements {
		//fmt.Println("Before move: ", knots[0].x, knots[0].y)
		makeSmallMove(move, knots) // move head as per instruction and adjust the rest of knots
		//fmt.Println("After move: ", knots[0].x, knots[0].y)
	}

	fmt.Println("---------")
	fmt.Println("Number of unique posittions as tail end of rope moved is: ", len(knots[len(knots)-1].previousPosition))
	fmt.Println("---------")

	//visualizeGrid(*knots[len(knots)-1])

}
