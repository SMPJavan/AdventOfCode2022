package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const stackCount = 9
const stackParseTxtGap = 4
const spaceChar = " "

func main() {
	f, err := os.Open("resources/crane.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	stacks := loadInitialStacks(parseInitialStackLines(scanner))

	for scanner.Scan() {
		crateCount, moveFromStack, moveToStack, successfullyParsed := parseMoveCommand(scanner.Text())
		if successfullyParsed {
			stacks = moveCrates(stacks, crateCount, moveFromStack, moveToStack)
		}
	}

	topRow := make([]string, 0)

	for i := 1; i <= stackCount; i++ {
		topRow = append(topRow, stacks[i][len(stacks[i])-1])
	}

	fmt.Printf("Final stack top row: %s", topRow)
}

func parseInitialStackLines(scanner *bufio.Scanner) []string {
	initialLines := make([]string, 0)

	for scanner.Scan() {
		lineContent := scanner.Text()
		if strings.Trim(lineContent, " ")[0:1] != "[" {
			break
		}
		initialLines = append(initialLines, lineContent)
	}

	return reverseStackOrder(initialLines)
}

func loadInitialStacks(initialLines []string) map[int][]string {
	initialStacks := make(map[int][]string, 9)

	for _, line := range initialLines {
		for i := 1; i <= stackCount; i++ {
			textIndex := 1 + ((i - 1) * stackParseTxtGap)
			crate := string(line[textIndex])
			if crate != spaceChar {
				initialStacks[i] = append(initialStacks[i], crate)
			}
		}
	}

	return initialStacks
}

func reverseStackOrder(stack []string) []string {
	reversedStack := make([]string, 0)
	for i := len(stack) - 1; i >= 0; i-- {
		reversedStack = append(reversedStack, stack[i])
	}

	return reversedStack
}

func parseMoveCommand(command string) (int, int, int, bool) {
	split1 := strings.Split(command, "move ")
	if len(split1) > 1 {
		split2 := strings.Split(split1[1], " from ")
		split3 := strings.Split(split2[1], " to ")
		crateCount := split2[0]
		moveFromStack := split3[0]
		moveToStack := split3[1]
		return stringToInt(crateCount), stringToInt(moveFromStack), stringToInt(moveToStack), true
	}
	return 0, 0, 0, false
}

func moveCrates(stacks map[int][]string, crateCount int, fromStack int, toStack int) map[int][]string {
	fromStackIndex := len(stacks[fromStack]) - (crateCount)
	stacks[toStack] = append(stacks[toStack], stacks[fromStack][fromStackIndex:len(stacks[fromStack])]...)
	stacks[fromStack] = stacks[fromStack][:fromStackIndex]
	return stacks
}

func stringToInt(stringValue string) int {
	intVal, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}
