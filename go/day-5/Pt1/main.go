package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func pop(arr *[]rune) rune {
	lastElement := (*arr)[len(*arr)-1]

	*arr = (*arr)[:len(*arr)-1]

	return lastElement
}

func main() {
	file, err := os.Open("../../../inputs/day5.txt")

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	/*
		- read in whacky input
			    [D]
			[N] [C]
			[Z] [M] [P]
			1   2   3 .... 9

			move 1 from 2 to 1
			move 3 from 1 to 3
			move 2 from 2 to 1
			move 1 from 1 to 2
		- needs to be read into 9 arrays
		- once we get to steps, push and pop to arrays one at a time
	*/

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

		for i := 0; i < amt; i++ {
			popped := pop(&arr[from-1])
			arr[to-1] = append(arr[to-1], popped)
		}
	}

	for _, s := range arr {
		fmt.Print(string(pop(&s)))
	}
}
