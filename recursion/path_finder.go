package main

import (
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

type Maze struct {
	mapa  []string
	wall  rune
	start rune
	end   rune
	size  Point
	path  []Point
	seen  [][]bool
}

var dirs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	mazeMap := []string{
		"xxxxxxxsxxx",
		"xxxxx    xx",
		"xxxxxxxx xx",
		"xxxxxxxx xx",
		"xx       xx",
		"xxexxxxxxxx",
	}
	maze := Maze{
		mapa:  mazeMap,
		wall:  'x',
		start: 's',
		end:   'e',
	}
	resolve(&maze)
}

func walk(maze *Maze, p Point) bool {
	// BASE CASES
	// If i am outside my map
	if p.x < 0 || p.x > maze.size.x ||
		p.y < 0 || p.y > maze.size.y {
		return false
	}
	// If i hit a wall
	if maze.mapa[p.y][p.x] == 'x' {
		return false
	}
	// If i am on the end i return true
	if maze.mapa[p.y][p.x] == 'e' {
		maze.path = append(maze.path, p)
		return true
	}
	// if i have already been in this location
	if maze.seen[p.y][p.x] {
		return false
	}
	maze.seen[p.y][p.x] = true
	maze.path = append(maze.path, p)
	for _, dir := range dirs {
		n := Point{y: p.y + dir[0], x: p.x + dir[1]}
		// fmt.Println(dir[0], ",", dir[1])
		if walk(maze, n) {
			return true
		}
	}
	// maze.path = maze.path[:len(maze.path)-1]
	return false
}

func getStart(mapa []string) *Point {
	for y, line := range mapa {
		for x, c := range line {
			if c == 's' {
				return &Point{x: x, y: y}
			}
		}
	}
	return nil
}

func resolve(maze *Maze) {
	maze.size.x = len(maze.mapa[0])
	maze.size.y = len(maze.mapa)
	start := getStart(maze.mapa)
	if start == nil {
		fmt.Println("Unable to find the start in the maze")
		os.Exit(1)
	}
	seen := make([][]bool, maze.size.y)
	for i := range seen {
		seen[i] = make([]bool, maze.size.x)
	}
	maze.seen = seen
	walk(maze, *start)
	for _, i := range maze.path {
		fmt.Printf("x: %d, y: %d\n", i.x, i.y)
	}
}
