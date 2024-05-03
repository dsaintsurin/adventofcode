package day1

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
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
	str2digit := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	digits := "0123456789"
	strFIdx := math.MaxInt
	fNum := -1
	strLIdx := -math.MaxInt
	lNum := -1
	for key := range str2digit {
		if i := strings.Index(str, key); i < strFIdx && i != -1 {
			strFIdx = i
			fNum = str2digit[key]

		}
		if i := strings.LastIndex(str, key); i > strLIdx && i != -1 {
			strLIdx = i
			lNum = str2digit[key]
		}
	}

	numFIdx := strings.IndexAny(str, digits)
	if numFIdx < strFIdx {
		fNum, _ = strconv.Atoi(fmt.Sprintf("%c", str[numFIdx]))
	}

	numLIdx := strings.LastIndexAny(str, digits)

	if numLIdx > strLIdx {
		lNum, _ = strconv.Atoi(fmt.Sprintf("%c", str[numLIdx]))
	}
	return 10*fNum + lNum
}
