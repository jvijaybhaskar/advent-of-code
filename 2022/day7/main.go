package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Overall logic [Possible overcomplicated!]

Step 1: Build folder structure

Process raw input to discover the folder structure

	cd <folder>
		intialize current folder
		create a new folder struct if not recorded
			record the parent if any

	cd ..
		update current folder


Step 2: Calculate folder size
Sum up size of files in a directory and update the size of folder and ALL the parents in the chain


	ls
		iterate through ls results
			create a list of directory if any
			create a list of files if any [TBD]
			summarize size of files in a dir and update directory size field
			update the size of ALL parents in the chain [due to the requirement of part 1]
		store the folder in the map ds
			key - folder name concatenated with parent
			value - the `dir` struct



Step 3: Loop through map to summarize directories with size UPTO 100,000 [for Part 1]

*/

/*
Data structure design


	Directory modelled as a Struct. It stores subdirectory and files under it.

		type Dir struct {
			dirName string
			subDir  []string
			files	[]string
			parentDir string
			dirSize	int
			subdirSize int
		}


	Map to store directory list in a flat format.
	Represent directoryName as full path < root/parentDir/dir...>

		[directoryName] directory struct


*/

type Directory struct {
	dirName       string
	subDir        []string
	files         []string
	parentDirPath string
	dirSize       int
	subDirSize    int
}

/*
initialize data structures
*/

var directoryMap = map[string]*Directory{}

var currentDirectoryName string
var currentDirPath string

func changeDirectory(dirName string) {

	if dirName == "$ cd .." {

		// go up the dir tree and set current dir and parent diretory path
		splitParentPath := strings.Split(currentDirPath, "/")
		parentHeight := len(splitParentPath) - 1
		currentDirectoryName = splitParentPath[parentHeight]
		currentDirPath = strings.Join(splitParentPath[:parentHeight], "/")

		//fmt.Println("Processing >> ", dirName)
		//fmt.Println(currentDirectoryName, currentDirPath)

	} else {

		if dirName == "$ cd /" {

			currentDirectoryName = "root"
			currentDirPath = "root"

		} else {
			folderName := strings.Split(dirName, " ")
			currentDirectoryName = folderName[2]
			currentDirPath = currentDirPath + "/" + currentDirectoryName
		}

		// Add new entry to directory map if the folder does not exist
		var directory Directory = Directory{}
		directory.dirName = currentDirectoryName
		directory.parentDirPath = currentDirPath

		directoryMap[currentDirPath] = &directory

		//fmt.Println("Processing >> ", dirName)
		//fmt.Println(currentDirectoryName, currentDirPath)

	}

}

func processDirContent(dirContent string) {

	childDetails := strings.Split(dirContent, " ")

	if childDetails[0] == "dir" {

		directory := directoryMap[currentDirPath]
		directory.subDir = append(directory.subDir, childDetails[1])

	} else {

		fileSize, _ := strconv.Atoi(childDetails[0])

		directory := directoryMap[currentDirPath]
		directory.dirSize += fileSize

		parentChain := strings.Split(currentDirPath, "/")
		parentChainLength := len(parentChain)

		var prevFolder string

		for _, folder := range parentChain[:parentChainLength-1] {

			var directory *Directory

			if folder == "root" {
				prevFolder = folder
			} else {
				prevFolder = prevFolder + "/" + folder
			}

			directory = directoryMap[prevFolder]
			directory.dirSize += fileSize

		}

	}

	//fmt.Println("After update", *directoryMap[currentDirPath])

}

// Read data from data file
func readDataFromFile(filepath string) []string {
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

func main() {

	rawData := readDataFromFile("./input/data.txt")

	i := 0
	for i < len(rawData) {

		line := rawData[i]
		if strings.Index(line, "$ cd") >= 0 {
			changeDirectory(line)
		} else if strings.Index(line, "$ ls") >= 0 {
			i++
			continue
		} else {
			processDirContent(line)
		}

		i++
	}

	var candidateFileSize int
	var seperator = "________________________________________________________"
	for path, directory := range directoryMap {

		fmt.Println(path, seperator[:len(seperator)-len(path)], directory.subDir, directory.dirSize)
		if directory.dirSize <= 100000 {
			candidateFileSize += directory.dirSize
		}

	}
	fmt.Println(candidateFileSize)

}
