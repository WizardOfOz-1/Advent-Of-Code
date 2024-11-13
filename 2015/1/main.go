package main

import (
	"bufio"
	"os"
)

func readFile() int {
	story := 0
	file, err := os.Open("input")
	if err != nil {
		panic("Couldnt open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		char := scanner.Text()
		if char == "(" {
			story++
		} else if char == ")" {
			story--
		}
	}
	return story
}
func main() {
    res := readFile()
    print(res)

}
