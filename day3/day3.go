package day3

import (
	"os"
	"bufio"
	"strings"
	"log"
	"math"
	"strconv"
)

func PowerConsumption(s string) (int, int, int) {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	arr := strings.Split(line, "")
	length := len(arr)
	c, g, e := make([]int, length), make([]int, length), make([]int, length)
	dg, de := 0, 0

	for scanner.Scan() {
		line = scanner.Text()
		arr = strings.Split(line, "")
		for i, s := range arr {
			if s == "1" {
				c[i]++
			} else if (s == "0") {
				c[i]--
			}
		}
	}

	for i, n := range c {
		if n >= 0 {
			g[i] = 1
			e[i] = 0
		} else {
			g[i] = 0
			e[i] = 1
		}
		base10 := int(math.Pow(2, float64(length-1-i)))
		dg += g[i] * base10
		de += e[i] * base10
	}

	return dg, de, dg*de
}


func LifeSupportRating(s string) (int, int, int) {
	m := make(map[string]int)
	length := 0

	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		length = len(line)
		for i := 0; i <= length; i++ {
			m[string(line[:i])] += 1
		}
	}

	ogr := ""
	csr := ""

	for i := 0; i < length; i++ {
		if m[ogr] == 1 {
			if m[ogr + "0"] == 1 {
				ogr += "0"
			}
			if m[ogr + "1"] == 1 {
				ogr += "1"
			}
		} else {
			if m[ogr + "0"] > m[ogr + "1"] {
				ogr += "0"
			} else {
				ogr += "1"
			}
		}

		if m[csr] == 1 {
			if m[csr + "0"] == 1 {
				csr += "0"
			}
			if m[csr + "1"] == 1 {
				csr += "1"
			}
		} else {
			if (m[csr + "0"] <= m[csr + "1"]) {
				csr += "0"
			} else {
				csr += "1"
			}
		}
	}

	ogrd, err := strconv.ParseInt(ogr, 2, 64);
	csrd, err := strconv.ParseInt(csr, 2, 64);
	if err != nil {
			log.Fatal(err)
	}

	return int(ogrd), int(csrd), int(ogrd*csrd)
}