package main

import (
	"bufio"
	"os"
)

func readFile() int {
	story, pos := 0, 1
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

		if story < 0 {
			print("Reached -1 at ", pos)
			break
		}
        pos++
	}
	return story
}
func main() {
	res := readFile()
	print(res)

}
