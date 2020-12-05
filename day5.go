package main

func findId(pass string, min, max int) int {
	value := 0

	for _, c := range pass {
		if c == 'F' {
			max = max / 2
		} else {
			min = max / 2
			if max % 2 != 0 {
				min++
			}
		}
	}

	if min != max {
		panic("invalid splits occurred")
	}




	return value
}

func day5(puzzle int) int {
	switch puzzle {
	case 1:
		return -1
	case 2:
		return -1
	default:
		panic("unknown puzzle number")
	}
}
