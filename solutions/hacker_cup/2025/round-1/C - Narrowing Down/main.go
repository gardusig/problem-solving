package main

import (
	"fmt"
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

const MAX_SIZE = 30

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {
	bitAccumulatedFrequency := make([][]int, input.N)
	for i := 0; i < input.N; i += 1 {
		bitAccumulatedFrequency[i] = make([]int, MAX_SIZE)
		for bit := 0; bit < MAX_SIZE; bit += 1 {
			if (input.A[i] & (1 << bit)) != 0 {
				bitAccumulatedFrequency[i][bit] = 1
			} else {
				bitAccumulatedFrequency[i][bit] = 0
			}
			if i > 0 {
				bitAccumulatedFrequency[i][bit] += bitAccumulatedFrequency[i-1][bit]
			}
		}
	}

	output.answer = 0
	for i := 0; i < input.N; i += 1 {
		leftmostPositive := -1
		rightmostPositive := -1
		for j := i; j < input.N; j += 1 {
			isValid := true

			if input.A[j] > 0 {
				if leftmostPositive == -1 {
					leftmostPositive = j
				}
				rightmostPositive = j
			}

			for bit := 0; bit < MAX_SIZE; bit += 1 {
				total := bitAccumulatedFrequency[j][bit]
				if i > 0 {
					total -= bitAccumulatedFrequency[i-1][bit]
				}
				if (total & 1) != 0 {
					isValid = false
				}
			}
			if !isValid {
				output.answer += (j - i + 1)
				continue
			}

			if leftmostPositive != -1 && rightmostPositive != -1 {
				output.answer += (rightmostPositive - leftmostPositive)
			}
		}
	}
}

// --------------------------------------------------------------------------------
// INPUT
// --------------------------------------------------------------------------------

type TestCaseInput struct {
	N int
	A []int
}

func readTestCaseInput(input *TestCaseInput) {
	readInput(&input.N)

	input.A = make([]int, input.N)
	readInput(input.A)
}

// --------------------------------------------------------------------------------
// OUTPUT
// --------------------------------------------------------------------------------

type TestCaseOutput struct {
	answer int
}

func NewTestCaseOutput() *TestCaseOutput {
	return &TestCaseOutput{}
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
