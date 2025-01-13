package utils

import "fmt"

func MyFunc() {
	fmt.Println("mypack.MyFunc()")
}

func MySum(a ...int) (int, error) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum, nil
}
