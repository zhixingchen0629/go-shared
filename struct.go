package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     uint
	Address Address
}

type Address struct {
	Province string
	City     string
}

func (p *Person) String() string {
	return fmt.Sprintf("the name is %s, age is %d, address %s", p.Name, p.Age, p.Address.String())
}

func (p *Person) PrintAge() {
	fmt.Println(p.Name, "age is", p.Age)
}

func (a *Address) String() string {
	return fmt.Sprintf("the province is %s, city is %s", a.Province, a.City)
}

func PrintString(s fmt.Stringer) {
	fmt.Println(s.String())
}

// func main() {
// 	var p Person
// 	p.Name = "lisi"
// 	p.Age = 60
// 	p.Address.Province = "guangdong"
// 	p.Address.City = "shenzhen"
// 	p.PrintAge()
// 	fmt.Println("stringer interface", p.String())

// 	person := Person{"zhangsan", 35, Address{"beijing", "beijing"}}
// 	person.PrintAge()
// 	fmt.Println("hello world")

// 	p2 := Person{Name: "zhaoliu", Age: 30, Address: Address{Province: "guangdong", City: "guangzhou"}}
// 	p2.PrintAge()

// 	pPerson := &Person{"wangwu", 40, Address{"shanghai", "shanghai"}}
// 	pPerson.PrintAge()

// 	PrintString(&p.Address)
// 	PrintString(&p)

// 	// ctx := context

// 	// var ch chan int

// 	// ch = make(chan int, 5)

// 	// str := "str test"
// 	// i, err := strconv.Atoi(str)

// 	// errors.New()

// 	// errors.Is()

// 	// f := os.ReadFile()
// }
