package main

import (
	"fmt"
	"reflect"
	"sort"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

type Pair struct {
	first  int
	second int
}

type PairList []Pair

func (s PairList) Len() int           { return len(s) }
func (s PairList) Less(i, j int) bool { return s[i].first < s[j].first }
func (s PairList) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {
	availableValues := make(map[int]int)
	for i, currentTemperature := range input.A {
		availableValues[currentTemperature] = i
	}

	for i := 0; i < input.N; i += 1 {
		if input.A[i] > input.B[i] {
			output.invalid = true
			return
		}
		_, ok := availableValues[input.B[i]]
		if !ok {
			output.invalid = true
			return
		}
	}

	bValues := make(PairList, input.N)
	for i, value := range input.B {
		bValues[i] = Pair{first: value, second: i}
	}

	sort.Sort(bValues)

	for _, goal := range bValues {
		idx := goal.second
		if input.A[idx] == input.B[idx] {
			continue
		}
		output.pairs = append(output.pairs,
			Pair{
				idx + 1,
				availableValues[input.B[idx]] + 1,
			},
		)
	}
}

// --------------------------------------------------------------------------------
// INPUT
// --------------------------------------------------------------------------------

type TestCaseInput struct {
	N int
	A []int
	B []int
}

func readTestCaseInput(input *TestCaseInput) {
	readInput(&input.N)

	input.A = make([]int, input.N)
	readInput(input.A)

	input.B = make([]int, input.N)
	readInput(input.B)
}

// --------------------------------------------------------------------------------
// OUTPUT
// --------------------------------------------------------------------------------

type TestCaseOutput struct {
	invalid bool
	pairs   []Pair
}

func NewTestCaseOutput() *TestCaseOutput {
	return &TestCaseOutput{
		invalid: false,
		pairs:   make([]Pair, 0),
	}
}

func printTestCaseOutput(testID int, output *TestCaseOutput) {
	if output.invalid {
		fmt.Printf("Case #%d: %d\n", testID, -1)
		return
	}
	fmt.Printf("Case #%d: %d\n", testID, len(output.pairs))
	for _, pair := range output.pairs {
		fmt.Printf("%d %d\n", pair.first, pair.second)
	}
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
