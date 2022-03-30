package main

import (
	"fmt"
	"aoc21/day1"
	"aoc21/day2"
)

func main() {
	fmt.Println("Day 1: Sonar Sweep")
	fmt.Println("Number of Depth Increases (Single Measurement): ", day1.SingleMeasure("day1/input.txt"))
	fmt.Println("Number of Depth Increases (Three-Measurement Window): ", day1.ThreeMeasure("day1/input.txt"))

	fmt.Println("Day 2: Dive!")
	x, y := day2.Dive("day2/input.txt")
	fmt.Println("Dive without Aim - Final Position (x,y): ", x, y)
	x, y = day2.AimDive("day2/input.txt")
	fmt.Println("Dive with Aim - Final Position (x,y): ", x, y)
	//fmt.Println("Number of Depth Increases (Three-Measurement Window): ", day1.ThreeMeasure("day1/input.txt"))
}