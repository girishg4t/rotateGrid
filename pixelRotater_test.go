package main

import (
	"fmt"
	"testing"
)

func TestPixelRotater(t *testing.T) {
	g := grid{
		{0, 16, 255},
		{8, 128, 32},
		{0, 0, 0},
	}
	r := grid{
		{0, 0, 0},
		{32, 128, 8},
		{255, 16, 0},
	}
	result := g.rotate(2)

	if !Equal(r, result) {
		t.Errorf("Not equal %v and %v", result, r)
	}
}

func Equal(a grid, b grid) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			fmt.Println(a[i][j], b[i][j])
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
