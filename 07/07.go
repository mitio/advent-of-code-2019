package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) (result int) {
	result, _ = strconv.Atoi(s)
	return
}

func debug(format string, args ...interface{}) {
	if os.Getenv("DEBUG") != "1" {
		return
	}

	fmt.Printf(format, args...)
}

func loadMemory(rawData string) (memory []int) {
	s := strings.Split(string(rawData), ",")
	memory = make([]int, len(s))

	for i := 0; i < len(memory); i++ {
		memory[i] = atoi(s[i])
	}

	return
}

func paramModeFor(instruction, paramIndex int) int {
	paramModes := instruction / 100

	return paramModes / int(math.Pow10(paramIndex)) % 10
}

func paramValueFor(memory []int, instruction, address, paramIndex int) int {
	paramMode := paramModeFor(instruction, paramIndex)

	switch paramMode {
	case 0:
		return memory[memory[address]]
	case 1:
		return memory[address]
	default:
		log.Fatalf("Unknown param mode: %d (instruction %d at address %d)", paramMode, instruction, address)
	}

	return 0
}

func execute(memory []int, inputs []int) ([]int, []int) {
	memory = append([]int(nil), memory...)
	var outputs []int
	inputsOffset := 0

	address := 0
	for {
		currentAddress := address
		instruction := memory[address]
		opcode := instruction % 100
		offset := 1
		result := 0
		var arg1, arg2, resultAddress int

		if opcode == 99 {
			break
		}

		switch opcode {
		case 1, 2, 5, 6, 7, 8:
			arg1 = paramValueFor(memory, instruction, address+offset, 0)
			offset++
			arg2 = paramValueFor(memory, instruction, address+offset, 1)
			offset++
		case 4:
			arg1 = paramValueFor(memory, instruction, address+offset, 0)
			offset++
		}

		switch opcode {
		case 1, 2, 3, 7, 8:
			resultAddress = memory[address+offset]
			offset++
		}

		switch opcode {
		case 1:
			result = arg1 + arg2
		case 2:
			result = arg1 * arg2
		case 3:
			result = inputs[inputsOffset]
			inputsOffset++
		case 4:
			outputs = append(outputs, arg1)
		case 5:
			if arg1 != 0 {
				offset = 0
				address = arg2
			}
		case 6:
			if arg1 == 0 {
				offset = 0
				address = arg2
			}
		case 7:
			if arg1 < arg2 {
				result = 1
			} else {
				result = 0
			}
		case 8:
			if arg1 == arg2 {
				result = 1
			} else {
				result = 0
			}
		default:
			log.Fatalf("Unknown opcode %d at address %d (instruction: %d)", opcode, address, instruction)
		}

		switch opcode {
		case 1, 2, 3, 7, 8:
			memory[resultAddress] = result
		}

		address += offset

		debug(
			"% 5d: instr: % 5d opcode: % 3d arg1: % 9d arg2: % 9d resultAt: % 5d result: % 9d offs: % 2d\n",
			currentAddress,
			instruction,
			opcode,
			arg1,
			arg2,
			resultAddress,
			result,
			offset,
		)
	}

	return memory, outputs
}

// Stolen from https://yourbasic.org/golang/generate-permutation-slice-string/
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func main() {
	var rawData string

	if len(os.Args) > 1 {
		rawData = strings.Join(os.Args[1:], "\n")
	} else {
		bytes, err := ioutil.ReadFile("input")
		if err != nil {
			panic(err)
		}
		rawData = string(bytes)
	}

	memory := loadMemory(string(rawData))

	maxSignal := 0
	var maxSignalPhases []int

	Perm([]int{0, 1, 2, 3, 4}, func(phases []int) {
		lastOutput := 0

		debug("------------------------------------------\n")
		debug("%v\n", phases)
		debug("------------------------------------------\n")

		for i := 0; i < 5; i++ {
			_, outputs := execute(memory, []int{phases[i], lastOutput})
			lastOutput = outputs[0]
		}

		if lastOutput > maxSignal {
			maxSignal = lastOutput
			maxSignalPhases = append([]int(nil), phases...)
		}
	})

	fmt.Printf("Part 1: %v (phases %v)", maxSignal, maxSignalPhases)
}
