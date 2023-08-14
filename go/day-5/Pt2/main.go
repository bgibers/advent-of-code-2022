package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func pop(arr *[]rune, num int) []rune {
	elements := (*arr)[len(*arr)-num:]

	*arr = (*arr)[:len(*arr)-num]

	return elements
}

func main() {
	file, err := os.Open("../../../inputs/day5.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	// Store all slices in an array
	var arr [9][]rune

	for scanner.Text() != " 1   2   3   4   5   6   7   8   9 " {
		for i, r := range scanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				// since we're reading in top to bottom, and we want top to be the last element we need to push to bottom
				arr[i/4] = append([]rune{r}, arr[i/4]...)
			}
		}
		scanner.Scan()
	}

	scanner.Scan()

	for scanner.Scan() {
		var amt, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &amt, &from, &to)

		popped := pop(&arr[from-1], amt)
		arr[to-1] = append(arr[to-1], popped...)
	}

	for _, s := range arr {
		fmt.Print(string(pop(&s, 1)))
	}
}
