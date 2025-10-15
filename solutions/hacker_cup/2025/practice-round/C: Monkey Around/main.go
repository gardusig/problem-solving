package main

import (
	"fmt"
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

type Pair struct {
	first  int
	second int
}

type Group struct {
	initialValue int
	maxValue     int
}

func getGroups(input *TestCaseInput) []*Group {
	groups := make([]*Group, 0)
	initialValue := input.A[0] - 1
	maxValue := input.A[0]
	seenValues := make(map[int]bool)
	seenValues[input.A[0]] = true
	for idx := 1; idx < input.N; idx += 1 {
		value := input.A[idx]
		_, hasSeen := seenValues[value]
		if hasSeen || input.A[idx] > input.A[idx-1]+1 {
			groups = append(groups, &Group{
				initialValue: initialValue,
				maxValue:     maxValue,
			})
			initialValue = value - 1
			maxValue = value
			seenValues = make(map[int]bool)
		}
		seenValues[value] = true
		maxValue = max(maxValue, value)
	}
	groups = append(groups, &Group{initialValue: initialValue, maxValue: maxValue})
	return groups
}

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {
	groups := getGroups(input)
	moves := make([]int, len(groups))
	totalMoves := 0
	for idx := len(groups) - 1; idx >= 0; idx -= 1 {
		current := totalMoves % groups[idx].maxValue
		if current > groups[idx].initialValue {
			moves[idx] = groups[idx].maxValue - current + groups[idx].initialValue
		} else {
			moves[idx] = groups[idx].initialValue - current
		}
		totalMoves += moves[idx]
	}
	for idx, group := range groups {
		output.answer = append(output.answer, Pair{1, group.maxValue})
		for i := 0; i < moves[idx]; i += 1 {
			output.answer = append(output.answer, Pair{2, 0})
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
	answer []Pair
}

func NewTestCaseOutput() *TestCaseOutput {
	return &TestCaseOutput{
		answer: make([]Pair, 0),
	}
}

func printTestCaseOutput(testID int, output *TestCaseOutput) {
	fmt.Printf("Case #%d: %d\n", testID, len(output.answer))
	for _, element := range output.answer {
		if element.first == 1 {
			fmt.Printf("%d %d\n", 1, element.second)
		} else {
			fmt.Printf("2\n")
		}
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
