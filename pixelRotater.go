package main

type grid [][]int

func main() {
	a := grid{
		{0, 16, 255},
		{8, 128, 32},
		{0, 0, 0},
	}
	a.rotate(2)
}

func (g *grid) rotate(k int) grid {
	rotation := 0

	for rotation != k {
		result := rotateOnesClockwise(*g)
		*g = result
		rotation++
	}
	return *g
}

func rotateOnesClockwise(g grid) grid {
	gridAfterRotation := make(grid, 0)

	n := len(g)
	for c := 0; c < n; c++ {
		rowArr := make([]int, 0)
		for r := n - 1; r > -1; r-- {
			rowArr = append(rowArr, g[r][c])
		}
		gridAfterRotation = append(gridAfterRotation, rowArr)
	}

	return gridAfterRotation
}
