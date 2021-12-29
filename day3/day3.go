package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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
		inputSlice = append(inputSlice, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return inputSlice

}

func dayOne(list []string) float64 {
	var gammaString string
	var epsilonString string
	zeros := 0
	ones := 0
	for bit := 0; bit < len(list[0]); bit++ {
		zeros = 0
		ones = 0
		for item := 0; item < len(list); item++ {
			switch string(list[item][bit]) {
			case "0":
				zeros++
			case "1":
				ones++
			}
		}

		if zeros > ones {
			gammaString += "0"
			epsilonString += "1"
		} else {
			gammaString += "1"
			epsilonString += "0"
		}

	}

	gamma := binaryConvert(gammaString)
	epsilon := binaryConvert(epsilonString)

	return gamma * epsilon
}

func dayTwo(list []string, rating string) float64 {
	filteredList := list

	zeros := 0
	ones := 0
	for bit := 0; bit < len(filteredList[0]); bit++ {
		zeros = 0
		ones = 0
		for item := 0; item < len(filteredList); item++ {
			if string(filteredList[item][bit]) == "0" {
				zeros++
			} else {
				ones++
			}
		}

		if rating == "oxygen" {
			if zeros > ones {
				filteredList = filter(filteredList, "0", bit)
			} else {
				filteredList = filter(filteredList, "1", bit)
			}
		} else {
			if zeros > ones {
				filteredList = filter(filteredList, "1", bit)
			} else {
				filteredList = filter(filteredList, "0", bit)
			}
		}

	}

	fmt.Println(filteredList[0])
	result := binaryConvert(filteredList[0])

	return result
}

func filter(list []string, criteria string, index int) []string {
	var filteredList []string
	for i := 0; i < len(list); i++ {
		if len(list) == 1 {
			filteredList = list
			break
		}
		if string(list[i][index]) == criteria {
			filteredList = append(filteredList, list[i])
		}
	}

	return filteredList
}

func binaryConvert(binaryString string) float64 {
	var num int
	var decimal float64
	i := 0
	for digit := len(binaryString) - 1; digit >= 0; digit-- {
		num, _ = strconv.Atoi(string(binaryString[digit]))
		decimal += float64(num) * math.Pow(2, float64(i))
		i++
	}

	return decimal
}

func main() {
	input := inputParse("input.txt")

	result1 := dayOne(input)

	fmt.Println(result1)

	oxygen := dayTwo(input, "oxygen")
	c02 := dayTwo(input, "c02")

	var total int64
	total = int64(oxygen) * int64(c02)
	fmt.Println(total)
}
