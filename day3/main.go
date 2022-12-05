package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	part2()
}

func part2() {

	f, err := os.Open("resources/rucksacks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPriority := 0

	for {
		elves, moreToScan := getThreeElfGroup(scanner)
		if !moreToScan {
			break
		}
		for _, elf := range elves {
			fmt.Printf("Elf: %s\n", elf)
		}
		commonCharsForLine := commonCharsIn3Maps(uniqueCharMap(elves[0]), uniqueCharMap(elves[1]), uniqueCharMap(elves[2]))

		for _, char := range commonCharsForLine {
			priority := intValueForChar(char)
			totalPriority = totalPriority + priority
			fmt.Printf("Common char value: %s\n", string(char))
		}
	}
	fmt.Printf("Total priority value: %d\n", totalPriority)
}
func part1() {

	f, err := os.Open("resources/rucksacks.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPriority := 0

	for scanner.Scan() {
		text := scanner.Text()
		compartmentSize := len(text) / 2
		firstCompartment := text[0:compartmentSize]
		secondCompartment := text[compartmentSize:]
		fmt.Printf("First: %s  |  Second: %s\n", firstCompartment, secondCompartment)
		commonCharsForLine := commonChars(firstCompartment, secondCompartment)

		for _, char := range commonCharsForLine {
			priority := intValueForChar(char)
			totalPriority = totalPriority + priority
			fmt.Printf("Common char value: %d\n", priority)
		}

		fmt.Printf("Common Chars: %s\n", string(commonCharsForLine))
	}
	fmt.Printf("Total priority: %d\n", totalPriority)
}

func commonChars(inputStr1 string, inputString2 string) []rune {
	return commonCharsInMaps(uniqueCharMap(inputStr1), uniqueCharMap(inputString2))
}

func commonCharsInMaps(inputStr1 map[rune]bool, inputString2 map[rune]bool) []rune {
	commonChars := make([]rune, 0)

	for key := range inputStr1 {
		if _, ok := inputString2[key]; ok {
			commonChars = append(commonChars, key)
		}
	}

	return commonChars
}

func commonCharsIn3Maps(inputStr1 map[rune]bool, inputString2 map[rune]bool, inputString3 map[rune]bool) []rune {
	commonChars := make([]rune, 0)

	for key := range inputStr1 {
		if ok := inputString2[key] && inputString3[key]; ok {
			commonChars = append(commonChars, key)
		}
	}

	return commonChars
}

func uniqueCharMap(inputStr string) map[rune]bool {
	chars := []rune(inputStr)
	boolMap := make(map[rune]bool)

	for _, val := range chars {
		if _, ok := boolMap[val]; !ok {
			boolMap[val] = true
		}
	}

	return boolMap
}

func intValueForChar(char rune) int {
	var intVal int
	if unicode.IsUpper(char) {
		intVal = int(char) - int('A') + 27
	} else {
		intVal = int(char) - int('a') + 1
	}
	return intVal
}

func getThreeElfGroup(scanner *bufio.Scanner) ([]string, bool) {
	elves := make([]string, 0)
	moreToScan := true

	for i := 1; i <= 3; i++ {
		if !scanner.Scan() {
			moreToScan = false
			break
		}
		elves = append(elves, scanner.Text())
	}

	return elves, moreToScan
}
