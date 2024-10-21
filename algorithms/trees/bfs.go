package search

/**
* Breadth First Search
**/

// Queue Implementation
type Queue struct {
}

// Simple Grid to simulate BFS
type Block int
type Grid [][]Block

func NewGrid() Grid {
	grid := make(Grid, 3)

	for y := range grid {
		gridRow := make([]int, 3)
		for x := range gridRow {
			grid[y][x] = Block(1)
		}
	}

	return grid
}
