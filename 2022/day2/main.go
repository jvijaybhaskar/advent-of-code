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

func iterateStrategyGuide(rounds []string) int {

	var roundScore int

	for _, round := range rounds {

		outcomeScore, shapeScore := scoreCalculatorPart2(round)
		//fmt.Println(outcomeScore, shapeScore)

		roundScore = roundScore + outcomeScore + shapeScore

	}

	return roundScore
}

/*
Function to identify score for part 2

Data structure for Part 2
A -> Rock    - 1
B -> Paper   - 2
C -> Scissor - 3

Outcome score

X -> Lose -> 0
Y -> Draw -> 3
Z -> Win  -> 6


round score = shape score + outcome score
total scope = sum of all round scores


Shape combination

A X - Lose - C
A Y - Draw - A
A Z - Win  - B
B X - Lose - A
B Y - Draw - B
B Z - Win  - C
C X - Lose - B
C Y - Draw - C
C Z - Win  - A

*/
func scoreCalculatorPart2(round string) (int, int) {

	var shapeScoreMap = map[string]int{
		"A X": 3,
		"A Y": 1,
		"A Z": 2,
		"B X": 1,
		"B Y": 2,
		"B Z": 3,
		"C X": 2,
		"C Y": 3,
		"C Z": 1,
	}

	var outcomeScoreMap = map[string]int{
		"X": 0,
		"Y": 3,
		"Z": 6,
	}

	outcomeScore := outcomeScoreMap[string(round[2])]
	shapeScore := shapeScoreMap[round]

	return outcomeScore, shapeScore

}

/*
Function to calculate total score for part 1
*/
func calculateScore(rounds []string) int {

	var outcomeScoreMap = map[string]int{
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

	var shapeScoreMap = map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	var roundScore int

	for _, round := range rounds {

		outcomeScore := outcomeScoreMap[round]
		shapeScore := shapeScoreMap[string(round[2])]

		roundScore = roundScore + outcomeScore + shapeScore

	}

	return roundScore
}

/* Data structure for Part 1
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

	/* For testing functions
	var roundScore int
	testRounds := []string{
		"A Y",
		"B X",
		"C Z",
	}

	//roundScore =  calculateScore(testRounds)
	//roundScore = iterateStrategyGuide(testRounds)

	//fmt.Println(roundScore)
	*/

	// Read strategy guide data
	rounds := processInputData()

	// Determine win/lose/draw - outcome score
	// Determine shape score
	// Calcuate total score

	fmt.Println("The total possible score is: ", calculateScore(rounds))

	fmt.Println("The total possible score is: ", iterateStrategyGuide(rounds))

}
