package main

import (
	"fmt"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		a, b         line
		intersects   bool
		intersection *point
	}{
		{
			line{point{0, 0}, point{0, 10}},
			line{point{0, 0}, point{10, 0}},
			true,
			&point{0, 0},
		},
		{
			line{point{0, 0}, point{0, 10}},
			line{point{1, 1}, point{10, 1}},
			false,
			nil,
		},
		{
			line{point{3, 1}, point{3, 3}},
			line{point{2, 2}, point{5, 2}},
			true,
			&point{3, 2},
		},
		{
			line{point{3, 3}, point{3, 1}},
			line{point{2, 2}, point{5, 2}},
			true,
			&point{3, 2},
		},
		{
			line{point{3, 3}, point{3, 1}},
			line{point{5, 2}, point{2, 2}},
			true,
			&point{3, 2},
		},
	}

	for _, testCase := range tests {
		testName := fmt.Sprintf("intersection(%v, %v)", testCase.a, testCase.b)
		t.Run(testName, func(t *testing.T) {
			intersection, intersects := testCase.a.intersectionWith(&testCase.b)
			if intersects != testCase.intersects ||
				(testCase.intersects && !intersection.Equals(testCase.intersection)) {
				t.Errorf("got (%v, %v) want (%v, %v)", intersection, intersects, testCase.intersection, testCase.intersects)
			}
		})
	}
}
