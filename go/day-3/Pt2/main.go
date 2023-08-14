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
		// Get the three elves items
		elf1 := addItems(scanner.Text())
		scanner.Scan()
		elf2 := addItems(scanner.Text())
		scanner.Scan()
		elf3 := addItems(scanner.Text())

		// Loop through elf 1 and see if its items are in both elfs
		for key := range elf1 {
			if elf2[key] && elf3[key] {
				ttlPriority += itemTypePriority(key)
				break
			}
		}
	}
	println(ttlPriority)

}

func addItems(items string) (set map[rune]bool) {
	set = make(map[rune]bool)
	for _, i := range items {
		set[i] = true
	}

	return set
}
