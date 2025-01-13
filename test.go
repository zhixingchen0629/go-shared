package main

import (
	"fmt"
)

// const (
// 	one int    = 1
// 	two string = "two"
// )

const (
	one = iota + 1
	two
	three
)

func add(a int, b int) (int, error) {
	return a + b, nil
}

func Sum(a ...int) (int, error) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum, nil
}

var globalJ = ""

var globalA *int

type A struct {
}

func colsure() func() int {
	i := 0
	// j := 1
	// fmt.Printf("%p\n", &j)
	// globalJ = fmt.Sprintf("%p", &j)
	fmt.Printf("%p\n", &i)

	f := func() int {
		// fmt.Println("hello world")
		// return 1
		i++
		fmt.Printf("%p\n", &i)
		return i
	}
	// i = 1

	// fmt.Printf("%p\n", &i)
	return f
}

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func main() {
	// 	// 	// var a int32 = 8
	// 	// 	// var b int16 = 10

	// 	// 	// fmt.Println(a + int32(b))
	// 	// 	// fmt.Println(a)
	// 	// 	// fmt.Println("hello world")
	// 	// 	// fmt.Println(one)
	// 	// 	// fmt.Println(two)

	// 	// 	// var c int = 10
	// 	// 	// i2s := strconv.Itoa(c)
	// 	// 	// fmt.Println("this is string", i2s)
	// 	// 	// s2i, err := strconv.Atoi(i2s)
	// 	// 	// if err != nil {
	// 	// 	// 	fmt.Println("error")
	// 	// 	// 	return
	// 	// 	// }
	// 	// 	// fmt.Println("this is number", s2i)

	// 	// 	// s1 := "Hello World"
	// 	// 	// fmt.Println(strings.HasPrefix(s1, "Hello"))
	// 	// 	// fmt.Println(strings.Index(s1, "o"))
	// 	// 	// fmt.Println(strings.ToUpper(s1))

	// 	// 	// strconv.FormatInt()
	// 	// 	// strconv.ParseInt()

	// 	// 	// i := 6

	// 	// 	// if i > 10 {
	// 	// 	// 	fmt.Println("i>10")
	// 	// 	// } else if i > 5 && i <= 10 {
	// 	// 	// 	fmt.Println("i>5 i<=10")
	// 	// 	// } else {
	// 	// 	// 	fmt.Println("i<=5")
	// 	// 	// }

	// 	// 	// switch i := 6; {
	// 	// 	// case i > 10:
	// 	// 	// 	fmt.Println("i>10")
	// 	// 	// case i > 5 && i <= 10:
	// 	// 	// 	fmt.Println("i>5 i<=10")
	// 	// 	// 	fallthrough
	// 	// 	// default:
	// 	// 	// 	fmt.Println("i<=5")
	// 	// 	// }

	// 	// 	// i := 0
	// 	// 	// for {
	// 	// 	// 	println("this is ", i, " line")
	// 	// 	// 	i++
	// 	// 	// }

	// 	// 	// sum := 0
	// 	// 	// for i := 1; i <= 100; i++ {
	// 	// 	// 	if i%2 == 0 {
	// 	// 	// 		continue
	// 	// 	// 	}
	// 	// 	// 	sum += i
	// 	// 	// 	fmt.Println("this is ", i)
	// 	// 	// }

	// 	// 	// array4 := [5]string{"a", "b", "c", "d", "e"}
	// 	// 	// array := [...]string{"a", "b", "c", "d", "e"}
	// 	// 	// array1 := [5]string{1: "b", 3: "d"}
	// 	// 	// array2 := [...]string{1: "b", 3: "d"}

	// 	// 	// for i := 0; i < len(array); i++ {
	// 	// 	// 	fmt.Println("array: ", array[i], " array1: ", array1[i])
	// 	// 	// }

	// 	// 	// fmt.Println(len(array2), len(array4), array4[3])

	// 	// 	// array := [5]string{"a", "b", "c", "d", "e"}
	// 	// 	// slice := array[2:5]
	// 	// 	// slice[1] = "f"

	// 	// 	// for i, v := range array {
	// 	// 	// 	fmt.Println(i, v)
	// 	// 	// }

	// 	// 	// slice := make([]string, 5, 8)
	// 	// 	slice1 := []string{"a", "b", "c", "d", "e"}
	// 	// 	fmt.Println(len(slice1), cap(slice1))

	// 	// 	slice2 := append(slice1, "f")
	// 	// 	slice3 := append(slice1, "f", "g")
	// 	// 	slice4 := append(slice1, slice2...)
	// 	// 	fmt.Println(slice1, slice2, slice3, slice4)

	// 	// 	nameAgeMap := make(map[string]int)
	// 	// 	nameAgeMap["a"] = 1
	// 	// 	nameAgeMap["a"] = 10
	// 	// 	fmt.Println(nameAgeMap)

	// 	// 	value, ok := nameAgeMap["a"]
	// 	// 	if ok {
	// 	// 		fmt.Println(value, ok)
	// 	// 	}

	// 	// 	// nameAgeMap1 := map[string]int{"a": 1, "b": 2, "c": 3}

	// 	// 	a2 := make([][]int, 5)
	// 	// 	a2[1] = append(a2[0], 5)
	// 	// 	fmt.Println(len(a2), cap(a2), len(a2[0]), cap(a2[0]))

	// 	// 	aa := [...][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// 	// 	fmt.Println(aa)

	// 	// 	add1, add2 := 1, 2

	// 	// 	sum, err := add(add1, add2)
	// 	// 	if err != nil {
	// 	// 		fmt.Println("error")
	// 	// 	}

	// 	// 	fmt.Println("1 + 2 =", sum)

	// 	// 	total, err := Sum(1, 2, 3, 4)
	// 	// 	if err != nil {
	// 	// 		fmt.Println("error")
	// 	// 	}
	// 	// 	fmt.Println("1 + 2 + 3 + 4 =", total)

	// 	// 	total1, err := Sum([]int{1, 2, 3, 4, 5}...)
	// 	// 	if err != nil {
	// 	// 		fmt.Println("error")
	// 	// 	}
	// 	// 	fmt.Println("1 + 2 + 3 + 4 + 5 =", total1)

	// 	// 	total2, err := utils.MySum(1, 2, 3)
	// 	// 	if err != nil {
	// 	// 		fmt.Println("error")
	// 	// 	}
	// 	// 	fmt.Println("1 + 2 + 3 =", total2)
	// 	// 	// fmt.Println()

	cl := colsure()
	fmt.Println(cl())
	fmt.Println(cl())

	// fmt.Println(cl())
	// fmt.Println(cl())

	// 	// c2 := colsure()
	// 	// fmt.Println(c2())

	// 	// c3 := colsure()
	// 	// fmt.Println(c3())

	// 	// fmt.Println(c2())
	// 	// fmt.Println(cl())

	// 	// // 假设这是一个内存地址的字符串
	// 	// addressStr := globalJ // 这里用一个示例地址，你需要替换为实际地址

	// 	// // 将字符串转换为整型表示的地址
	// 	// address, err := strconv.ParseUint(addressStr, 0, 64)
	// 	// if err != nil {
	// 	// 	fmt.Println("地址转换错误:", err)
	// 	// 	return
	// 	// }

	// 	// // 将 uint64 类型的地址转换为指针
	// 	// ptr := unsafe.Pointer(uintptr(address))

	// 	// // 将指针转换为 int 类型的指针
	// 	// intPtr := (*int)(ptr)

	// 	// // 读取该地址上的整数值
	// 	// // 注意：在解引用之前，你需要确保该地址是有效的并且可以安全地读取。
	// 	// value := *intPtr

	// 	// fmt.Printf("地址 %s 上的整数值是: %d\n", addressStr, value)

	// 	// fmt.Println("===")
	// 	// fmt.Println(colsure()())
	// 	// fmt.Println(colsure()())
	// 	// fmt.Println(colsure()())

	// 	// 	age1 := Age(10)
	// 	// 	age1.String()

}
