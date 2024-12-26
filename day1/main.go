package main

import (
	"fmt"
	"strconv"
)

// hello world
func TestCase1() {
	fmt.Println("Hello, world!")
}

// variable declaration in 4 ways
func TestCase2() {
	// 1
	var a1 int
	a1 = 1

	// 2
	var a2 int = 2

	// 3
	var a3 = 3

	// 4
	a4 := 4

	fmt.Println(a1, a2, a3, a4)
}

// basic data types
func TestCase3() {
	var _uint8 uint8 = 1
	var _uint16 uint16 = 2
	var _uint32 uint32 = 3
	var _uint64 uint64 = 4
	var _uint uint = 5
	fmt.Println(_uint8, _uint16, _uint32, _uint64, _uint)

	var _int8 int8 = 1
	var _int16 int16 = 2
	var _int32 int32 = 3
	var _int64 int64 = 4
	var _int int = 5
	fmt.Println(_int8, _int16, _int32, _int64, _int)

	var _float32 float32 = 1.1
	var _float64 float64 = 2.2
	fmt.Println(_float32, _float64)

	var _bool bool = true
	fmt.Println(_bool)

	var _string string = "hello world"
	fmt.Println(_string)
}

// default value nil
func TestCase4() {
	var _int int
	var _float32 float32
	var _bool bool
	var _string string
	fmt.Println(_int, _float32, _bool, _string)
}

// pointer
func TestCase5() {
	var a int = 10
	var b *int = &a
	fmt.Println(a, b, *b)
	*b = 20
	fmt.Println(a, b, *b)
}

// const
func TestCase6() {
	const _int int = 10
	const _float32 float32 = 1.1
	const _bool bool = true
	const _string string = "hello world"

	// _int = 20
	fmt.Println(_int, _float32, _bool, _string)
}

// iota
func TestCase7() {
	const (
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one, two, three, four)
}

// convert between string and int
func TestCase8() {

	// int to string
	_int := 10
	fmt.Printf("Type of _int is %T, value is %d\n", _int, _int)

	_string := strconv.Itoa(_int)
	fmt.Printf("Type of _string is %T, value is %s\n", _string, _string)

	_intReverse, ok := strconv.Atoi(_string)
	fmt.Printf("Type of _intReverse is %T, %v, value is %d\n", _intReverse, ok, _intReverse)

	// float to string
	_float64 := 1.1
	fmt.Printf("Type of _float32 is %T, value is %f\n", _float64, _float64)

	_stringFloat64 := strconv.FormatFloat(_float64, 'f', -1, 64)
	fmt.Printf("Type of _stringFloat32 is %T, value is %s\n", _stringFloat64, _stringFloat64)

	_float64Reverse, err := strconv.ParseFloat(_stringFloat64, 64)
	fmt.Printf("Type of _float64Reverse is %T, %v, value is %f\n", _float64Reverse, err, _float64Reverse)

	// bool to string
	_bool := true
	fmt.Printf("Type of _bool is %T, value is %t\n", _bool, _bool)

	_stringBool := strconv.FormatBool(_bool)
	fmt.Printf("Type of _stringBool is %T, value is %s\n", _stringBool, _stringBool)

	_boolReverse, err := strconv.ParseBool(_stringBool)
	fmt.Printf("Type of _boolReverse is %T, %v, value is %t\n", _boolReverse, err, _boolReverse)
}

// if control block
func TestCase9() {
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	} else {
		fmt.Println("a is less than 5")
	}
}

// switch control block
func TestCase10() {
	a := 10
	switch a {
	case 1:
		fmt.Println("a is 1")
		fallthrough
	case 2:
		fmt.Println("a is 2")
	case 3:
		fmt.Println("a is 3")
	default:
		fmt.Println("a is not 1, 2, 3")
	}
}

// for loop
func TestCase11() {
	for i := 0; i < 10; i++ {
		fmt.Println("this is ", i, " line")
	}
}

// for loop implement while loop in other language
func TestCase12() {
	i := 0
	for i < 10 {
		fmt.Println("this is ", i, " line")
		i++
	}
}

// array
func TestCase13() {
	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array)

	array1 := [5]string{1: "b", 3: "d"}
	fmt.Println(array1)

	array2 := [...]string{1: "b", 3: "d"}
	fmt.Println(array2)

	for i := 0; i < len(array); i++ {
		fmt.Println("array: ", array[i], " array1: ", array1[i])
	}

	// range loop
	for i, v := range array {
		fmt.Println("array: ", v, " array1: ", array1[i])
	}

	fmt.Println(len(array2), array2[3])
}

// slice base on array
func TestCase14() {
	array := [5]string{"a", "b", "c", "d", "e"}
	slice := array[2:5]
	slice[1] = "f"
	fmt.Println(slice)

	for i, v := range array {
		fmt.Println(i, v)
	}

	slice1 := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice1)
}

// slice
func TestCase15() {
	slice := make([]string, 5, 8)
	fmt.Println(slice)

	slice = append(slice, "f")
	fmt.Println(slice)

	fmt.Println(len(slice), cap(slice))
}

func main() {
	TestCase1()
	TestCase2()
	TestCase3()
	TestCase4()
	TestCase5()
	TestCase6()
	TestCase7()
	TestCase8()
	TestCase9()
	TestCase10()
	TestCase11()
	TestCase12()
	TestCase13()
	TestCase14()
	TestCase15()
}
