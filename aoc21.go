package main

import (
	"fmt"
	"aoc21/day1"
)

func main() {
	fmt.Println("Day 1: Sonar Sweep")
	fmt.Println("Number of Depth Increases (Single Measurement): ", day1.SingleMeasure("day1/input.txt"))
	fmt.Println("Number of Depth Increases (Three-Measurement Window): ", day1.ThreeMeasure("day1/input.txt"))
}