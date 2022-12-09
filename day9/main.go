package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var headPosX int
var headPosY int
var tailPosX int
var tailPosY int
var positionsVisited map[string]bool

func main() {

	f, err := os.Open("resources/knots.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	headPosX = 0
	headPosY = 0
	tailPosX = 0
	tailPosY = 0
	positionsVisited = make(map[string]bool, 0)

	fmt.Printf("\nStart|\n")
	drawState()
	for scanner.Scan() {
		headMovement := scanner.Text()
		fmt.Printf("\ncommand: %s\n ", headMovement)
		commandParts := strings.Split(headMovement, " ")
		distance, _ := strconv.Atoi(string(commandParts[1]))
		for i := 0; i < distance; i++ {
			if commandParts[0] == "R" {
				headPosX = moveHead(headPosX, true)
			}
			if commandParts[0] == "L" {
				headPosX = moveHead(headPosX, false)
			}
			if commandParts[0] == "U" {
				headPosY = moveHead(headPosY, true)
			}
			if commandParts[0] == "D" {
				headPosY = moveHead(headPosY, false)
			}
			updateTailPos()
			drawState()
		}
		fmt.Printf("HeadPosX: %d, HeadPosY: %d| TailPosX: %d, TailPosY: %d\n", headPosX, headPosY, tailPosX, tailPosY)
	}
	fmt.Printf(strconv.Itoa(len(positionsVisited)))
}

func drawState() {
	for y := 5; y > -5; y-- {
		for x := -10; x < 10; x++ {
			if tailPosX == x && tailPosY == y {
				fmt.Printf("T")
			} else if headPosX == x && headPosY == y {
				fmt.Printf("H")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}

func updateTailPos() {
	differenceX := headPosX - tailPosX
	differenceY := headPosY - tailPosY
	if differenceX > 1 {
		tailPosX = tailPosX + 1
		if differenceY > 0 {
			tailPosY = tailPosY + 1
		} else if differenceY < 0 {
			tailPosY = tailPosY - 1
		}
	} else if differenceX < -1 {
		tailPosX = tailPosX - 1
		if differenceY > 0 {
			tailPosY = tailPosY + 1
		} else if differenceY < 0 {
			tailPosY = tailPosY - 1
		}
	} else if differenceY > 1 {
		tailPosY = tailPosY + 1
		if differenceX > 0 {
			tailPosX = tailPosX + 1
		} else if differenceX < 0 {
			tailPosX = tailPosX - 1
		}
	} else if differenceY < -1 {
		tailPosY = tailPosY - 1
		if differenceX > 0 {
			tailPosX = tailPosX + 1
		} else if differenceX < 0 {
			tailPosX = tailPosX - 1
		}
	}
	positionsVisited[strconv.Itoa(tailPosX)+","+strconv.Itoa(tailPosY)] = true
}

func moveHead(headPos int, increment bool) int {
	if increment {
		headPos = headPos + 1
	} else {
		headPos = headPos - 1
	}
	return headPos
}
