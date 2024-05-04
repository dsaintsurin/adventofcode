package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cubes struct {
	red   int
	green int
	blue  int
}

func Solution(file string) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	sum := 0
	sum_2 := 0
	for scanner.Scan() {

		inputString := scanner.Text()

		gameNumber := helper(inputString)

		productGameNumber := day2(inputString)
		sum += gameNumber
		sum_2 += productGameNumber
	}
	fmt.Printf("Day 1: %d\nDay 2: %d\n", sum, sum_2)
}

func day2(str string) int {
	_, plays, err := parseInput(str)
	if err != nil {
		fmt.Println("Error parsing input", err)
		return 0
	}
	if err != nil {
		fmt.Println("Error parsing input", err)
		return 0
	}

	return getProductMinimumCubes(plays)

}

func helper(str string) int {
	cubes := Cubes{
		red:   12,
		green: 13,
		blue:  14,
	}

	gameNumber, plays, err := parseInput(str)
	if err != nil {
		fmt.Println("Error parsing input", err)
		return 0
	}
	if !checkCubeAvailability(cubes, plays) {
		return 0
	}

	return gameNumber
}

func parseGameNumber(input string) int {
	splitStr := strings.Split(input, " ")
	if len(splitStr) < 2 {
		return -1
	}
	number, _ := strconv.Atoi(splitStr[1])
	return number
}

func checkCubeAvailability(cubes Cubes, plays []string) bool {
	for _, play := range plays {
		play = strings.TrimSpace(play)
		hands := strings.Split(play, ",")
		for _, hand := range hands {
			hand = strings.TrimSpace(hand)
			parts := strings.Split(hand, " ")
			count, _ := strconv.Atoi(parts[0])
			color := parts[1]
			switch color {
			case "red":
				if cubes.red < count {
					return false
				}
			case "green":
				if cubes.green < count {
					return false
				}
			case "blue":
				if cubes.blue < count {
					return false
				}
			}
		}
	}
	return true
}

func getProductMinimumCubes(plays []string) int {
	maxRed := math.MinInt
	maxGreen := math.MinInt
	maxBlue := math.MinInt
	for _, play := range plays {
		play = strings.TrimSpace(play)
		hands := strings.Split(play, ",")
		for _, hand := range hands {
			hand = strings.TrimSpace(hand)
			parts := strings.Split(hand, " ")
			count, _ := strconv.Atoi(parts[0])
			color := parts[1]
			switch color {
			case "red":
				if maxRed < count {
					maxRed = count
				}
			case "green":
				if maxGreen < count {
					maxGreen = count
				}
			case "blue":
				if maxBlue < count {
					maxBlue = count
				}
			}

		}
	}
	return maxRed * maxGreen * maxBlue
}

func parseInput(input string) (int, []string, error) {
	gamePlays := strings.Split(input, ":")
	if len(gamePlays) != 2 {
		return 0, nil, fmt.Errorf("input must containt exactly one ':' character")
	}

	gameNumber := parseGameNumber(gamePlays[0])
	plays := strings.Split(gamePlays[1], ";")
	return gameNumber, plays, nil
}
