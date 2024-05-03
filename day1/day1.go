package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Solution(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		inputStr := scanner.Text()
		sum += helper(inputStr)
	}
	fmt.Println(sum)
}

func helper(str string) int {
	var c1 rune
	var c2 rune
	for _, c := range str {
		if unicode.IsDigit(c) {
			if c1 == 0 {
				c1 = c
			}
			c2 = c
		}
	}

	k := fmt.Sprintf("%c%c", c1, c2)
	num, _ := strconv.Atoi(k)
	return num
}
