package main

import (
	"fmt"
	"time"
)

//func sum(number *int) {
// This one is kind of erratic, sometimes the second thread gets to execute and other times it doesnt, it needs to be inside of a loop, else it wont execute. Probably cuz main ends before the second thread gets to execute code
/*	for i := 0; i < 1; i++ {
		runtime.Gosched()
		*number++
		fmt.Println("Your new number is")
	}
*/
/*}

func subs(number *int) {
	runtime.Gosched()
	for i := 0; i < 1; i++ {
		*number--
	}
}

func main() {
	var tro int
	var tra int
	// GOMAXPROCS() doesnt seem to affect anything, moreover, without Gosched() it doesnt seem to be doing anything at all.
	runtime.GOMAXPROCS(2)
	go subs(&tra)
	subs(&tro)
	fmt.Println(tro)
	fmt.Println(tra)
}
*/

//Experimenting with channels
/*
func sum(to_sum *int, sum int, data chan int) {
	*to_sum = *to_sum + sum
	data <- *to_sum
}

func main() {
	a := 1
	tro := make(chan int, 2)
	go sum(&a, 1, tro)
	go sum(&a, 1, tro)
	_, _ = <-tro, <-tro
	fmt.Println(a)
	fmt.Println()
	// when you discharge a channel, (line 44), its capacity is restored to.
}
*/

//Experimenting with range in channels
/*
func sum(times int, to_sum *int, sum int, data chan int) {
	for i := 0; i < times; i++ {
		*to_sum = *to_sum + sum
		data <- *to_sum
	}
	close(data)
}

func main() {
	var a int
	c := make(chan int, 9)
	go sum(cap(c), &a, 1, c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(<-c)
}
*/

//Using select

func fibonacci(c, quit chan int) {
	x := 0
	for {
		select {
		case c <- x:
			x++
		case <-quit:
			fmt.Println("quit")
			return
		case <-time.After(5 * time.Second):
			fmt.Println("pics")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
	}()
	fibonacci(c, quit)
}
