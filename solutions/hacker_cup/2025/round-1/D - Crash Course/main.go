package main

import (
	"fmt"
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

func solveDp(s string) int {
	n := len(s)
	dpA := make([][]int, n+2)
	dpB := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dpA[i] = make([]int, n+2)
		dpB[i] = make([]int, n+2)
	}

	for length := 1; length <= n; length++ {
		for l := 1; l+length-1 <= n; l++ {
			r := l + length - 1

			resA := dpA[l+1][r]
			if s[l-1] == 'A' {
				resA = max(resA, dpB[l+1][r]^1)
			}
			dpA[l][r] = resA

			resB := dpB[l][r-1]
			if s[r-1] == 'B' {
				resB = max(resB, dpA[l][r-1]^1)
			}
			dpB[l][r] = resB
		}
	}

	return dpA[1][n]
}

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {
	answer := solveDp(input.S)
	if answer == 1 {
		output.answer = "Alice"
	} else {
		output.answer = "Bob"
	}
}

// --------------------------------------------------------------------------------
// INPUT
// --------------------------------------------------------------------------------

type TestCaseInput struct {
	N int
	S string
}

func readTestCaseInput(input *TestCaseInput) {
	readInput(&input.N)
	readInput(&input.S)
}

// --------------------------------------------------------------------------------
// OUTPUT
// --------------------------------------------------------------------------------

type TestCaseOutput struct {
	answer string
}

func NewTestCaseOutput() *TestCaseOutput {
	return &TestCaseOutput{}
}

func printTestCaseOutput(testID int, output *TestCaseOutput) {
	fmt.Printf("Case #%d: %s\n", testID, output.answer)
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
