package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

//
// More info on how to read file here https://golangdocs.com/reading-files-in-golang
//
func processInputData() [][]int {
	file, err := os.Open("./input/data.txt")

	if err != nil {
		fmt.Println("Err")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var inventory [][]int
	var calories []int

	for scanner.Scan() {

		scannedText := scanner.Text()

		if scannedText != "" {
			//fmt.Println("Read data:", scanner.Text())
			dataPoint, err := strconv.Atoi(scannedText)
			if err != nil {
				fmt.Println("Unable to process input data::", err)
				panic(err)
			}

			calories = append(calories, dataPoint)

		} else {

			inventory = append(inventory, calories)
			calories = nil

		}
	}

	if !scanner.Scan() {
		inventory = append(inventory, calories)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return inventory
}

func calorieCount(calories []int) int {
	var totalCalorie int
	for _, calorie := range calories {
		totalCalorie += calorie
	}

	return totalCalorie
}

func largestNumber(calorieSummaryList []int) (largestCalorieCount int, position int) {

	inventoryLen := len(calorieSummaryList)
	largestCalorie := calorieSummaryList[0]

	for j := 1; j < inventoryLen; j++ {
		if largestCalorie < calorieSummaryList[j] {
			largestCalorie = calorieSummaryList[j]
			position = j
		}
	}

	return largestCalorie, position
}

func topSnackHolders(calorieSummaryList []int) int {

	sort.Slice(calorieSummaryList, func(i, j int) bool {
		return calorieSummaryList[i] > calorieSummaryList[j]
	})

	//fmt.Println(calorieSummaryList[0], calorieSummaryList[1], calorieSummaryList[2])
	topThreeCalorieCount := calorieSummaryList[0] + calorieSummaryList[1] + calorieSummaryList[2]
	return topThreeCalorieCount

}

func computeCalorieTotalList(inventory [][]int) []int {

	var calorieSummary []int

	for _, calories := range inventory {
		totalCalorie := calorieCount(calories)
		calorieSummary = append(calorieSummary, totalCalorie)
	}

	return calorieSummary

}

func main() {
	fmt.Println(">>Solving day 1 requirement<<")
	// Process data into a suitable data structure
	// Calorie > []Slice int
	// Calories > []*Slice

	// Test with trial data

	/*
		var calorie1 = []int{1000, 200, 300}
		var calorie2 = []int{1000, 500, 3000}

		var inventory = [][]int{calorie1, calorie2}

		for _, calories := range inventory {
			fmt.Println(calories)
			fmt.Println(calorieCount(calories))
		}
	*/

	inventory := processInputData()
	fmt.Println("Processed data >> ", inventory)

	// Loop through the data struct to compute the largest a data point

	calorieSummary := computeCalorieTotalList(inventory)
	fmt.Println("CalorieSummary data >> ", calorieSummary)

	// Output the the entity with the largest data point

	largetClorieCount, elfPosition := largestNumber(calorieSummary)

	fmt.Println("")
	fmt.Printf("The lasrgest calorie count is: %d \nIt is held by elf at position: %d \n", largetClorieCount, elfPosition)
	fmt.Println("")
	fmt.Printf("The sum of calories held up by top three elfs is: %d \n", topSnackHolders(calorieSummary))

}
