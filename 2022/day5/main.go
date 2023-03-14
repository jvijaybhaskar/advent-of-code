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

/*
	Data strcuture design

	Stack data
	A map might be suitable. The keys will represent stack.
	The values will be a string based slice representing the crates in a stack


	Move instructions data
	This will be represented using multi dimentional slices.
	Each row from raw data will be an entry in the slice.
	The move steps in a single row will be represented as an element of a slice after being parsed

*/
func buildStack(rawData []string) map[int][]string {

	var stackData = make(map[int][]string)

	for i := len(rawData) - 1; i >= 0; i-- {

		var stack int
		row := rawData[i]

		for j := 0; j <= len(row); j = j + 4 {

			/*
				if j == len(row) {
					fmt.Println("Break: ", j)
					break
				}
			*/

			//fmt.Println(j, row[j:j+3])

			stack++
			cellEntry := strings.ReplaceAll(row[j:j+3], " ", "")
			if cellEntry == "" {
				//stackData[stack] = append(stackData[stack], "")
			} else {
				//cellEntry = strings.ReplaceAll(row[j:j+3], "[", "")
				//stackData[stack] = strings.ReplaceAll(cellEntry, "]", ""))
				stackData[stack] = append(stackData[stack], row[j:j+3])
			}

			//stackData[stack] = reverseSlice(stackData[stack])

		}

	}

	return stackData

}

/*
	Function to prepare moves data and normalize it into an array
*/
func prepareMoveSteps(rawData []string) [][]int {

	var allMoves [][]int

	for _, row := range rawData {

		splitAction := strings.Split(row, " ")
		//fmt.Println(splitAction[1], splitAction[3], splitAction[5])

		crateCount, _ := strconv.Atoi(splitAction[1])
		fromStack, _ := strconv.Atoi(splitAction[3])
		toStack, _ := strconv.Atoi(splitAction[5])

		move := []int{crateCount, fromStack, toStack}

		allMoves = append(allMoves, move)

	}

	return allMoves

}

/*
Function to make the actual moves and rearrange stacks as per move data
*/
func makeMove(actions [][]int, stackData map[int][]string) {

	for _, action := range actions {

		//fmt.Println("Action >", action)

		toSlice := stackData[action[2]]
		fromSlice := stackData[action[1]]

		for j := action[0]; j > 0; j-- {

			toSlice = append(toSlice, fromSlice[len(fromSlice)-1])
			fromSlice = stackData[action[1]][:len(fromSlice)-1]

		}

		stackData[action[1]] = fromSlice
		stackData[action[2]] = toSlice

	}

	for k := 1; k <= len(stackData); k++ {
		fmt.Println(k, stackData[k])
	}

}

/*
Function to make moves as per part 2 requirement
*/
func makeMove9001(actions [][]int, stackData map[int][]string) {
	for _, action := range actions {

		//fmt.Println("Action >", action)

		toSlice := stackData[action[2]]
		fromSlice := stackData[action[1]]

		cratesToMove := fromSlice[len(fromSlice)-action[0]:]

		toSlice = append(toSlice, cratesToMove...)
		fromSlice = stackData[action[1]][:len(fromSlice)-action[0]] //Remove crates after move

		stackData[action[1]] = fromSlice
		stackData[action[2]] = toSlice

	}

	for k := 1; k <= len(stackData); k++ {
		fmt.Println(k, stackData[k])
	}
}

func processRawData(rawData []string) {

}

func main() {

	//TestData

	/*
		rawData := ReadDataFromFile("./input/data.txt")
		stackData := buildStack(rawData[:3])
		allMoves := prepareMoveSteps(rawData[5:9])
	*/

	rawData := ReadDataFromFile("./input/data_large.txt")
	stackData := buildStack(rawData[:8])
	allMoves := prepareMoveSteps(rawData[10:])

	//makeMove(allMoves, stackData)
	makeMove9001(allMoves, stackData)

}
