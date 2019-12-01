package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func necessaryFuelFor(mass int) int {
	fuel := mass/3 - 2

	if fuel <= 0 {
		return 0
	}

	return fuel + necessaryFuelFor(fuel)
}

func naiveFuelFor(mass int) int {
	return mass/3 - 2
}

func partOneWithStrings(fuelCalculator func(int) int) int {
	input, _ := ioutil.ReadFile("input")
	total := 0
	for _, word := range strings.Split(string(input), "\n") {
		mass, _ := strconv.Atoi(word)
		total += fuelCalculator(mass)
	}
	return total
}

func partOneWithScanner(fuelCalculator func(int) int) int {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	total := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		total += fuelCalculator(mass)
	}

	return total
}

func main() {
	fmt.Println("part one (strings):", partOneWithStrings(naiveFuelFor))
	fmt.Println("part one (scanner):", partOneWithScanner(naiveFuelFor))
	fmt.Println("part two (strings):", partOneWithStrings(necessaryFuelFor))
	fmt.Println("part two (scanner):", partOneWithScanner(necessaryFuelFor))
}
