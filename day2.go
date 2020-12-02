package main

import (
	"strconv"
	"strings"
)

type Password struct {
	lowerBound int
	upperBound int
	character  string
	pass       string
}

func (p Password) IsValid() bool {
	count := 0
	for _, c := range p.pass {
		if p.character == string(c) {
			count++
		}
	}
	return count >= p.lowerBound && count <= p.upperBound
}

func (p Password) IsNewValid() bool {
	first := string(p.pass[p.lowerBound]) == p.character
	second := string(p.pass[p.upperBound]) == p.character
	return !(first && second) && (first || second)
}

func day2(puzzle int) int {
	passwords := make([]Password, 0)
	readCsv("data/day2.txt", func(line string) error {
		a := strings.Split(line, ":")
		rule := a[0]
		actual := a[1]

		a = strings.Split(rule, " ")
		validRange := a[0]
		ch := a[1]

		a = strings.Split(validRange, "-")
		lowerChars := a[0]
		upperChars := a[1]

		lower, e := strconv.Atoi(lowerChars)
		if e != nil {
			return e
		}
		upper, e := strconv.Atoi(upperChars)
		if e != nil {
			return e
		}
		passwords = append(passwords, Password{
			lowerBound: lower,
			upperBound: upper,
			character:  ch,
			pass: actual,
		})

		return nil
	}, logError)

	switch puzzle {
	case 1:
		count := 0
		for _, p := range passwords {
			if p.IsValid() {
				count++
			}
		}
		return count
	case 2:
		count := 0
		for _, p := range passwords {
			if p.IsNewValid() {
				count++
			}
		}
		return count
	default:
		panic("unknown puzzle number")
	}
}
