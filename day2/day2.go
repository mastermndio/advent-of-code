package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Function that reads a file line by line and puts each in line into an array and returns it
func inputParse(x string) []string {
	file, err := os.Open(x)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var inputSlice []string
	for scanner.Scan() {

		// Append each line item to a slice/array called inputSlice
		inputSlice = append(inputSlice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputSlice

}

func main() {
	input := inputParse("input.txt")

	fmt.Println(partOne(input))
	fmt.Println(partTwo(input))
}

func partOne(directions []string) int {
	// Create variables to track position, depth, orientation, distance and an array to split our string into
	var directionSplit []string
	horizontalPosition := 0
	depth := 0
	var orientation string
	var distance int

	// Create a loop to iterates through our input(directions slice)
	for i := 0; i < len(directions); i++ {
		// Split element on space ex "forward 2" --> ["forward", 2]
		directionSplit = strings.Split(directions[i], " ")

		orientation = directionSplit[0]
		// converting number string to integer
		distance, _ = strconv.Atoi(directionSplit[1])

		if orientation == "forward" {
			horizontalPosition += distance
		} else if orientation == "up" {
			depth -= distance
		} else {
			depth += distance
		}
	}

	return horizontalPosition * depth
}

func partTwo(directions []string) int {
	var directionSplit []string
	horizontalPosition := 0
	depth := 0
	aim := 0

	var orientation string
	var distance int

	for i := 0; i < len(directions); i++ {
		directionSplit = strings.Split(directions[i], " ")

		orientation = directionSplit[0]
		distance, _ = strconv.Atoi(directionSplit[1])

		if orientation == "forward" {
			horizontalPosition += distance
			depth = depth + (aim * distance)
		} else if orientation == "up" {
			aim -= distance
		} else {
			aim += distance
		}
	}

	return horizontalPosition * depth
}
