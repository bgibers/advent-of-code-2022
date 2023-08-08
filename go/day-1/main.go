package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func calCounts(filePath string) []int {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	var allElvesCals []int
	var currentCalCount int

	// Iterate through each line using the scanner
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			cals, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("bad item %q: %v", line, err)
			}

			currentCalCount += cals
		} else {
			allElvesCals = append(allElvesCals, currentCalCount)
			currentCalCount = 0
		}
	}

	allElvesCals = append(allElvesCals, currentCalCount)

	return allElvesCals
}

func main() {
	cals := calCounts("../../inputs/day1.txt")

	sort.Sort(sort.Reverse(sort.IntSlice(cals)))

	topThree := cals[:3]

	sumTopThree := 0
	for _, val := range topThree {
		sumTopThree += val
	}

	fmt.Printf("Top 3 values %d with a total sum of %d", topThree, sumTopThree)
}
