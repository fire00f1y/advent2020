package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	colorMatch  = regexp.MustCompile(`#[0-9a-z]{6}`)
	numberMatch = regexp.MustCompile(`^[0-9]{9}$`)

	requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	allowedColors  = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
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
			v, _ := strconv.Atoi(value)
			return len(value) == 4 && v >= 1920 && v <= 2002
		}
	case "iyr":
		{
			v, _ := strconv.Atoi(value)
			return len(value) == 4 && v >= 2010 && v <= 2020
		}
	case "eyr":
		{
			v, _ := strconv.Atoi(value)
			return len(value) == 4 && v >= 2020 && v <= 2030
		}
	case "hgt":
		{
			cm := strings.HasSuffix(value, "cm")
			inches := strings.HasSuffix(value, "in")
			if !cm && !inches {
				return false
			}

			if cm {
				c, _ := strconv.Atoi(strings.Replace(value, "cm", "", -1))
				return c >= 150 && c <= 193
			}
			if inches {
				c, _ := strconv.Atoi(strings.Replace(value, "in", "", -1))
				return c >= 59 && c <= 76
			}

		}
	case "hcl":
		{
			return colorMatch.Match([]byte(value))
		}
	case "ecl":
		{
			count := 0
			for _, color := range allowedColors {
				if strings.TrimSpace(value) == strings.TrimSpace(color) {
					count++
				}
			}

			return count == 1
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
