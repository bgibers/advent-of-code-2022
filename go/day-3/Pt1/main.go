package main

import (
	"bufio"
	"log"
	"os"
)

func itemTypePriority(item rune) int {
	if 'a' <= item && item <= 'z' {
		return int(item - 'a' + 1)
	}
	if 'A' <= item && item <= 'Z' {
		return int(item - 'A' + 27)
	}
	return 0 // Invalid item type
}

func main() {
	file, err := os.Open("../../../inputs/day3.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ttlPriority := 0

	for scanner.Scan() {
		items := make(map[rune]bool)

		for _, i := range scanner.Text()[:len(scanner.Text())/2] {
			items[i] = true
		}

		for _, i := range scanner.Text()[len(scanner.Text())/2:] {
			if items[i] {
				ttlPriority += itemTypePriority(i)
				break
			}
		}
	}
	println(ttlPriority)

}
