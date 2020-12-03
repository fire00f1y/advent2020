package main

import (
	"strconv"
)

const (
	TargetNumber = 2020
)

func day1(puzzle int) int {
	output := make([]int, 0)
	readCsv("data/day1.txt", func(b string) error {
		i, e := strconv.Atoi(b)
		if e != nil {
			return e
		}
		output = append(output, i)

		return nil
	}, logError)
	switch puzzle {
	case 1:
		return productOfTwo(output)
	case 2:
		return productOfThree(output)
	default:
		panic("invalid module")
	}
}

func productOfTwo(nums []int) int {
	for _, i := range nums {
		for _, j := range nums {
			if i+j == TargetNumber {
				return i * j
			}
		}
	}
	return 0
}

func productOfThree(nums []int) int {
	for _, i := range nums {
		for _, j := range nums {
			for _, k := range nums {
				if i+j+k == TargetNumber {
					return i * j * k
				}
			}
		}
	}
	return 0
}
