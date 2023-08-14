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

	numPairs := 0

	for scanner.Scan() {
		/*
			- Each line contains a pair 2-4,6-8
			- Need to see if one pair fully contains the other. 2-8,3-7
			- If p1x <= p2x && p1y >= p2y || p1x >= p2x && p1y <= p2y
		*/

		line := scanner.Text()

		var p1x, p1y, p2x, p2y int
		_, err := fmt.Sscanf(line, "%d-%d,%d-%d", &p1x, &p1y, &p2x, &p2y)

		if err != nil {
			log.Fatalf("failed to parse line '%s': %s", line, err)
		}

		if p1x <= p2x && p1y >= p2y {
			numPairs++
		}
	}

	println(numPairs)
}
