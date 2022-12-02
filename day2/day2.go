package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	score      int
	tournament []turn
)

type turn struct {
	player1 string
	player2 string
}

func main() {
	fmt.Println("Day 2 : Rock, Paper, Scisors")
	getTxt()
	for _, round := range tournament {
		//part1(round)
		part2(round)
	}
	fmt.Println(score)
}

func part1(round turn) {
	battlePart1(round)
	getPoints(round.player2)

}
func part2(round turn) {
	choice := battlePart2(round)
	getPoints(choice)
}

func getTxt() {
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
	for scanner.Scan() {
		player1 := strings.Split(scanner.Text(), " ")[0]
		player2 := strings.Split(scanner.Text(), " ")[1]
		tournament = append(tournament, turn{player1, player2})
	}
}

func getPoints(me string) {
	switch me {
	case "X":
		score += 1
		break
	case "Y":
		score += 2
		break
	case "Z":
		score += 3
		break
	}
}

func battlePart1(round turn) {
	switch round.player2 {
	case "X":
		if round.player1 == "A" {
			score += 3
		} else if round.player1 == "C" {
			score += 6
		}
		break
	case "Y":
		if round.player1 == "B" {
			score += 3
		} else if round.player1 == "A" {
			score += 6
		}
		break
	case "Z":
		if round.player1 == "C" {
			score += 3
		} else if round.player1 == "B" {
			score += 6
		}
		break
	}
}

func battlePart2(round turn) string {
	var choice string
	switch round.player2 {
	case "X":
		if round.player1 == "A" {
			choice = "Z"
		} else if round.player1 == "B" {
			choice = "X"
		} else {
			choice = "Y"
		}
		break
	case "Y":
		if round.player1 == "A" {
			choice = "X"
		} else if round.player1 == "B" {
			choice = "Y"
		} else {
			choice = "Z"
		}
		score += 3
		break
	case "Z":
		if round.player1 == "A" {
			choice = "Y"
		} else if round.player1 == "B" {
			choice = "Z"
		} else {
			choice = "X"
		}
		score += 6
		break

	}
	return choice
}
