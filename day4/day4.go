package day4

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func getScore(markers [][]int, card [][]string, lastball string) int {
	unmarkedSum := 0
	rows := len(markers)
	cols := len(markers[0])

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if markers[i][j] == 0 {
				unmarked, err := strconv.Atoi(card[i][j])
				if err != nil {
					log.Println(err)
				}
				unmarkedSum += unmarked
			}
		}
	}

	last, err := strconv.Atoi(lastball)
	if err != nil {
		log.Println(err)
	}

	return unmarkedSum * last
}

func hasBingo(markers [][]int) bool {
	rows := len(markers)
	cols := len(markers[0])
	bingo := false

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if markers[i][j] == 0 { break }
			if j == cols - 1 { bingo = true }
		}
	}

	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if markers[j][i] == 0 { break }
			if j == rows - 1 { bingo = true }
		}
	}

	return bingo
}

func mark(card [][]string, balls []string) (int, int) {
	rows := len(card)
	cols := len(card[0])

	markers := make([][]int, rows)
	for i := 0; i < rows; i++ {
		markers[i] = make([]int, cols)
	}

	for i, ball := range balls {
		for j := 0; j < rows; j++ {
			for k, num := range card[j] {
				if ball == num {
					markers[j][k] = 1
				}
				if hasBingo(markers) {
					return i, getScore(markers, card, ball)
				}
			}
		}
	}
	return 0, 0
}

func boardScore(s string, findFirst bool, findLast bool) int {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	balls := strings.Split(line, ",")
	scanner.Scan()
	card := make([][]string, 5)
	earliestBingoBall := len(balls)
	latestBingoBall := 0
	finalScore := 0
	rownum := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			bingoBall, boardScore := mark(card, balls)
			if bingoBall > latestBingoBall && findLast {
				finalScore = boardScore
				latestBingoBall = bingoBall
			}
			if bingoBall < earliestBingoBall && findFirst {
				finalScore = boardScore
				earliestBingoBall = bingoBall
			}
			rownum = 0
		} else {
			card[rownum] = strings.Fields(line)
			rownum++
		}
	}

	return finalScore
}

func FirstBoardScore(s string) int {
	return boardScore(s, true, false)
}

func LastBoardScore(s string) int {
	return boardScore(s, false, true)
}