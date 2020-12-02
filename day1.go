package main

import (
	"fmt"
	"strconv"
)

const (
	TargetNumber = 2020
)

func day1() {
	output := make([]int, 0)
	readCsv("data/day1.txt", func(b []byte) error {
		i, e := strconv.Atoi(string(b))
		if e != nil {
			return e
		}
		output = append(output, i)

		return nil
	}, logError)
	fmt.Printf("[1:1] Product: %d\n", productOfTwo(output))
	fmt.Printf("[1:2] Product: %d\n", productOfThree(output))
}

func productOfTwo(nums []int) int {
	for _, i := range nums {
		for _, j := range nums {
			if i + j == TargetNumber {
				return i*j
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
					return i*j*k
				}
			}
		}
	}
	return 0
}
