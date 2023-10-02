package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	content, error := os.ReadFile("input.txt")

	if error != nil {
		log.Fatal(error)
	}

	str := string(content)
	data := strings.Split(strings.ReplaceAll(str, "\r\n", "\n"), "\n\n")

	var calories []int

	for _, v := range data {
		var values = strings.Split(v, "\n")

		var totalCalories = 0

		for _, v := range values {

			calories, err := strconv.Atoi(v)

			if err != nil {
				fmt.Print(err)
				continue
			}

			totalCalories += calories
		}

		calories = append(calories, totalCalories)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	fmt.Print(calories[0] + calories[1] + calories[2])

}
