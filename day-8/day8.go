package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	row, col int
}

type Dist struct {
	dRow, dCol int
}

type Grid struct {
	nCol, nRow int
	nodes map[byte][]Coord
	antiNodes map[Coord]struct{}
}

func main() {
	fileName := "day8.txt"
	fmt.Printf("part1: %d\n", part1(fileName))
	fmt.Printf("part2: %d\n", part2(fileName))
}

func part1(fileName string) int {
	grid := buildGridFromFile(fileName)
	grid.getAntiNodes()
	return len(grid.antiNodes)
}

func part2(fileName string) int {
	grid := buildGridFromFile(fileName)
	grid.getAntiNodesP2()
	return len(grid.antiNodes)
}

func buildGridFromFile(fileName string) Grid {
	b, err := os.ReadFile(fileName); if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")

	nRow := len(lines)
	nCol := len(lines[0])

	nodes := make(map[byte][]Coord)
	antiNodes := make(map[Coord]struct{})
	for iRow, line := range lines {
		bytes := []byte(line)
		for iCol, bt := range bytes {
			if bt == '.' {
				continue
			}
			c := Coord{iRow, iCol}
			_, ok := nodes[bt]; if !ok {
				nodes[bt] = make([]Coord, 0)
			}
			nodes[bt] = append(nodes[bt], c)
		}
	}
	return Grid{
		nRow,
		nCol,
		nodes,
		antiNodes,
	}
}

func (g* Grid) getAntiNodes() {
	for b := range g.nodes {
		g.getAntiNodesForFreq(b)
	}
}

func (g* Grid) getAntiNodesP2() {
	for b := range g.nodes {
		g.getAntiNodesForFreqP2(b)
	}
}
func (g* Grid) getAntiNodesForFreq(b byte) {
	coords, ok := g.nodes[b]; if !ok {
		log.Fatalf("no nodes for frequency %b\n", b)
	}
	l := len(coords)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			c1 := coords[i]
			c2 := coords[j]
			d := c1.dist(c2)
			an1 := c1.addDist(d)
			an2 := c2.subDist(d)

			if g.inbounds(an1) {
				g.antiNodes[an1] = struct{}{}
			}
			
			if g.inbounds(an2) {
				g.antiNodes[an2] = struct{}{}
			}
		}
	}
}

func (g* Grid) getAntiNodesForFreqP2(b byte) {
	coords, ok := g.nodes[b]; if !ok {
		log.Fatalf("no nodes for frequency %b\n", b)
	}

	l := len(coords)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			c1 := coords[i]
			c2 := coords[j]
			d := c1.dist(c2)

			an := c1
			for g.inbounds(an) {
				g.antiNodes[an] = struct{}{}
				an = an.addDist(d)
			}

			an = c2
			for g.inbounds(an) {
				g.antiNodes[an] = struct{}{}
				an = an.subDist(d)
			}
		}
	}

}

func (g* Grid) inbounds(c Coord) bool {
	return c.col >= 0 && c.row >= 0 && c.col < g.nCol && c.row < g.nRow
}

func (c Coord) addDist(d Dist) Coord {
	return Coord{c.row + d.dRow, c.col + d.dCol}
}

func (c Coord) subDist(d Dist) Coord {
	return Coord{c.row - d.dRow, c.col - d.dCol}
}

func (c Coord) dist(c2 Coord) Dist {
	return Dist{c.row - c2.row, c.col - c2.col}
}

func (g* Grid) PrintAntiNodes() {
	fmt.Println("Printing AntiNode Positions:")
	for an := range g.antiNodes {
		fmt.Println(an)
	}
	fmt.Println("")
}