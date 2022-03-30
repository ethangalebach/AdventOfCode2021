package main

import (
	"fmt"
	"aoc21/day1"
	"aoc21/day2"
	"aoc21/day3"
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

	fmt.Println("Day 3: Binary Diagnostic")
	gamma, epsilon, pc := day3.PowerConsumption("day3/input.txt")
	fmt.Println("Gamma Rate, Epsilon Rate, Power Consumption: ", gamma, ", ", epsilon, ", ", pc)
	ogr, csr, lsr := day3.LifeSupportRating("day3/input.txt")
	fmt.Println("Oxygen Generator Rating, CO2 Scrubber Rating, Life Support Rating: ", ogr, ", ", csr, ", ", lsr)
}