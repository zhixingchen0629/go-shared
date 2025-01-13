package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// factory function
type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// go routine
func TestCase1() {
	go fmt.Println("this is go routine 1")
	go fmt.Println("this is go routine 2")

	for i := 3; i < 10; i++ {
		go fmt.Printf("this is go routine %d\n", i)
	}

	fmt.Println("this is main routine")
	time.Sleep(time.Second * 2)
}

// defer LIFO
func TestCase2() {
	defer fmt.Println("this is defer 1")
	defer fmt.Println("this is defer 2")
	defer fmt.Println("this is defer 3")
	fmt.Println("this is main routine")
}

// defer execution order compared to return
// 1. 先return赋值
// 2. defer执行
// 3. return返回
func TestCase3() int {
	i := 9

	defer func() {
		i++
		fmt.Println("defer 1, i =", i)
	}()

	defer func() {
		i++
		fmt.Println("defer 2, i =", i)
	}()

	return i
}

// channel
func TestCase4() {
	ch := make(chan int)
	// fmt.Println(cap(ch))
	go func() {
		// send to channel
		// ch <- 1
		fmt.Println("send to channel")
	}()

	go func() {
		for {
			i := 0
			i = 1
			fmt.Println(i)
		}
	}()
	// receive from channel
	// fmt.Println("receive from channel")
	fmt.Println(<-ch)
}

// non-blocking channel
func TestCase5() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// buffered channel
func TestCase6() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	// ch <- 3 // deadlock
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
}

// unidirectional channel
func TestCase7() {
	ch := make(chan int)
	onlySend := (chan<- int)(ch)
	onlyReceive := (<-chan int)(ch)

	go func() {
		onlySend <- 1
		close(onlySend)
	}()

	go func() {
		fmt.Println(<-onlyReceive)
	}()

	// Wait for goroutines to finish
	time.Sleep(time.Second)
}

// select and channel
func TestCase8() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				fmt.Println("ch1", v)
			case v := <-ch2:
				fmt.Println("ch2", v)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
			ch2 <- i * 10
		}
	}()

	time.Sleep(time.Second)
}

// simulation for download file in multiple goroutines use select and channel
func TestCase9() {
	f_download := func(channel string) string {
		// sleep random time between 1-5 seconds
		time.Sleep(time.Second * time.Duration(1+time.Now().Unix()%5))
		return channel + " download success"
	}

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func() {
		ch1 <- f_download("channel 1")
	}()

	go func() {
		ch2 <- f_download("channel 2")
	}()

	go func() {
		ch3 <- f_download("channel 3")
	}()

	select {
	case res := <-ch1:
		fmt.Println(res)
	case res := <-ch2:
		fmt.Println(res)
	case res := <-ch3:
		fmt.Println(res)
	}

	time.Sleep(time.Second * 10)
}

// sync package use WaitGroup
func TestCase10() {
	f_download := func(channel string) string {
		// sleep random time between 1-5 seconds
		time.Sleep(time.Second * time.Duration(1+time.Now().Unix()%5))
		return channel + " download success"
	}

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		ch1 <- f_download("channel 1")
	}()

	go func() {
		defer wg.Done()
		ch2 <- f_download("channel 2")
	}()

	go func() {
		defer wg.Done()
		ch3 <- f_download("channel 3")
	}()

	channels := []<-chan string{ch1, ch2, ch3}
	for i := 0; i < len(channels); i++ {
		select {
		case res := <-ch1:
			fmt.Println(res)
		case res := <-ch2:
			fmt.Println(res)
		case res := <-ch3:
			fmt.Println(res)
		}
	}

	wg.Wait()
}

// sync package use infinite loop
func TestCase11() {
	f_download := func(channel string) string {
		// sleep random time between 1-5 seconds
		time.Sleep(time.Second * time.Duration(1+time.Now().Unix()%5))
		return channel + " download success"
	}

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ { // Simulate multiple messages
			ch1 <- f_download("channel 1")
		}
		close(ch1)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ { // Simulate multiple messages
			ch2 <- f_download("channel 2")
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ { // Simulate multiple messages
			ch3 <- f_download("channel 3")
		}
		close(ch3)
	}()

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
		close(ch3)
	}()

	for ch1 != nil || ch2 != nil || ch3 != nil {
		select {
		case res, ok := <-ch1:
			if ok {
				fmt.Println(res)
			} else {
				ch1 = nil
			}
		case res, ok := <-ch2:
			if ok {
				fmt.Println(res)
			} else {
				ch2 = nil
			}
		case res, ok := <-ch3:
			if ok {
				fmt.Println(res)
			} else {
				ch3 = nil
			}
		}
	}
}

// context package
func TestCase12() {

}

var globalNum = 0

func TestCase13() {
	rwMutex := sync.RWMutex{}
	var wg sync.WaitGroup
	var wg1 sync.WaitGroup

	wg.Add(50)
	wg1.Add(50)

	for i := 0; i < 100; i++ {
		go func() {
			wg1.Wait()
			rwMutex.Lock()
			defer rwMutex.Unlock()

			globalNum++

			wg.Done()
		}()
	}

	wg.Wait()
	// time.Sleep(time.Second * 2)

	context.WithCancel()

	fmt.Println(globalNum)
}

// // condition variable
// func TestCase14() {
// 	// cond := sync.Cond{}

// }

func main() {
	// TestCase1()
	// TestCase2()
	// fmt.Println(TestCase3())
	// TestCase4()
	// TestCase5()
	// TestCase6()
	// TestCase7()
	// TestCase8()
	// TestCase9()
	// TestCase10()
	// TestCase11()
	TestCase13()
}
