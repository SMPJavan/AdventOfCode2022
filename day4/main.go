package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("resources/assignments.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	//fullyContainedZones := 0
	overlappingZones := 0

	for scanner.Scan() {

		assignmentZones := strings.Split(scanner.Text(), ",")
		currentLineZones1 := assignmentZoneRangeToArray(assignmentZones[0])
		currentLineZones2 := assignmentZoneRangeToArray(assignmentZones[1])
		/*		if isAssignmentFullyInsideOtherAssignment(currentLineZones1, currentLineZones2) || isAssignmentFullyInsideOtherAssignment(currentLineZones2, currentLineZones1) {
				fullyContainedZones = fullyContainedZones + 1
			}*/
		if assignmentsOverlap(currentLineZones1, currentLineZones2) {
			overlappingZones = overlappingZones + 1
		}
	}

	//fmt.Printf("Total fully enclosed zones: %d\n", fullyContainedZones)
	fmt.Printf("Total overlapping zones: %d\n", overlappingZones)
}

func assignmentZoneRangeToArray(assignmentZonesString string) []int {
	assignmentZoneArray := make([]int, 0)

	assignmentZoneBoundaries := strings.Split(assignmentZonesString, "-")

	for i := stringToInt(assignmentZoneBoundaries[0]); i <= stringToInt(assignmentZoneBoundaries[1]); i++ {
		assignmentZoneArray = append(assignmentZoneArray, i)
	}

	return assignmentZoneArray
}

func stringToInt(stringValue string) int {
	intVal, err := strconv.Atoi(stringValue)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

func assignmentsOverlap(assignment1 []int, assignment2 []int) bool {
	assignmentFullyInsideOtherAssignment := false
	for _, assignment1Zone := range assignment1 {
		for _, assignment2Zone := range assignment2 {
			if assignment1Zone == assignment2Zone {
				assignmentFullyInsideOtherAssignment = true
				break
			}
			if assignmentFullyInsideOtherAssignment {
				break
			}
		}
	}
	return assignmentFullyInsideOtherAssignment
}

func isAssignmentFullyInsideOtherAssignment(assignment1 []int, assignment2 []int) bool {
	if len(assignment1) > len(assignment2) {
		return false
	}
	assignmentFullyInsideOtherAssignment := true
	for _, assignment1Zone := range assignment1 {
		zoneMatched := false
		for _, assignment2Zone := range assignment2 {
			if assignment1Zone == assignment2Zone {
				zoneMatched = true
				break
			}
		}
		if !zoneMatched {
			assignmentFullyInsideOtherAssignment = false
			break
		}
	}
	return assignmentFullyInsideOtherAssignment
}
