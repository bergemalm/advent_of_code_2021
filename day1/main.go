package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	numbers := make([]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}

	part1 := part1(numbers)
	fmt.Printf("part1: %v\n", part1)

	part2 := part2(numbers)
	fmt.Printf("part2: %v\n", part2)
}

func part1(numbers []int) int {
	n := 0
	var previous int
	for i, number := range numbers {
		if i > 0 && number > previous {
			n++
		}
		previous = number
	}

	return n
}

func part2(numbers []int) int {
	n := 0
	var previous int
	for i := 0; i < len(numbers)-2; i++ {
		slice := numbers[i : i+3]
		sum := 0
		for _, nr := range slice {
			sum += nr
		}
		if i > 0 && sum > previous {
			n++
		}
		previous = sum
	}

	return n
}
