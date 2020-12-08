package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BootCode struct {
	Instructions []Instruction
	visited      map[int]struct{}
	iter         int
	acc          int
}

func (bc *BootCode) Step() (int, bool) {
	if _, visited := bc.visited[bc.iter]; visited {
		return bc.acc, true
	}
	bc.visited[bc.iter] = struct{}{}

	instr := bc.Instructions[bc.iter]

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

	return bc.acc, false
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
			value, repeated := bootCode.Step()
			if repeated {
				return value
			}
			fmt.Printf("stage %d: %d\n", bootCode.iter, value)
		}
		return -1
	case 2:
		return -1
	default:
		panic("unknown puzzle number")
	}
}
