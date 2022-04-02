package day5

import (
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)

type Coord struct {
	x int
	y int
}

func getPoints(line string, c chan Coord, inclDiag bool) {
	fields := strings.Split(line, " ")
	start := strings.Split(fields[0], ",")
	end := strings.Split(fields[len(fields)-1], ",")
	x1, err := strconv.Atoi(start[0])
	y1, err := strconv.Atoi(start[1])
	x2, err := strconv.Atoi(end[0])
	y2, err := strconv.Atoi(end[1])
	if err != nil {
		log.Fatal(err)
	}

	if x1 == x2 {
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		for i := y1; i <= y2; i++ {
			c <- Coord{x1, i}
		}
	} else if y1 == y2 {
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		for i := x1; i <= x2; i++ {
			c <- Coord{i, y1}
		}
	} else if inclDiag && math.Abs(float64(y2-y1)) == math.Abs(float64(x2-x1)) {
		if x2 >= x1 && y2 >= y1 {
			for i := x1; i <= x2; i++ {
				c <- Coord{i, y1+i-x1}
			}
		} else if x2 >= x1 && y1 > y2 {
			for i := x1; i <= x2; i++ {
				c <- Coord{i, y1-i+x1}
			}
		} else if x1 > x2 && y2 >= y1 {
			for i := x2; i <= x1; i++ {
				c <- Coord{i, y2-i+x2}
			}
		} else if x1 > x2 && y1 > y2 {
			for i := x2; i <= x1; i++ {
				c <- Coord{i, y2+i-x2}
			}
		}
	}
}

func CountIntersectionPoints(s string, inclDiag bool) int {
	num_ip := 0
	diagram := make([][]int, 2000)
	for i := 0; i < 2000; i++ {
		diagram[i] = make([]int, 2000)
	}
	c := make(chan Coord, 1000)

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	go func() {
		for scanner.Scan() {
			line := scanner.Text()
			getPoints(line, c, inclDiag)
		}
		close(c)
	}()

	for pt := range c {
		diagram[pt.x][pt.y]++
		if diagram[pt.x][pt.y] == 2 {
			num_ip++
		}
	}

	return num_ip
}