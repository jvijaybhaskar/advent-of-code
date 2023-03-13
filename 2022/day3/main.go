package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	var rucksackInventory = []string{}

	for scanner.Scan() {

		rucksackContent := scanner.Text()

		rucksackInventory = append(rucksackInventory, rucksackContent)

	}

	return rucksackInventory
}

/*
 Data structure design

 Break string into two halves

 For every charachter in first half
 	 	search the second half
		if there is a match
			increment count
 Output the charachter that has the maximum count



*/

func calculateCharachterOccurance(backpackContent string) (string, int) {

	var compartmentLength int
	compartmentLength = len(backpackContent) / 2

	firstCompartment := backpackContent[:compartmentLength]
	secondCompartment := backpackContent[compartmentLength:]
	//fmt.Println(firstCompartment, secondCompartment)

	//var firstCompartmentCount []string

	var duplicateItem string
	var firstIndex int

	for i, charachter := range firstCompartment {

		j := strings.Index(secondCompartment, string(charachter))

		if j >= 0 {
			duplicateItem = string(charachter)
			firstIndex = i
			break
		}

		/*
			if j >= 0 {
				firstCompartmentCount = append(firstCompartmentCount, string(charachter))
			} else {
				firstCompartmentCount = append(firstCompartmentCount, "_")
			}
		*/

	}

	//fmt.Println(firstCompartmentCount)

	return duplicateItem, firstIndex

}

type itemPosition struct {
	item       string
	firstIndex int
}

// Find duplicate items for all rucksacks
func findDuplicateItems(rucksackInventory []string) []itemPosition {

	var duplicateItem string
	var firstIndex int
	var itemPositions []itemPosition

	for _, rucksackInventory := range rucksackInventory {
		duplicateItem, firstIndex = calculateCharachterOccurance(rucksackInventory)

		itemPosition := itemPosition{
			item:       duplicateItem,
			firstIndex: firstIndex,
		}
		itemPositions = append(itemPositions, itemPosition)

	}

	return itemPositions

}

const lowerPriority = "abcdefghijklmnopqrstuvwxyz"
const higherPriority = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func calculatePriorities(itemPositions []itemPosition) int {
	var priority int

	for _, itemPosition := range itemPositions {

		var position int

		if position = strings.Index(lowerPriority, itemPosition.item); position > 0 {
			priority += position + 1
			//fmt.Println(itemPosition.item, position+1)
		} else if position = strings.Index(higherPriority, itemPosition.item); position > 0 {
			priority += position + 27
			//fmt.Println(itemPosition.item, position+27)
		}

	}

	return priority

}

func calculateGroupPriorities(rucksackInventory []string) int {

	var groupStickers []itemPosition

	// Loop through each group
	for i := 0; i <= len(rucksackInventory); i = i + 3 {

		if i >= len(rucksackInventory) {
			break
		}

		// Loop through each charachter in the first bag of the group
		for _, rucksackContent := range rucksackInventory[i] {

			rucksackContentString := string(rucksackContent)
			var foundSticker bool
			var groupSticker itemPosition

			//Seach for common badge/charachter across rucksacks in a group
			for pos, charachter := range rucksackContentString {

				k := strings.Index(rucksackInventory[i+1], string(charachter))
				l := strings.Index(rucksackInventory[i+2], string(charachter))

				if k >= 0 && l >= 0 {
					foundSticker = true
					groupSticker.item = string(charachter)
					groupSticker.firstIndex = pos
					groupStickers = append(groupStickers, groupSticker)

					break
				}
			}

			if foundSticker {
				foundSticker = false
				break
			}

		}

	}

	//fmt.Println(groupStickers)

	// Calculate overall group priorities
	grpPriority := calculatePriorities(groupStickers)

	return grpPriority

}

func main() {

	// Process data
	rucksackInventory := processInputData()

	// Find duplicate items across compartments in a rucksack
	itemPositions := findDuplicateItems(rucksackInventory)

	// Calcuate priorities of dupliate items across all backpacks
	priority := calculatePriorities(itemPositions)

	fmt.Println("The sum of the priorities of duplicate items across rucksacks is: ", priority)

	// Calcuate priorities of badges across all elf groups
	groupPriority := calculateGroupPriorities(rucksackInventory)

	fmt.Println(groupPriority)

}
