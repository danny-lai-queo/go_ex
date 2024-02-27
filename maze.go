// Write your answer here, and then test your code.
// Your job is to implement the countPaths() function.

package main

import "fmt"

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false
// const showHints = false

type Maze struct {
	values        [][]int
	computedPaths map[string]int
	rows, cols    int
}

func NewMaze(vals [][]int) *Maze {
	m := &Maze{
		values:        vals,
		computedPaths: make(map[string]int),
	}
	m.rows = len(vals)
	m.cols = len(vals[0])
	return m
}

func (m *Maze) GetCost(row, col int) int {
	return m.values[row][col]
}

func (m *Maze) IsDestination(row, col int) bool {
	return row == (m.rows-1) && col == (m.cols-1)
}

func (m *Maze) IsLastRow(row int) bool {
	return row == (m.rows - 1)
}

func (m *Maze) IsLastCol(col int) bool {
	return col == (m.cols - 1)
}

func (m *Maze) getKey(row, col, cost int) string {
	return fmt.Sprintf("%d-%d-%d", row, col, cost)
}

func (m *Maze) SetPathCount(row, col, cost, value int) {
	key := m.getKey(row, col, cost)
	m.computedPaths[key] = value
}

func (m *Maze) GetPathCount(row, col, cost int) (int, bool) {
	key := m.getKey(row, col, cost)
	val, ok := m.computedPaths[key]
	return val, ok
}

func countPaths(maze *Maze, row int, col int, cost int) int {
	//fmt.Printf("maze rows %d cols %d\n", maze.rows, maze.cols)
	maze.SetPathCount(0, 0, maze.GetCost(0, 0), 1)
	var queue [][]int = [][]int{{0, 0, maze.GetCost(0, 0)}}
	totalPaths := 0
	for len(queue) > 0 {
		//fmt.Printf("%d %v\n", totalPaths, queue)
		i, j, tc := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		if tc == cost && maze.IsDestination(i, j) {
			totalPaths++
			continue
		}

		if !maze.IsLastRow(i) {
			tc_down := maze.GetCost(i+1, j) + tc
			if tc_down <= cost {
				rec := []int{i + 1, j, tc_down}
				queue = append(queue, rec)
			}
		}

		if !maze.IsLastCol(j) {
			tc_left := maze.GetCost(i, j+1) + tc
			if tc_left <= cost {
				rec := []int{i, j + 1, tc_left}
				queue = append(queue, rec)
			}
		}
	}

	return totalPaths
}
