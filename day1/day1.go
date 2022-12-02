package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var totalScore []int

func main() {
	part1()
}

func part1() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	scanner := bufio.NewScanner(file)
	var score int
	for scanner.Scan() {
		if scanner.Text() == "" {
			totalScore = append(totalScore, score)
			score = 0
		} else {
			result, _ := strconv.Atoi(scanner.Text())
			score += result
		}
	}
	getHighestNumber()
	getTheThirdHighestNumber()
}

func getHighestNumber() {
	max := totalScore[0]
	for _, value := range totalScore {
		if value > max {
			max = value
		}
	}
	fmt.Println("Part 1", max)
}

func getTheThirdHighestNumber() {
	var max1, max2, max3 int
	max1, max2, max3 = totalScore[0], totalScore[0], totalScore[0]
	for _, value := range totalScore {
		if value > max1 {
			max3 = max2
			max2 = max1
			max1 = value
		} else if value > max2 {
			max3 = max2
			max2 = value
		} else if value > max3 {
			max3 = value
		}
	}
	fmt.Println("Part 2", max1+max2+max3)
}
