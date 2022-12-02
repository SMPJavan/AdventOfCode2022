package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var shapes = setupShapes()

func main() {

	f, err := os.Open("resources/strategy.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalPoints := 0

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		opponent := s[0]
		result := s[1]
		shapeToPick := calculateSelectionForResult(opponent, result)
		fmt.Printf("Shape to pick for line (%s): %s\n", text, shapeToPick)
		lineScore := shapes[shapeToPick].points + calculatePointsForResult(opponent, shapeToPick)
		fmt.Printf("Points for line (%s %s): %d\n", opponent, shapeToPick, lineScore)
		totalPoints += lineScore
	}

	fmt.Printf("Total points: %d\n", totalPoints)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateSelectionForResult(opponent string, result string) string {
	var selection string
	switch result {
	case "X":
		selection = shapes[opponent].beats
	case "Y":
		selection = opponent
	case "Z":
		selection = shapes[opponent].isBeatenBy
	default:
		selection = ""
	}

	return selection
}

func calculatePointsForResult(opponent string, mine string) int {
	points := 0

	if isDraw(opponent, mine) {
		points = 3
	} else if isVictory(opponent, mine) {
		points = 6
	}

	return points
}

func isDraw(opponent string, mine string) bool {
	return opponent == mine
}

func isVictory(opponent string, mine string) bool {
	return shapes[mine].beats == opponent
}

func setupShapes() map[string]Shape {
	shapes := make(map[string]Shape)
	shapes["A"] = Shape{beats: "C", isBeatenBy: "B", points: 1}
	shapes["B"] = Shape{beats: "A", isBeatenBy: "C", points: 2}
	shapes["C"] = Shape{beats: "B", isBeatenBy: "A", points: 3}
	return shapes
}

type Shape struct {
	beats      string
	isBeatenBy string
	points     int
}
