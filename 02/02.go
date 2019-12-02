package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func atoi(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return
}

func loadMemory(rawData string) (memory []int) {
	s := strings.Split(string(rawData), ",")
	memory = make([]int, len(s))

	for i := 0; i < len(memory); i++ {
		memory[i] = atoi(s[i])
	}

	return
}

func execute(memory []int, noun int, verb int) []int {
	memory = append([]int(nil), memory...)

	memory[1] = noun
	memory[2] = verb

	address := 0
	for {
		opcode := memory[address]

		if opcode == 99 {
			break
		}

		arg1 := memory[memory[address+1]]
		arg2 := memory[memory[address+2]]
		resultAddress := memory[address+3]
		result := 0

		switch opcode {
		case 1:
			result = arg1 + arg2
		case 2:
			result = arg1 * arg2
		default:
			panic(fmt.Sprintf("Unknown opcode %d at address %d", opcode, address))
		}

		memory[resultAddress] = result

		address += 4
	}

	return memory
}

func main() {
	rawData, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	memory := loadMemory(string(rawData))

	part1 := execute(memory, 12, 2)
	fmt.Printf("Part 1: %d\n", part1[0])

	for x := 0; x <= 99; x++ {
		for y := 0; y <= 99; y++ {
			part2 := execute(memory, x, y)

			if part2[0] == 19690720 {
				fmt.Printf("Part 2: noun=%d verb=%d 100 * noun + verb = %d\n", x, y, 100*x+y)
				return
			}
		}
	}
}
