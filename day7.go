package main

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	bagExtractor = regexp.MustCompile(`(\d+ |^)(\w+ \w+) bag`)
)

type BagCount struct {
	count int
	bag   string
}

type Graph struct {
	inverted  map[string][]string
	hierarchy map[string][]BagCount
}

func (g Graph) AddRelation(outer, inner string, count int) {
	if count > 0 {
		l := g.inverted[inner]
		if l == nil {
			l = make([]string, 0)
		}
		g.inverted[inner] = append(l, outer)
	}

	ll := g.hierarchy[outer]
	if ll == nil {
		ll = make([]BagCount, 0)
	}
	if count > 0 {
		g.hierarchy[outer] = append(ll, BagCount{count, inner})
	} else {
		g.hierarchy[outer] = nil
	}
}

func (g Graph) CountLeafs(key string, seen map[string]struct{}) int {
	l := g.inverted[key]
	sum := 0
	for _, v := range l {
		_, skip := seen[v]
		if skip {
			continue
		}
		sum += 1
		sum += g.CountLeafs(v, seen)
		seen[v] = struct{}{}
	}
	return sum
}

func (g Graph) CountBags(target string) int {
	sum := 0

	l := g.hierarchy[target]
	if l == nil || len(l) == 0 {
		return 0
	}

	for _, bag := range l {
		sum += bag.count * g.CountBags(bag.bag) + bag.count
	}

	return sum
}

func ParseBag(bagArray []string) (int, string) {
	count := -1
	bag := bagArray[2]

	if bagArray[1] != "" {
		count, _ = strconv.Atoi(strings.TrimSpace(bagArray[1]))
	}

	return count, bag
}

func day7(puzzle int) int {
	graph := Graph{inverted: make(map[string][]string, 0), hierarchy: make(map[string][]BagCount, 0)}
	readCsv("data/day7.txt", func(line string) error {
		matches := bagExtractor.FindAllStringSubmatch(line, -1)
		var root string
		for _, bagString := range matches {
			count, bag := ParseBag(bagString)
			if count <= 0 {
				if len(matches) == 1 {
					graph.AddRelation(bag, "", 0)
					continue
				}
				root = bag
				continue
			}
			graph.AddRelation(root, bag, count)
		}

		return nil
	}, logError)

	switch puzzle {
	case 1:
		return graph.CountLeafs("shiny gold", make(map[string]struct{}, 0))
	case 2:
		return graph.CountBags("shiny gold")
	default:
		panic("unknown puzzle number")
	}
}
