package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coords struct {
	row, col int
}

func (c Coords) AddRow(dRow int) Coords {
	return c.Add(dRow, 0)
}

func (c Coords) AddCol(dCol int) Coords {
	return c.Add(0, dCol)
}

func (c Coords) Add(dRow, dCol int) Coords {
	return Coords{c.row + dRow, c.col + dCol}
}

type Grid struct {
	nRow, nCol int;
	grid [][]byte;
	visited map[Coords]struct{}
	obst map[Coords](map[byte]struct{})
}

type Cursor struct {
	coords Coords
	direction byte;
}

func main() {
	b, err := os.ReadFile("day6.txt")

	if err != nil {
		log.Fatal(err)
	}

	grid := createGrid(string(b))
	cursor := getCursorFromGrid(&grid)
	startCoords := cursor.coords
	grid.traverse(&cursor)

	part2 := 0
	for visited := range grid.visited {
		if visited == startCoords {
			continue
		}
		// fmt.Printf("\nAdded Obstacle At: %+v\n", visited)
		
		newGrid := createGrid(string(b))
		newGrid.AddObstacle(visited)
		newCursor := Cursor{startCoords, 0}

		hasCycle := newGrid.traverse(&newCursor)
		if hasCycle {
			part2++
		}
	}

	part1 := len(grid.visited)
	// grid.print()
	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
	
}

func createGrid(input string) Grid {
	inputRows := strings.Split(input, "\n");
	grid := Grid{len(inputRows), 
			len(inputRows[0]), 
			make([][]byte, len(inputRows)), 
			make(map[Coords]struct{}),
			make(map[Coords](map[byte]struct{}))}

	for i, row := range inputRows {
		grid.grid[i] = []byte(row)
	}
	return grid
}

func (g* Grid) AddObstacle(c Coords) {
	g.grid[c.row][c.col] = '#';
} 

func getCursorFromGrid(g *Grid) Cursor {
	for row := 0; row < g.nRow; row++ {
		for col := 0; col < g.nCol; col++ {
			if g.grid[row][col] == '^' {
				return Cursor{Coords{row, col}, 0}
			}
		}
	}
	log.Fatalf("could not find cursor in grid")
	return Cursor{Coords{0, 0}, 'x'}
}

func (g* Grid) traverse(c* Cursor) (bool) {
	if (!g.inbounds(c.coords)) {
		return false
	}

	dRow, dCol := 0, 0
	switch c.direction {
	case 0: 
		dRow = -1
	case 1:
		dCol = 1
	case 2:
		dRow = 1
	case 3:
		dCol = -1
	default:
		log.Fatalf("expected an arrow, got %d", c.direction)
	}

	nextCoord := c.coords.Add(dRow, dCol)
	if g.inbounds(nextCoord) && g.at(nextCoord) == '#' {
		_, ok := g.obst[nextCoord]
		if !ok {
			g.obst[nextCoord] = make(map[byte]struct{})
		}
		_, ok = g.obst[nextCoord][c.direction]
		if ok {
			return true
		}
		
		g.obst[nextCoord][c.direction] = struct{}{}
		// fmt.Println(nextCoord, g.obst)
		nc := c.rotate()
		// fmt.Println("Rotating")
		return g.traverse(&nc)
	} else {
		g.visited[c.coords] = struct{}{}
		nc := c.step(dRow, dCol)
		// fmt.Println("Stepping")
		return g.traverse(&nc)
	}
}

func (g* Grid) inbounds(c Coords) bool {
	return c.row >= 0 && c.col >= 0 && c.row < g.nRow && c.col < g.nCol
}

func (g* Grid) at(c Coords) byte {
	return g.grid[c.row][c.col]
}

func (c Cursor) rotate() Cursor {
	return Cursor{c.coords, (c.direction + 1) % 4}
}

func (c Cursor) step(dRow, dCol int) Cursor {
	return Cursor{c.coords.Add(dRow, dCol), c.direction}
}
