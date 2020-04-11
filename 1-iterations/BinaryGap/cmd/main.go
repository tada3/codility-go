package main

import (
	solution "BinaryGap"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Start!")

	file, err := os.Open("test-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		l := scanner.Text()
		fmt.Printf("Input: %s\n", l)

		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		answer := solution.Solution(n)

		fmt.Printf("Output: %d\n", answer)
	}
}
