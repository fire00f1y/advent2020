package main

import (
	"sort"
	"strings"
)

func calculateSeat(code string, bit int) int {
	seat := 0
	step := len(code) - 1
	for i := 1; i < bit; i = i * 2 {
		if code[step] == 'B' || code[step] == 'R' {
			seat += i
		}
		step--
	}

	return seat
}

func readTickets(filename string) map[int]string {
	tickets := make(map[int]string, 0)
	readCsv(filename, func(line string) error {
		rowBits := strings.TrimSpace(line)[:7]
		colBits := strings.TrimSpace(line)[7:]
		row := calculateSeat(rowBits, 128)
		col := calculateSeat(colBits, 8)
		id := (row * 8) + col
		tickets[id] = strings.TrimSpace(line)

		return nil
	}, logError)

	return tickets
}

func day5(puzzle int) int {
	tickets := readTickets("data/day5.txt")
	switch puzzle {
	case 1:
		max := 0
		for k := range tickets {
			if k > max {
				max = k
			}
		}
		return max
	case 2:
		ids := make([]int, 0)
		for k := range tickets {
			ids = append(ids, k)
		}
		sort.Ints(ids)

		for i, id := range ids {
			if i == 0 || i == len(ids) - 1 {
				continue
			}
			if id - ids[i-1] != 1 {
				return id-1
			}
		}

		return 0
	default:
		panic("unknown puzzle number")
	}
	return 0
}
