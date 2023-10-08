package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

/* TODO:
- Split string on half of character length
- Find values that match between compartments
- Map priority to lowercase letters and  upercase letters 1 - 26 + 27 - 52
- Return priority of selected values
- Collect sum of all values in input

*/


func main() {
	content, error := os.ReadFile("input.txt")

	if error != nil {
		log.Fatal(error)
	}

	rucksacks := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	totalSum := 0;
	for _, rucksack := range rucksacks {

		length := len(rucksack)
		compartmentA := rucksack[:length/2]
		compartmentB := rucksack[length/2:length]

		matchingCharacters := ""

		for _, character := range compartmentA {

			if (strings.Contains(compartmentB, string(character)) && !strings.Contains(matchingCharacters, string(character))) {
				matchingCharacters += string(character)
			}
		}

		if len(matchingCharacters) == 0 {
			continue
		}

		for _, character := range matchingCharacters {

			if (unicode.IsUpper(character)) {
				totalSum += int(character - 'A' + 1) + 26
			} else {
				totalSum += int(character - 'a' + 1) 
			}
		}
	}

	fmt.Println(totalSum);
}