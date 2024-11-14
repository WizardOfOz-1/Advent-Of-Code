package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func returnRibbon(a, b, c int) int {
	var smallest int

	if a <= b && a <= c {
		smallest = a
	} else if b <= a && b <= c {
		smallest = b
	} else {
		smallest = c
	}

	return smallest

}

func findVolume(a, b, c int) int {
	return a * b * c
}

func convert(a string) (ia int) {

	ia, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	sum := 0
	file, err := os.Open("input")
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		scanned := scanner.Text()
		parts := strings.Split(scanned, "x")
		a, b, c := convert(parts[0]), convert(parts[1]), convert(parts[2])
		sum += findVolume(a, b, c) + returnRibbon(2*(a+b), 2*(b+c), 2*(c+a))
	}
	print(sum)

}
