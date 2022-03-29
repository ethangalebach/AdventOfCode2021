package day1

import (
	"os"
	"bufio"
	"strconv"
	"log"
)

func SingleMeasure(s string) int {
	increaseCount := 0

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nextString := scanner.Text()
	next, err := strconv.Atoi(nextString)
	if err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		curr := next
		nextString = scanner.Text()
		next, err = strconv.Atoi(nextString)
		if err != nil {
			log.Fatal(err)
		}
		if (next > curr) {
			increaseCount++
		}
	}

	return increaseCount
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
	 result += v
	}
	return result
 }

func ThreeMeasure(s string) int {
	increaseCount := 0

	queue := make([]int, 0, 4)

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nextString := scanner.Text()
		nextDepth, err := strconv.Atoi(nextString)
		if err != nil {
			log.Fatal(err)
		}
		queue = append(queue, nextDepth)
		if (len(queue) == 4) {
			curr := sum(queue[:3])
			next := sum(queue[1:])
			if (next > curr) {
				increaseCount++
			}
			queue = queue[1:]
		}
	}

	return increaseCount
}

