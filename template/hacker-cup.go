package template

import (
	"fmt"
	"reflect"
	"sync"
)

// --------------------------------------------------------------------------------
// SOLUTION
// --------------------------------------------------------------------------------

func solveTestCase(input *TestCaseInput, output *TestCaseOutput) {

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
	invalid bool
}

func NewTestCaseOutput() *TestCaseOutput {
	return &TestCaseOutput{
		invalid: false,
	}
}

func printTestCaseOutput(testID int, output *TestCaseOutput) {
	if output.invalid {
		fmt.Println("Case #", testID, ": ", -1)
		return
	}
	fmt.Println("Case #", testID, ": ", output.invalid)
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
