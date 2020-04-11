package main

import (
	solution "CyclicRotation"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

		l1 := unpack(l)

		if l1 == "" {
			fmt.Println("Empty input, ignored.")
			continue
		}

		r1 := []rune(l1)

		a, pos := readIntArray(r1, 0)
		k, _ := readInt(r1, pos)

		fmt.Printf("a=%+v\n", a)
		fmt.Printf("k=%d\n", k)

		answer := solution.Solution(a, k)

		fmt.Printf("Output: %d\n", answer)
	}
}

func unpack(s string) string {
	return strings.Trim(s, " \t()")
}

func readIntArray(r []rune, pos int) ([]int, int) {
	start := 0
	end := 0
	i := pos
	for ; i < len(r); i++ {
		if r[i] == ' ' {
			continue
		}
		if r[i] == '[' {
			start = i + 1
		} else if r[i] == ']' {
			end = i
			break
		}
	}
	r1 := r[start:end]
	aStr := strings.Split(string(r1), ",")
	a := make([]int, len(aStr))
	for i := 0; i < len(aStr); i++ {
		x, err := strconv.Atoi(strings.TrimSpace(aStr[i]))
		if err != nil {
			panic(err)
		}
		a[i] = x
	}

	next := nextComma(r, end)
	if next < len(r) {
		next++
	}
	return a, next
}

func readInt(r []rune, pos int) (int, int) {
	start := pos
	next := nextComma(r, pos)
	s1 := string(r[start:next])
	x, err := strconv.Atoi(strings.TrimSpace(s1))
	if err != nil {
		panic(err)
	}

	if next < len(r) {
		next++
	}
	return x, next
}

func nextComma(r []rune, pos int) int {
	i := pos
	for ; i < len(r); i++ {
		if r[i] == ',' {
			return i
		}
	}
	return i
}
