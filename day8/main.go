package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("resources/trees.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	rows := make(map[int]map[int]int, 0)
	columns := make(map[int]map[int]int, 0)

	scanner := bufio.NewScanner(f)

	rowCount := 0

	for scanner.Scan() {
		treeRow := scanner.Text()
		fmt.Printf("row: %s\n", treeRow)
		rows[rowCount] = stringToIntMap(treeRow)
		for index, tree := range rows[rowCount] {
			if columns[index] == nil {
				columns[index] = make(map[int]int, 0)
			}
			columns[index][rowCount] = tree
		}
		rowCount = rowCount + 1
	}

	visibilityValues := make([]int, 0)
	for row, _ := range rows {
		for column, _ := range columns {
			visibilityValue := lookLeft(rows, row, column) * lookRight(rows, row, column) * lookDown(rows, row, column) * lookUp(rows, row, column)
			fmt.Printf("Visibility value for %d, %d: %d\n", row, column, visibilityValue)
			visibilityValues = append(visibilityValues, visibilityValue)
		}
	}

	sort.Ints(visibilityValues)

	fmt.Printf("Largest visibility value: %d\n", visibilityValues[len(visibilityValues)-1])
}

func lookLeft(rows map[int]map[int]int, row int, column int) int {
	treeVal := rows[row][column]
	visibleTrees := 0
	for i := column - 1; i >= 0; i-- {
		visibleTrees = visibleTrees + 1
		if rows[row][i] >= treeVal {
			break
		}
	}
	return visibleTrees
}

func lookRight(rows map[int]map[int]int, row int, column int) int {
	treeVal := rows[row][column]
	visibleTrees := 0
	for i := column + 1; i < len(rows[row]); i++ {
		visibleTrees = visibleTrees + 1
		if rows[row][i] >= treeVal {
			break
		}
	}
	return visibleTrees
}

func lookDown(rows map[int]map[int]int, row int, column int) int {
	treeVal := rows[row][column]
	visibleTrees := 0
	for i := row + 1; i < len(rows); i++ {
		visibleTrees = visibleTrees + 1
		if rows[i][column] >= treeVal {
			break
		}
	}
	return visibleTrees
}

func lookUp(rows map[int]map[int]int, row int, column int) int {
	treeVal := rows[row][column]
	visibleTrees := 0
	for i := row - 1; i >= 0; i-- {
		visibleTrees = visibleTrees + 1
		if rows[i][column] >= treeVal {
			break
		}
	}
	return visibleTrees
}

func stringToIntMap(intString string) map[int]int {
	intArray := make(map[int]int, len(intString))
	for index, char := range intString {
		intArray[index], _ = strconv.Atoi(string(char))
	}
	return intArray
}

func getPositionsOfVisibleTrees(keys []int, treeRowOrColumn map[int]int) []int {
	visibleTreePositionsInGroup := make([]int, 0)
	largestTreeSoFar := 0
	for i, index := range keys {
		tree := treeRowOrColumn[index]
		if i == 0 || tree > largestTreeSoFar {
			visibleTreePositionsInGroup = append(visibleTreePositionsInGroup, index)
			largestTreeSoFar = tree
		}
	}
	return visibleTreePositionsInGroup
}
