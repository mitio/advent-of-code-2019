package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func atoi(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return number
}

type point struct {
	x, y int
}

func (a *point) Equals(b *point) bool {
	return a.x == b.x && a.y == b.y
}

func (a *point) distanceFrom(b *point) int {
	return int(math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y)))
}

type line struct {
	start, end point
}

func (line *line) isVertical() bool {
	return line.start.x == line.end.x
}

func (line *line) isHorizontal() bool {
	return line.start.y == line.end.y
}

func (l *line) normalized() *line {
	if l.isHorizontal() {
		if l.start.x < l.end.x {
			return l
		} else {
			return &line{l.end, l.start}
		}
	} else {
		if l.start.y < l.end.y {
			return l
		} else {
			return &line{l.end, l.start}
		}
	}
}

func (a *line) intersectionWith(b *line) (*point, bool) {
	a = a.normalized()
	b = b.normalized()

	if a.isVertical() {
		if b.isHorizontal() {
			if a.start.y <= b.start.y && b.start.y <= a.end.y &&
				b.start.x <= a.start.x && a.start.x <= b.end.x {
				return &point{a.start.x, b.start.y}, true
			}
		} else {
			// TODO: Not handled. May have multiple intersection points.
		}
	} else {
		if b.isVertical() {
			return b.intersectionWith(a)
		}
	}

	return nil, false
}

func (a *line) steps() int {
	return a.start.distanceFrom(&a.end)
}

type intersection struct {
	point point
	steps int
}

func pointsFrom(lineVectors string) []point {
	var x, y int

	var points []point
	points = append(points, point{x, y})

	for _, vector := range strings.Split(lineVectors, ",") {
		offset := atoi(vector[1:])

		switch vector[0] {
		case 'R':
			x += offset
		case 'L':
			x -= offset
		case 'U':
			y += offset
		case 'D':
			y -= offset
		}

		points = append(points, point{x, y})
	}

	return points
}

func debug(format string, args ...interface{}) {
	if os.Getenv("DEBUG") != "1" {
		return
	}

	fmt.Printf(format, args...)
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

	lineVectors := strings.Split(rawData, "\n")

	pointsA := pointsFrom(lineVectors[0])
	pointsB := pointsFrom(lineVectors[1])

	var intersections []intersection
	var stepsA, stepsB int

	for i, endA := range pointsA {
		if i == 0 {
			stepsA = 0
			continue
		}

		a := line{pointsA[i-1], endA}

		debug("A: %v steps %d total steps %d\n", a, a.steps(), stepsA)

		for j, endB := range pointsB {
			if j == 0 {
				stepsB = 0
				continue
			}

			b := line{pointsB[j-1], endB}

			debug("  B: %v steps %d total steps %d\n", b, b.steps(), stepsB)

			intersectionPoint, intersects := a.intersectionWith(&b)
			if intersects && !intersectionPoint.Equals(&point{0, 0}) {
				steps := stepsA + stepsB + a.start.distanceFrom(intersectionPoint) + b.start.distanceFrom(intersectionPoint)
				intersections = append(intersections, intersection{*intersectionPoint, steps})

				debug("    intersection at %v after %d\n", intersectionPoint, steps)
			}

			stepsB += b.steps()
		}

		stepsA += a.steps()
	}

	debug("points of line A: %v\n", pointsA)
	debug("points of line B: %v\n", pointsB)

	fmt.Println(intersections)

	origin := point{0, 0}
	var closestIntersection point
	var minDistance int
	var minSteps int

	for i, intersection := range intersections {
		distance := intersection.point.distanceFrom(&origin)

		if i == 0 || distance < minDistance {
			closestIntersection = intersection.point
			minDistance = distance
		}

		if i == 0 || intersection.steps < minSteps {
			minSteps = intersection.steps
		}
	}

	fmt.Printf("closest point: %v at distance %d. Min steps are %d\n", closestIntersection, minDistance, minSteps)

}
