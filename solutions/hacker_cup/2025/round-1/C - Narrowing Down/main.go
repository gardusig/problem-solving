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
	bitPrefixSum := make([][]int, input.N)
	for i := 0; i < input.N; i += 1 {
		bitPrefixSum[i] = make([]int, MAX_SIZE)
		for bit := 0; bit < MAX_SIZE; bit += 1 {
			bitPrefixSum[i][bit] = (input.A[i] & (1 << bit))
			if bitPrefixSum[i][bit] > 0 {
				bitPrefixSum[i][bit] = 1
			}
			if i > 0 {
				bitPrefixSum[i][bit] += bitPrefixSum[i-1][bit]
			}
		}
	}

	output.answer = 0
	for i := 0; i < input.N; i += 1 {
		size := input.N - i
		output.answer += ((1 + size) * size) / 2
	}

	for i := 0; i < input.N; i += 1 {
		for bit := 0; bit < MAX_SIZE; bit += 1 {
			for j := i + 1; j < input.N; j += 1 {
				total := bitPrefixSum[j][bit]
				if i > 0 {
					total -= bitPrefixSum[i-1][bit]
				}
				if (total & 1) == 0 {
					fmt.Println("bitPrefixSum:", input.A, ", bit:", bit, ", i:", i, ", j:", j)
				}
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
