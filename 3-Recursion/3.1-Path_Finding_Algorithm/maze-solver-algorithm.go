package main

import (
	"fmt"
)

// Point represents a point with 'x' and 'y' coordinates
type Point struct {
	x int
	y int
}

// Directions that we need to go in the array
var directions = [4][2]int{
	{-1, 0}, //Up
	{1, 0},  //Down
	{0, -1}, //Left
	{0, 1},  //Right
}

// Implementing the walk function, that calls itself
func walk(maze []string, wall string, currentPoint Point, end Point, seen [][]bool, path []Point) bool {
	//Base cases
	// 1. Off the map
	if currentPoint.x < 0 || currentPoint.x >= len(maze[0]) ||
		currentPoint.y < 0 || currentPoint.y >= len(maze) {
		return false
	}

	//2. On a wall
	if maze[currentPoint.y][currentPoint.x] == wall[0] {
		return false
	}

	//3. It's the end
	if currentPoint == end {
		path = append(path, end)
		return true
	}

	//4. If we have seen it
	if seen[currentPoint.y][currentPoint.x] {
		return false
	}

	//Applying the 3 steps for recursion
	// Step 1: Pre
	seen[currentPoint.y][currentPoint.x] = true
	path = append(path, currentPoint) //path.push() in JS

	//Step 2: Recurse
	for _, direction := range directions {
		x := currentPoint.x + direction[0]
		y := currentPoint.y + direction[1]
		if walk(maze, wall, Point{x, y}, end, seen, path) {
			return true
		}
	}

	//Step 3: Post
	path = path[:len(path)-1] //path.pop() in JS

	// If we do all of that and don't find the end,
	// return false
	return false
}

// solve finds a path from the start point to the end point in the maze
func solve(maze []string, wall string, start, end Point) []Point {
	seen := make([][]bool, len(maze))
	for i := range seen {
		seen[i] = make([]bool, len(maze[0]))
	}

	var path []Point

	walk(maze, wall, start, end, seen, path)

	return path
}

func main() {
	maze := []string{
		"##########",
		"#        #",
		"#  ## ## #",
		"#     #  #",
		"#  ##    #",
		"#     #  #",
		"#  ## ## #",
		"#        #",
		"##########",
	}

	wall := "#"
	start := Point{1, 1}
	end := Point{8, 8}

	path := solve(maze, wall, start, end)

	fmt.Println("Maze:")
	for _, row := range maze {
		fmt.Println(row)
	}

	fmt.Println("Path from start to end:")
	for _, point := range path {
		// Mark the path in the maze by replacing the corresponding characters with dots
		row := []byte(maze[point.y])
		row[point.x] = '.'
		maze[point.y] = string(row)
	}

	fmt.Println("Modified Maze:")
	for _, row := range maze {
		fmt.Println(row)
	}
}
