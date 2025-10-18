package main

import (
	"fmt"
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

func dfs(idx int, visited []bool, height int, input *TestCaseInput) {
	visited[idx] = true
	for _, nxtIdx := range []int{idx - 1, idx + 1} {
		if nxtIdx < 0 || nxtIdx >= input.N {
			continue
		}
		if visited[nxtIdx] {
			continue
		}
		diff := input.arr[idx] - input.arr[nxtIdx]
		if diff < 0 {
			diff = -diff
		}
		if height < diff {
			continue
		}
		dfs(nxtIdx, visited, height, input)
	}
}

func isValid(input *TestCaseInput, height int) bool {
	visited := make([]bool, input.N)
	for i := 0; i < input.N; i += 1 {
		if visited[i] {
			continue
		}
		if height >= input.arr[i] {
			dfs(i, visited, height, input)
		}
	}
	for i := 0; i < input.N; i += 1 {
		if !visited[i] {
			return false
		}
	}
	return true
}

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {
	lo := 1
	hi := 1123456789
	for lo <= hi {
		mid := lo + ((hi - lo) >> 1)
		if isValid(input, mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	output.answer = hi + 1
}

// --------------------------------------------------------------------------------
// INPUT
// --------------------------------------------------------------------------------

type TestCaseInput struct {
	N   int
	arr []int
}

func readTestCaseInput(input *TestCaseInput) {
	readInput(&input.N)

	input.arr = make([]int, input.N)
	readInput(input.arr)
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
		go solveTestCaseAsync(wg, testCaseInput, testCasesOutput[testID-1])
	}
	wg.Wait()

	for testID, testOutput := range testCasesOutput {
		printTestCaseOutput(testID+1, testOutput)
	}
}

// --------------------------------------------------------------------------------
// DO NOT TOUCH
// --------------------------------------------------------------------------------
