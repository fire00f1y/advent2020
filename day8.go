package main

import (
	"strconv"
	"strings"
)

type BootCode struct {
	Instructions []Instruction
	visited      map[int]struct{}
	iter         int
	acc          int
}

func (bc *BootCode) Reset() {
	bc.acc = 0
	bc.visited = make(map[int]struct{}, 0)
	bc.iter = 0
}

func (bc *BootCode) Step(modIndex int) (int, bool, bool) {
	if _, visited := bc.visited[bc.iter]; visited {
		return bc.acc, true, false
	}
	if bc.iter >= len(bc.Instructions) {
		return bc.acc, false, true
	}
	bc.visited[bc.iter] = struct{}{}

	instr := bc.Instructions[bc.iter]
	if bc.iter == modIndex {
		if instr.Op == "nop" {
			instr.Op = "jmp"
		} else if instr.Op == "jmp" {
			instr.Op = "nop"
		}
	}

	switch instr.Op {
	case "nop":
		{
			bc.iter++
			break
		}
	case "acc":
		{
			bc.acc += bc.Instructions[bc.iter].Arg
			bc.iter++
			break
		}
	case "jmp":
		{
			bc.iter += bc.Instructions[bc.iter].Arg
			break
		}
	default:
		panic("invalid instruction: " + bc.Instructions[bc.iter].Op)
	}

	return bc.acc, false, false
}

func (bc *BootCode) AddInstruction(instruction Instruction) {
	bc.Instructions = append(bc.Instructions, instruction)
}

type Instruction struct {
	Op  string
	Arg int
}

func day8(puzzle int) int {
	bootCode := &BootCode{
		Instructions: make([]Instruction, 0),
		visited:      make(map[int]struct{}, 0),
	}
	readCsv("data/day8.txt", func(line string) error {
		parts := strings.Split(line, " ")
		arg, err := strconv.Atoi(parts[1][1:])
		if err != nil {
			return err
		}
		if parts[1][0] == '-' {
			arg *= -1
		}
		instr := Instruction{
			Op:  parts[0],
			Arg: arg,
		}
		bootCode.AddInstruction(instr)
		return nil
	}, logError)

	switch puzzle {
	case 1:
		for {
			value, repeated, _ := bootCode.Step(-1)
			if repeated {
				return value
			}
		}
		return -1
	case 2:
		for i := 0; i < len(bootCode.Instructions); i++ {
			for {
				value, repeated, finished := bootCode.Step(i)
				if finished {
					return value
				}
				if repeated {
					bootCode.Reset()
					break
				}
			}
		}
		return -1
	default:
		panic("unknown puzzle number")
	}
}
