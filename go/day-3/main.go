package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../../inputs/day3.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanner.Text()
	}
}
