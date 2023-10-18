package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, error := os.ReadFile("input.txt")

	if error != nil {
		log.Fatal(error)
	}

	instructions := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	stacks := [][] string {
		{"S", "L", "W"}, // 1
		{"J", "T", "N", "Q"}, // 2
		{"S", "C", "H", "F", "J"}, // 3
		{"T", "R", "M", "W", "N", "G", "B"}, // 4
		{"T", "R", "L", "S", "D", "H", "Q", "B"}, // 5
		{"M", "J", "B", "V", "F", "H", "R", "L"}, // 6
		{"D", "W", "R", "N", "J", "M"}, // 7
		{"B", "Z", "T", "F", "H", "N", "D", "J"}, // 8
		{"H", "L", "Q", "N", "B", "F", "T"}, // 9
	}

	finalCargos := []string{}

	for _, instruction := range instructions {

		amount, _ := strconv.Atoi(strings.Split(instruction, " ")[1])
		originIndex, _ := strconv.Atoi(strings.Split(instruction, " ")[3])
		originIndex = originIndex -1
		targetIndex, _ := strconv.Atoi(strings.Split(instruction, " ")[5])
		targetIndex = targetIndex -1

		transfers := []string{}

		for i := 0; i < amount; i++ {

			origin := stacks[originIndex];
			finalCargo := origin[len(origin)-1]

			transfers = append(transfers, finalCargo)	
			stacks[originIndex] = origin[0:len(origin)-1]
		}

		for _, cargo := range transfers {
			stacks[targetIndex] = append(stacks[targetIndex], string(cargo))
		}
	}

	for _, stack := range stacks {
		finalCargos = append(finalCargos, stack[len(stack)- 1])
	}

	fmt.Println(finalCargos)
}