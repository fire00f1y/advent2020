package main

import (
	"sync"
)

type Map struct {
	rows   []string
	mu     *sync.Mutex
}

func (m *Map) AddRow(row string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.rows = append(m.rows, row)
}

func (m *Map) TreeSlope(x, y int) int {
	if m.rows == nil || len(m.rows) == 0 {
		panic("bad map input")
	}

	currentX := 0
	currentY := 0
	trees := 0

	for currentY < len(m.rows) {
		row := m.rows[currentY]
		if currentX >= len(row) {
			currentX -= len(row)
		}

		if row[currentX] == '#' {
			trees++
		}

		currentY += y
		currentX += x
	}

	return trees
}

func loadMap() *Map {
	level := &Map{
		rows:   make([]string, 0),
		mu:     &sync.Mutex{},
	}
	readCsv("data/day3.txt", func(line string) error {
		level.AddRow(line)
		return nil
	}, logError)

	return level
}

func day3(puzzle int) int {
	m := loadMap()
	switch puzzle {
	case 1:
		return m.TreeSlope(3, 1)
	case 2:
		return m.TreeSlope(1,1) * m.TreeSlope(3,1) * m.TreeSlope(5,1) * m.TreeSlope(7,1) * m.TreeSlope(1,2)
	default:
		panic("unknown puzzle number")
	}
}
