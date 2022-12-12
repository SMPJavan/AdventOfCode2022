package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var cycleCount = 0
var xRegistry = 1
var signalStrengthSum = 0

var pixels = make(map[int]map[int]bool, 6)

func main() {

	for i := 0; i < 6; i++ {
		pixels[i] = make(map[int]bool, 40)
		for j := 0; j < 40; j++ {
			pixels[i][j] = false
		}
	}

	f, err := os.Open("resources/operations.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	fmt.Printf("\nStart|\n")
	for scanner.Scan() {
		operation := scanner.Text()
		fmt.Printf("operation: %s\n", operation)
		runOperation(operation)
	}
	printScreen()
}

func printScreen() {
	for i := 0; i < len(pixels); i++ {
		for j := 0; j < 40; j++ {
			if pixels[i][j] {
				fmt.Printf("# ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func runOperation(operation string) {
	operationParts := strings.Split(operation, " ")
	if operationParts[0] == "addx" {
		size, _ := strconv.Atoi(operationParts[1])
		addX(size)
	} else {
		cycle()
	}
}

func addX(size int) {
	cycle()
	cycle()
	xRegistry = xRegistry + size
}

func cycle() {
	cycleCount++
	signalStrength := cycleCount * xRegistry
	//fmt.Printf("Cycle: %d, xRegistry: %d, signalStrength: %d\n", cycleCount, xRegistry, signalStrength)
	if isSummableCycle(cycleCount) {
		signalStrengthSum = signalStrengthSum + signalStrength
		//fmt.Printf("Cycle:%d, SignalStrengthSum: %d\n", cycleCount, signalStrengthSum)
	}
	enablePixels()
}

func enablePixels() {
	row := cycleCount / 40
	xPos := cycleCount % 40
	if xPos > 0 && xPos >= xRegistry && xPos <= xRegistry+2 {
		pixels[row][xPos-1] = true
	}
}

func isSummableCycle(cycle int) bool {
	switch cycle {
	case
		20,
		60,
		100,
		140,
		180,
		220:
		return true
	}
	return false
}
