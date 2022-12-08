package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var currentlyListingContent bool
var baseDir Content

const diskSize = 70000000
const updateSize = 30000000

func main() {
	currentlyListingContent = false

	f, err := os.Open("resources/files.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	baseDir = Content{isDir: true}
	var currentDir *Content

	for scanner.Scan() {
		cmdString := scanner.Text()
		fmt.Printf("%s\n", cmdString)
		if cmdString[0:1] == "$" {
			currentlyListingContent = false
			currentDir = actionCommand(cmdString, currentDir)
		} else if currentlyListingContent {
			currentDir = addFileOrDirectoryToContents(currentDir, cmdString)
		}
	}

	content := calculateDirectorySize(&baseDir)

	totalFreeSpace := diskSize - content.size

	fmt.Printf("Total free space: %d\n", totalFreeSpace)

	spaceRequiredForUpdate := updateSize - totalFreeSpace

	fmt.Printf("Total space required for update: %d\n", spaceRequiredForUpdate)

	directories := filterDirectoriesBySize(content, spaceRequiredForUpdate, math.MaxInt)

	var smallestDir *Content

	for _, directory := range directories {
		if smallestDir == nil || directory.size < smallestDir.size {
			smallestDir = directory
		}
	}

	fmt.Printf("Smallest single directory to delete (%d): %s\n", smallestDir.size, smallestDir.name)
}

func filterDirectoriesBySize(content *Content, min int, max int) []*Content {
	directories := make([]*Content, 0)
	for _, contentEntry := range content.contents {
		if contentEntry.isDir {
			if contentEntry.size < max && contentEntry.size > min {
				directories = append(directories, contentEntry)
			}
			directories = append(directories, filterDirectoriesBySize(contentEntry, min, max)...)
		}
	}
	return directories
}

func calculateDirectorySize(content *Content) *Content {
	total := 0
	for _, contentEntry := range content.contents {
		if contentEntry.isDir {
			total = total + calculateDirectorySize(contentEntry).size
		} else {
			total = total + contentEntry.size
		}
	}

	content.size = total

	return content
}

func actionCommand(command string, currentDir *Content) *Content {
	if command[2:4] == "cd" {
		dir := strings.Split(command, "$ cd ")[1]
		currentDir = changeDir(currentDir, dir)
	}
	if command[2:4] == "ls" {
		currentlyListingContent = true
	}

	return currentDir
}

func changeDir(currentDir *Content, dir string) *Content {
	var contentPointer *Content
	if dir == "/" {
		contentPointer = &baseDir
	} else if dir == ".." {
		if currentDir.parentDir == nil {
			contentPointer = &baseDir
		} else {
			contentPointer = currentDir.parentDir
		}
	} else {
		if val, ok := currentDir.contents[dir]; ok {
			contentPointer = val
		} else {
			content := Content{name: dir, parentDir: currentDir, path: currentDir.path + "/" + dir}
			contentPointer = &content
		}
	}
	return contentPointer
}

func addFileOrDirectoryToContents(dir *Content, entry string) *Content {
	entryName := strings.Split(entry, " ")[1]
	entryDir := Content{name: entryName, parentDir: dir}
	if len(entry) > 3 && entry[0:3] == "dir" {
		entryDir.isDir = true
	} else {
		entryDir.isDir = false
		entryDir.size = stringToInt(strings.Split(entry, " ")[0])
	}
	if dir.contents == nil {
		dir.contents = make(map[string]*Content)
	}
	dir.contents[entryName] = &entryDir

	return dir
}

func stringToInt(stringValue string) int {
	intVal, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

type Content struct {
	name      string
	isDir     bool
	size      int
	path      string
	parentDir *Content
	contents  map[string]*Content
}
