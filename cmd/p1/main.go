package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getInput() string {
	return `forward 5
down 5
forward 8
up 3
down 8
forward 2`
}

type Direction struct {
	x int
	y int
}

func parseLine(line string) Direction {
	parts := strings.Split(line, " ")

	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("this should not happen")
		panic(err)
	}
	if parts[0] == "forward" {
		return Direction{
			x: amount,
			y: 0,
		}
	}

	if parts[0] == "up" {
		return Direction{
			x: 0,
			y: -amount,
		}
	}
	return Direction{
		x: 0,
		y: amount,
	}
}

func main() {
	lines := strings.Split(getInput(), "\n")
	pos := Direction{
		x: 0,
		y: 0,
	}
	for _, line := range lines {
		position := parseLine(line)
		pos.x += position.x
		pos.y += position.y

	}
	fmt.Printf("point: %+v", pos)
}
