package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../../../inputs/day4.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	overlapCount := 0

	for scanner.Scan() {
		/*
			- Each line contains a pair 2-4,6-8
			- looking for any overlap at all
			- if  p1x <= p2y && p2x <= p1y overlap
		*/

		line := scanner.Text()

		var p1x, p1y, p2x, p2y int
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &p1x, &p1y, &p2x, &p2y)

		if err != nil {
			log.Fatalf("failed to parse line '%s': %s", line, err)
		}

		if p1x <= p2y && p2x <= p1y {
			overlapCount++
		}
	}

	println(overlapCount)
}
