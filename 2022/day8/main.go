package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	for x, treeRow := range rawFileData {

		var row []*treeProp

		for y, tree := range treeRow {

			treeHeight, _ := strconv.Atoi(string(tree))
			var treeVisibiity bool

			if x == 0 || y == 0 || x == len(treeRow)-1 || y == len(treeRow)-1 {
				treeVisibiity = true
			}

			var treeDetails = treeProp{
				height:  treeHeight,
				visible: treeVisibiity,
			}

			row = append(row, &treeDetails)
		}

		treeGrid = append(treeGrid, row)

	}
	printTreeGrid()

}

func printTreeGrid() {

	fmt.Println("____________________________")
	fmt.Println("")
	for _, row := range treeGrid {
		for _, tree := range row {
			fmt.Printf("%v", tree.height)
		}
		fmt.Println("")
	}

}

/*
Function to mark trees as visible

0. Read raw data

1. Build Grid from raw data

2. For trees on the edge of the grid
	mark as visible

3. Do this for all trees inside the grid
	To find neighbours of a tree:

	Loop vertically
		If there is a taller neighbour
			mark as invisible
			break
		else
			mark as visible


	Loop horizontally
		If there is taller tree
			mark as invisible
			break
		else
			mark as visible

*/
func determineTreeVisiblility(x int, y int, tree *treeProp) int {

	treeRow := treeGrid[x]
	var isVisibleLeft bool
	var isVisibleRight bool

	var isVisibleTop bool
	var isVisibleBottom bool

	var visibleTrees int

	for i := 0; i < y; i++ {

		if treeRow[i].height >= tree.height {
			isVisibleLeft = false
			break
		} else {
			isVisibleLeft = true
		}

	}

	for i := y + 1; i < len(treeRow); i++ {

		if treeRow[i].height >= tree.height {
			isVisibleRight = false
			break
		} else {
			isVisibleRight = true
		}

	}

	for i := 0; i < x; i++ {
		if treeGrid[i][y].height >= tree.height {
			isVisibleTop = false
			break
		} else {
			isVisibleTop = true
		}

	}

	for i := x + 1; i < len(treeGrid); i++ {

		if treeGrid[i][y].height >= tree.height {
			isVisibleBottom = false
			break
		} else {
			isVisibleBottom = true
		}

	}

	tree.visible = isVisibleLeft || isVisibleRight || isVisibleTop || isVisibleBottom
	//tree.visible = false

	if tree.visible {
		visibleTrees++
	}

	return visibleTrees

}

/*
Funcation to iterate through a grid to determine total visible trees
*/
func assessVisibility() {

	var visibleTrees int
	for x, treeRow := range treeGrid {
		for y, tree := range treeRow {
			if !tree.visible {
				visibleTrees += determineTreeVisiblility(x, y, tree)
			} else {
				visibleTrees++
			}
		}
	}

	printTreeGrid()

	fmt.Println("Visible Trees: ", visibleTrees)
}

/*
Data structure design

1. Properties of the tree (height and visibility) is tracked using a struct
2. The location of tree in a grid is tracked using a two dimentional slice of trees
3. Individual tree visibility is tracked in a struct. Total visible trees are determined by looping through the grid

*/

type treeProp struct {
	height  int
	visible bool
}

var treeGrid [][]*treeProp
var visbibleTrees map[string]int

func main() {

	rawFileContent := readDataFromFile("./input/data.txt")
	processRawData(rawFileContent)
	assessVisibility()

}
