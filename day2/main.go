package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1 := part1()
	part2 := part2()

	fmt.Printf("part1: %v\n", part1)
	fmt.Printf("part2: %v\n", part2)
}

func part1() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	horizontal := 0
	depth := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		d := parts[0]
		l, _ := strconv.Atoi(parts[1])
		switch d {
		case "forward":
			horizontal += l
		case "up":
			depth -= l
		case "down":
			depth += l
		default:
			log.Fatal("unknown direction!")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := horizontal * depth
	return sum
}

func part2() int {
	file, _ := os.Open("input.txt")
	defer file.Close()

	horizontal := 0
	depth := 0
	aim := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		d := parts[0]
		l, _ := strconv.Atoi(parts[1])
		switch d {
		case "forward":
			horizontal += l
			depth += aim * l
		case "up":
			aim -= l
		case "down":
			aim += l
		default:
			log.Fatal("unknown direction!")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := horizontal * depth
	return sum
}
