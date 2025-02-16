package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func sortSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}

func GetSlicesFromInput(input string) ([]int, []int) {
	file, err := os.Open("./input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var leftSlice []int
	var rightSlice []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		slice := strings.Split(line, " ")

		leftInt, _ := strconv.Atoi(slice[0])
		rightInt, _ := strconv.Atoi(slice[1])

		leftSlice = append(leftSlice, leftInt)
		rightSlice = append(rightSlice, rightInt)
	}

	return leftSlice, rightSlice
}

func main() {
	leftSlice, rightSlice := GetSlicesFromInput("input")

	resultForSecondPart := SolutionSecondPart(leftSlice, rightSlice)

	fmt.Println(resultForSecondPart)

	sortSlice(leftSlice)
	sortSlice(rightSlice)

	var sum int

	for i := 0; i < len(leftSlice); i++ {
		distance := abs(leftSlice[i] - rightSlice[i])
		sum += distance
	}

	fmt.Println(sum)
}

func SolutionSecondPart(firstSlice, secondSlice []int) int {
	var result int

	storage := make(map[int]int)

	for _, number := range firstSlice {
		for j := 0; j < len(secondSlice); j++ {
			if number == secondSlice[j] {
				storage[number] += 1
			}
		}

		result += number * storage[number]
		storage[number] = 0
	}

	return result
}
