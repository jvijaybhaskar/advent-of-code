package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Function to read input data and store in appropriate data structure for further processing
*/
func processInputData() []string {
	file, err := os.Open("./input/data.txt")

	if err != nil {
		fmt.Println("Err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rounds = []string{}

	for scanner.Scan() {

		scannedText := scanner.Text()

		rounds = append(rounds, scannedText)

	}

	return rounds
}

/*
Function to calculate total score
*/
func calculateScore(rounds []string) int {

	var outcomeScore = map[string]int{
		"A X": 3,
		"A Y": 6,
		"A Z": 0,
		"B X": 0,
		"B Y": 3,
		"B Z": 6,
		"C X": 6,
		"C Y": 0,
		"C Z": 3,
	}

	var shapeScore = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var roundScore int

	for _, round := range rounds {

		outcomeScore := outcomeScore[round]
		shapeScore := shapeScore[string(round[2])]

		roundScore = roundScore + outcomeScore + shapeScore

	}

	return roundScore
}

/* Data structure
A -> Rock
B -> Paper
C -> Scissor

X -> 1 -> Rock
Y -> 2 -> Paper
Z -> 3 -> Scissor

lose -> 0
draw -> 3
win ->  6

round score = shape score + outcome score
total scope = sum of all round scores


Outcome win/lose/draw combination

A X - Draw
A Y - Win
A Z - Lose
B X - Lose
B Y - Draw
B Z - Win
C X - Win
C Y - Lose
C Z - Draw

*/

func main() {

	// Test with dummy data

	/*
		var roundScore int

		testRounds := []string{
		"A X",
		"B Y",
		"C Z",
		}

		roundScore :=  calculateScore(testRounds)

		fmt.Println(roundScore)
	*/

	// Read strategy guide data
	rounds := processInputData()

	// Determine win/lose/draw - outcome score
	// Determine shape score
	// Calcuate total score

	fmt.Println("The total possible score is: ", calculateScore(rounds))

}
