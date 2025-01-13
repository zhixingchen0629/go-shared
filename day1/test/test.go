package test

import "fmt"

type TestStruct struct {
	Name string
}

func TestInTest() {
	fmt.Println("hello world")
}

func testInTest() {
	fmt.Println("hello world 2")
}

func Test3() {
	testInTest()
}
