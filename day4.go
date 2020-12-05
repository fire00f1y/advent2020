package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	byrMatch         = regexp.MustCompile(`^(?:19[2-9][0-9]|200[0-2])$`)
	iyrMatch         = regexp.MustCompile(`^(?:201[0-9]|2020)$`)
	eyrMatch         = regexp.MustCompile(`^(202[0-9]|2030)$`)
	colorMatch       = regexp.MustCompile(`#[0-9a-z]{6}`)
	numberMatch      = regexp.MustCompile(`^[0-9]{9}$`)
	dimensionMatch   = regexp.MustCompile(`^(?:1[5-8][0-9]|19[0-3])cm|(?:59|6[0-9]|7[0-6])in$`)
	inColorSet       = regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`)

	requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
)

type Passport map[string]string

func (p Passport) IsValid(matchField bool) bool {
	valid := true

	for _, field := range requiredFields {
		value, ok := p[field]
		if !ok && field != "cid" {
			valid = false
		}
		if matchField && !isFieldValid(field, value) {
			valid = false
		}
	}

	return valid
}

func isFieldValid(field string, value string) bool {
	switch field {
	case "byr":
		{
			return byrMatch.Match([]byte(value))
		}
	case "iyr":
		{
			return iyrMatch.Match([]byte(value))
		}
	case "eyr":
		{
			return eyrMatch.Match([]byte(value))
		}
	case "hgt":
		{
			return dimensionMatch.Match([]byte(value))
		}
	case "hcl":
		{
			return colorMatch.Match([]byte(value))
		}
	case "ecl":
		{
			return inColorSet.Match([]byte(value))
		}
	case "pid":
		{
			return numberMatch.Match([]byte(value))
		}
	case "cid":
		{
			return true
		}
	}
	return false
}

func readPassport(filename string) []Passport {
	r := make([]Passport, 0)

	f, err := os.OpenFile(filename, os.O_RDONLY, 0665)
	if err != nil {
		panic("error opening file " + filename)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	total := 0
	p := Passport{}
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			if len(p) > 0 {
				r = append(r, p)
			}
			break
		}
		line := strings.TrimSpace(string(lineBytes))
		if err != nil {
			log.Printf("failed to read line: %v\n", err)
			continue
		}
		if line == "" {
			r = append(r, p)
			total++
			p = Passport{}
			continue
		}
		attrs := strings.Split(line, " ")
		for _, attr := range attrs {
			t := strings.Split(attr, ":")
			p[t[0]] = t[1]
		}
	}

	return r
}

func day4(puzzle int) int {
	passports := readPassport("data/day4.txt")
	switch puzzle {
	case 1:
		count := 0
		for _, p := range passports {
			if p.IsValid(false) {
				count++
			}
		}
		return count
	case 2:
		count := 0
		for _, p := range passports {
			if p.IsValid(true) {
				count++
			}
		}
		return count
	default:
		panic("unknown puzzle number")
	}
}
