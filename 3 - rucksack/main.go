package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	content, error := os.ReadFile("input.txt")

	if error != nil {
		log.Fatal(error)
	}

	rucksacks := strings.Split(strings.ReplaceAll(string(content), "\r\n", "\n"), "\n")

	totalSum := 0;

	groupSize := 3;
	groupEnd := 0;
	for groupStart := 0; groupStart < len(rucksacks); groupStart += groupSize {
		groupEnd += groupSize
		if groupEnd > len(rucksacks) {
			groupEnd = len(rucksacks)
		}

		group := rucksacks[groupStart:groupEnd]

		uniqueCharacters := ""

		characters := group[0]

		for _, char := range characters {
			if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)) && !strings.Contains(uniqueCharacters, string(char)) {
				uniqueCharacters += string(char)
			}
		}

		for _, character := range uniqueCharacters {
	
			if (unicode.IsUpper(character)) {
				totalSum += int(character - 'A' + 1) + 26
			} else {
				totalSum += int(character - 'a' + 1) 
			}
		}
	}

	fmt.Println(totalSum);
}