package main

import (
	"fmt"
	"strconv"
	"strings"
)

// function for getting input
func getInput() string {
	return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

}

type Point struct {
	x int
	y int
}

type Line struct {
	p1 Point
	p2 Point
}

func parsePoint(line string) (*Point, error) {
	parts := strings.Split(line, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}
	return &Point{x: x, y: y}, nil
}

func parseLine(line string) (*Line, error) {
	parts := strings.Split(line, " -> ")
	fmt.Println(parts)
	p1, err := parsePoint(parts[0])
	if err != nil {
		return nil, err
	}
	p2, err := parsePoint(parts[1])
	if err != nil {
		return nil, err
	}
	return &Line{p1: *p1, p2: *p2}, nil
}

func main() {
	lines := []Line{}
	input := getInput()
	for _, line := range strings.Split(input, "\n") {
		line, err := parseLine(line)
		if err != nil {
			panic(err)
		}
		lines = append(lines, *line)
	}
	filtered := filterLines(lines)
	// filter lines when they are both horizontal or vertical
	// and save them to new list
	// then print the list

	fmt.Println("%+v", filtered)
}

func filterLines(lines []Line) (filtered []Line) {
	for _, line := range lines {
		if isHorizontal(line) || isVertical(line) {
			filtered = append(filtered, line)
		}
	}
	return
}

func isHorizontal(line Line) bool {
	return line.p1.y == line.p2.y
}

func isVertical(line Line) bool {
	return line.p1.x == line.p2.x
}
