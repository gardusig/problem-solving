package main

import (
	"fmt"
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

var DELTA_X = []int{0, 0, -1, +1}
var DELTA_Y = []int{-1, +1, 0, 0}

const MAX_INT = 1023456789

type GridPosition struct {
	row    int
	column int
}

func dfs(row int, column int, S int, visited [][]bool, distance [][]int) int {
	visited[row][column] = true
	total := 1

	for delta := 0; delta < 4; delta += 1 {
		nextRow := row + DELTA_X[delta]
		nextColumn := column + DELTA_Y[delta]

		if nextRow < 0 || nextRow >= len(distance) {
			continue
		}

		if nextColumn < 0 || nextColumn >= len(distance[row]) {
			continue
		}

		if !visited[nextRow][nextColumn] && distance[nextRow][nextColumn] >= S {
			total += dfs(nextRow, nextColumn, S, visited, distance)
		}
	}

	return total
}

func getLargestContiguousSpace(distance [][]int, S int) int {
	visited := make([][]bool, len(distance))
	for i := 0; i < len(distance); i += 1 {
		visited[i] = make([]bool, len(distance[i]))
	}

	largest := 0
	for i := 0; i < len(distance); i += 1 {
		for j := 0; j < len(distance[i]); j += 1 {
			if !visited[i][j] && distance[i][j] >= S {
				contiguousSpace := dfs(i, j, S, visited, distance)
				largest = max(largest, contiguousSpace)
			}
		}
	}

	return largest
}

func getDistance(input *TestCaseInput) [][]int {
	distance := make([][]int, input.R)
	for row := 0; row < input.R; row += 1 {
		distance[row] = make([]int, input.C)
		for column := 0; column < input.C; column += 1 {
			distance[row][column] = MAX_INT
		}
	}

	queue := make([]GridPosition, 0)
	for row := 0; row < input.R; row += 1 {
		for column := 0; column < input.C; column += 1 {
			if input.grid[row][column] != '#' {
				continue
			}
			distance[row][column] = -1
			queue = append(queue, GridPosition{row: row, column: column})
		}
	}
	for row := 0; row < input.R; row += 1 {
		for column := 0; column < input.C; column += 1 {
			if row == 0 || row == input.R-1 || column == 0 || column == input.C-1 {
				if input.grid[row][column] == '#' {
					continue
				}
				distance[row][column] = 0
				queue = append(queue, GridPosition{row: row, column: column})
			}
		}
	}

	for queueIndex := 0; queueIndex < len(queue); queueIndex += 1 {
		row := queue[queueIndex].row
		column := queue[queueIndex].column
		for delta := 0; delta < 4; delta += 1 {
			nextRow := row + DELTA_X[delta]
			nextColumn := column + DELTA_Y[delta]

			if nextRow < 0 || nextRow >= input.R {
				continue
			}

			if nextColumn < 0 || nextColumn >= input.C {
				continue
			}

			if distance[nextRow][nextColumn] > distance[row][column]+1 {
				distance[nextRow][nextColumn] = distance[row][column] + 1
				queue = append(queue, GridPosition{row: nextRow, column: nextColumn})
			}
		}
	}

	return distance
}

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {
	distance := getDistance(input)
	output.answer = getLargestContiguousSpace(distance, input.S)
}

// --------------------------------------------------------------------------------
// INPUT
// --------------------------------------------------------------------------------

type TestCaseInput struct {
	R int
	C int
	S int

	grid []string
}

func readTestCaseInput(input *TestCaseInput) {
	readInput(&input.R)
	readInput(&input.C)
	readInput(&input.S)

	input.grid = make([]string, input.R)
	for i := 0; i < input.R; i += 1 {
		readInput(&input.grid[i])
	}

}

// --------------------------------------------------------------------------------
// OUTPUT
// --------------------------------------------------------------------------------

type TestCaseOutput struct {
	answer int
}

func NewTestCaseOutput() *TestCaseOutput {
	return &TestCaseOutput{
		answer: 0,
	}
}

func printTestCaseOutput(testID int, output *TestCaseOutput) {
	fmt.Printf("Case #%d: %d\n", testID, output.answer)
}

// --------------------------------------------------------------------------------
// DO NOT TOUCH
// --------------------------------------------------------------------------------

func readInput(variable any) {
	v := reflect.ValueOf(variable)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice {
		readInputValue(v)
		return
	}
	for i := range v.Len() {
		readInputValue(v.Index(i))
	}
}

func readInputValue(v reflect.Value) {
	_, err := fmt.Scan(v.Addr().Interface())
	if err != nil {
		panic(err)
	}
}

func solveTestCaseAsync(wg *sync.WaitGroup, input *TestCaseInput, output *TestCaseOutput) {
	defer wg.Done()
	solveTestCase(input, output)
}

func main() {
	var testCases int
	readInput(&testCases)
	testCasesOutput := make([]*TestCaseOutput, testCases)

	wg := new(sync.WaitGroup)
	for testID := 1; testID <= testCases; testID += 1 {
		testCaseInput := new(TestCaseInput)
		readTestCaseInput(testCaseInput)
		testCasesOutput[testID-1] = NewTestCaseOutput()

		wg.Add(1)
		solveTestCaseAsync(wg, testCaseInput, testCasesOutput[testID-1])
	}
	wg.Wait()

	for testID, testOutput := range testCasesOutput {
		printTestCaseOutput(testID+1, testOutput)
	}
}

// --------------------------------------------------------------------------------
// DO NOT TOUCH
// --------------------------------------------------------------------------------
