package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const uniquePacketMarkerSize = 14

func main() {
	f, err := os.Open("resources/comms.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		dataString := scanner.Text()
		fmt.Printf("%s\n", dataString)
		for i := 0; i < len(dataString)-(uniquePacketMarkerSize-1); i++ {
			if fourUniqueChars([]rune(dataString[i : i+uniquePacketMarkerSize])) {
				i = i + uniquePacketMarkerSize
				fmt.Printf("%d\n", i)
			}
		}
	}
}

func fourUniqueChars(chars []rune) bool {
	charMap := make(map[rune]bool, 0)
	for _, char := range chars {
		charMap[char] = true
	}
	return len(charMap) == len(chars) && len(chars) == uniquePacketMarkerSize
}
