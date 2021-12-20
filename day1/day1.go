package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Function that reads a file line by line and puts each in line into an array and returns it
func inputParse(x string) []int {
	file, err := os.Open(x)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	var inputSlice []int
	var inputInt int
	for scanner.Scan() {
		inputInt, _ = strconv.Atoi(scanner.Text())
		inputSlice = append(inputSlice, inputInt)
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

func partOne(input []int) int {
	// created variable to track number of depth increases
	var count int

	//Iterate through each item in our input slice/array
	for i := 0; i < len(input)-1; i++ {
		// Compared current element to next element to see if next element was larger
		if input[i] < input[i+1] {
			// Added 1 to the count variable
			count++
		}
	}

	// returned value of count
	return count

}

func partTwo(input []int) int {
	// created variable to track number of depth increases
	var count int
	var sum1 int
	var sum2 int
	for i := 0; i < len(input)-3; i++ {
		sum1 = input[i] + input[i+1] + input[i+2]
		sum2 = input[i+1] + input[i+2] + input[i+3]

		if sum1 < sum2 {
			count++
		}
	}

	return count
}
