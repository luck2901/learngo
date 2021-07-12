package main

import (
	"fmt"
)

func initOpMap() {
	opMap = make(map[string]func(int, int) int)

	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
	opMap["**"] = pow

}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}

func pow(a, b int) int {
	rst := 1
	for i := 0; i < b; i++ {
		rst *= a
	}
	return rst
}

func main() {
	initOpMap()
	Test()
}

var opMap map[string]func(int, int) int

func Calculate(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func Test() {
	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}
	if !testCalculate("Test2", "+", 5, 4, 9) {
		return
	}
	if !testCalculate("Test3", "-", 5, 3, 2) {
		return
	}
	if !testCalculate("Test4", "-", 3, 6, -3) {
		return
	}
	if !testCalculate("Test5", "*", 3, 7, 21) {
		return
	}
	if !testCalculate("Test6", "*", 3, 0, 0) {
		return
	}
	if !testCalculate("Test7", "*", 3, -3, -9) {
		return
	}
	if !testCalculate("Test8", "/", 9, 3, 3) {
		return
	}
	if !testCalculate("Test9", "**", 2, 3, 8) {
		return
	}
	if !testCalculate("Test9", "**", 2, 0, 1) {
		return
	}

	fmt.Println("Success")
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output: %d\n", testcase, expected, o)
		return false
	}
	return true
}
