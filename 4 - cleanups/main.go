package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)
func main() {
	content, error := os.ReadFile("input.txt")

	if error != nil {
		log.Fatal(error)
	}

	assignments := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	overlappingPairs := 0
	for _, assignment := range assignments {

		assignmentA := strings.Split(assignment, ",")[0]
		assignmentB := strings.Split(assignment, ",")[1]

		startA, _ := strconv.Atoi(strings.Split(assignmentA, "-")[0])
		endA, _ := strconv.Atoi(strings.Split(assignmentA, "-")[1])

		startB, _ := strconv.Atoi(strings.Split(assignmentB, "-")[0])
		endB, _ := strconv.Atoi(strings.Split(assignmentB, "-")[1])

		var sectionsA []int
		var sectionsB []int

		for i := startA; i <  endA +1; i++ {
			sectionsA = append(sectionsA, i)
		}

		for i := startB; i <  endB +1; i++ {
			sectionsB = append(sectionsB, i)
		}

		pairOverlaps := false

		for _, sectionAId := range sectionsA {
			if (pairOverlaps == false && slices.Contains(sectionsB, sectionAId)) {
				pairOverlaps = true
				overlappingPairs++
			}
		}
	}

	fmt.Println(overlappingPairs)
}