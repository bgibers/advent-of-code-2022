package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	scores := map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}

	totalScore := 0

	file, err := os.Open("../../inputs/day2.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		totalScore += scores[scanner.Text()]
	}

	println(totalScore)
}
