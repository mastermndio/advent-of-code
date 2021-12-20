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

	fmt.Println(input)
}
