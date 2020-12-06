package main

type Group []string

func (g Group) CountYesQuestions(unique bool) int {
	set := make(map[string]int, 0)
	for _, answer := range g {
		for _, question := range answer {
			count := set[string(question)]
			set[string(question)] = count+1
		}
	}

	rt := len(set)

	if unique {
		rt = 0
		for _, v := range set {
			if len(g) == v {
				rt++
			}
		}
	}

	return rt
}

func readBatch(filename string) []Group {
	groups := make([]Group, 0)
	readCsvBatch(filename, "", func(batch []string) error {
		groups = append(groups, batch)
		return nil
	}, logError)
	return groups
}

func day6(puzzle int) int {
	groups := readBatch("data/day6.txt")
	switch puzzle {
	case 1:
		sum := 0
		for _, group := range groups {
			sum += group.CountYesQuestions(false)
		}
		return sum
	case 2:
		sum := 0
		for _, group := range groups {
			sum += group.CountYesQuestions(true)
		}
		return sum
	default:
		panic("unknown puzzle number")
	}
}
