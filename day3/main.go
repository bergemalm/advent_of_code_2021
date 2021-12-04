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
	file, _ := os.Open("input.txt")
	defer file.Close()

	binaries := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bin := strings.TrimSpace(scanner.Text())
		binaries = append(binaries, bin)
	}

	part1 := part1(binaries)
	fmt.Printf("part1: %v\n", part1)
	part2 := part2(binaries)
	fmt.Printf("part2: %v\n", part2)
}

func part1(binaries []string) int64 {
	gamma, epsilon := getByPosition(binaries)
	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)

	return gammaInt * epsilonInt
}

func part2(binaries []string) int64 {
	o2 := getRating(binaries, "o2")
	co2 := getRating(binaries, "co2")
	o2Int, _ := strconv.ParseInt(o2, 2, 64)
	co2Int, _ := strconv.ParseInt(co2, 2, 64)

	return o2Int * co2Int
}

func getByPosition(binaries []string) (string, string) {
	var gamma, epsilon string
	for i := 0; i < len(binaries[0]); i++ {
		var zeroCount, oneCount int
		for _, binary := range binaries {
			switch binary[i] {
			case '0':
				zeroCount++
			case '1':
				oneCount++
			default:
				log.Fatalf("neither 0 nor 1!")
			}
		}
		if zeroCount > oneCount {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	return gamma, epsilon
}

func isPos(binaries []string, pos int, d string) []string {
	filtered := make([]string, 0)
	for _, n := range binaries {
		if string(n[pos]) == d {
			filtered = append(filtered, n)
		}
	}
	if len(filtered) > 0 {
		return filtered
	}

	return binaries
}

func getRating(binaries []string, f string) string {
	for i := 0; i < len(binaries[0]); i++ {
		var zeroCount, oneCount int
		for _, binary := range binaries {
			switch binary[i] {
			case '0':
				zeroCount++
			case '1':
				oneCount++
			default:
				log.Fatalf("neither 0 nor 1!")
			}
		}
		switch f {
		case "o2":
			if oneCount >= zeroCount {
				binaries = isPos(binaries, i, "1")
			} else {
				binaries = isPos(binaries, i, "0")
			}
		case "co2":
			if oneCount < zeroCount {
				binaries = isPos(binaries, i, "1")
			} else {
				binaries = isPos(binaries, i, "0")
			}
		}
	}

	return binaries[0]
}
