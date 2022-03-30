package day2

import (
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

func Dive(s string) (int, int) {
	x, y := 0, 0

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		dir := arr[0]
		dist, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}

		if dir == "forward" {
			x += dist
		} else if dir == "down" {
			y += dist
		} else if dir == "up" {
			y -= dist
		} else {
			log.Fatal("missed row", dir)
		}
	}

	return x, y
}

func AimDive(s string) (int, int) {
	aim, x, y := 0, 0, 0

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, " ")
		dir := arr[0]
		dist, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal(err)
		}

		if dir == "forward" {
			x += dist
			y += aim * dist
		} else if dir == "down" {
			aim += dist
		} else if dir == "up" {
			aim -= dist
		} else {
			log.Fatal("missed row", dir)
		}
	}

	return x, y
}