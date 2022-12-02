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

	f, err := os.Open("resources/calories.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	elfNum := 1

	var p ElfList

	for {
		elfCalories, moreToScan, scanErr := getNextCalories(scanner)

		if scanErr != nil {
			log.Fatal(scanErr)
		}

		p = append(p, Elf{elfNum, elfCalories})

		fmt.Printf("Elf: %d, Calories: %d\n", p[p.Len()-1], p[p.Len()-1].calories)

		elfNum++
		if !moreToScan {
			break
		}
	}

	sort.Sort(sort.Reverse(p))

	fmt.Println(p)

	fmt.Printf("The largest calories is: %d\n", p[0].calories)

	fmt.Printf("The total for the three largest calories is: %d\n",
		p[0].calories+p[1].calories+p[2].calories)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getNextCalories(scanner *bufio.Scanner) (int, bool, error) {
	total := 0
	for {
		if !scanner.Scan() {
			return total, false, nil
		}
		text := scanner.Text()
		if text == "" {
			break
		} else {
			intVal, err := strconv.Atoi(text)
			total += intVal

			if err != nil {
				return 0, true, err
			}
		}
	}
	return total, true, nil
}

type Elf struct {
	number   int
	calories int
}

type ElfList []Elf

func (p ElfList) Len() int           { return len(p) }
func (p ElfList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p ElfList) Less(i, j int) bool { return p[i].calories < p[j].calories }
