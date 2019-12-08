package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func numberOfOrbitsFrom(orbiter string, orbits map[string]string) int {
	center, ok := orbits[orbiter]

	if !ok {
		return 0
	}
	if center == "COM" {
		return 1
	}

	return 1 + numberOfOrbitsFrom(center, orbits)
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

	orbits := make(map[string]string)

	for _, orbit := range strings.Split(rawData, "\n") {
		objects := strings.Split(orbit, ")")
		center := objects[0]
		orbiter := objects[1]

		orbits[orbiter] = center
	}

	totalNumberOfOrbits := 0

	for orbiter, _ := range orbits {
		totalNumberOfOrbits += numberOfOrbitsFrom(orbiter, orbits)
	}

	fmt.Println(totalNumberOfOrbits)

	myParent, _ := orbits["YOU"]
	var myTransfers, santaTransfers int
	var ok bool

	for {
		santaParent, _ := orbits["SAN"]
		santaTransfers = 0

		for {
			if myParent == santaParent || santaParent == "COM" {
				break
			}

			santaParent, ok = orbits[santaParent]
			if !ok {
				break
			}

			santaTransfers++
		}

		if myParent == santaParent || myParent == "COM" {
			break
		}

		myParent, ok = orbits[myParent]
		if !ok {
			break
		}
		myTransfers++
	}

	fmt.Printf("Transfers: %d (my side: %d Santa's side %d)\n", myTransfers+santaTransfers, myTransfers, santaTransfers)
}
