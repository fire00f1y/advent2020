package main

import (
	"fmt"
	"os"
)

var (
	knownAnswers = map[string]int{
		"1:1": 970816,
		"1:2": 96047280,
		"2:1": 572,
		"2:2": 306,
	}
)

func main() {
	days := 3
	successes := playAll(days)
	if successes != days*2 {
		fmt.Printf("%d scenarios tested, %d of them succeeded\n", days*2, successes)
		os.Exit(0)
	}

	fmt.Println("all days were successful")
}

func playAll(days int) int {
	count := 0
	for day := 1; day <= days; day++ {
		for puzzle := 1; puzzle <=2; puzzle++ {
			label := ""
			a := 0

			switch day {
			case 1:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day1(puzzle)
			case 2:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day2(puzzle)
			case 3:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day3(puzzle)
			case 4:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day4(puzzle)
			case 5:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day5(puzzle)
			case 6:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day6(puzzle)
			case 7:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day7(puzzle)
			case 8:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day8(puzzle)
			case 9:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day9(puzzle)
			case 10:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day10(puzzle)
			case 11:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day11(puzzle)
			case 12:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day12(puzzle)
			case 13:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day13(puzzle)
			case 14:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day14(puzzle)
			case 15:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day15(puzzle)
			case 16:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day16(puzzle)
			case 17:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day17(puzzle)
			case 18:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day18(puzzle)
			case 19:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day19(puzzle)
			case 20:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day20(puzzle)
			case 21:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day21(puzzle)
			case 22:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day22(puzzle)
			case 23:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day23(puzzle)
			case 24:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day24(puzzle)
			case 25:
				label = fmt.Sprintf("%d:%d", day, puzzle)
				a = day25(puzzle)
			default:
				fmt.Printf("[%d:%d] has not been implemented yet!\n", day, puzzle)
				continue
			}

			result := knownAnswers[label] == a
			if result {
				count++
			}
			fmt.Printf("[%s] Answer: %d matches known answer: %t\n", label, a, result)
		}
	}

	return count
}
