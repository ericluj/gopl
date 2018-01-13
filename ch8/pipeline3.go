package main

import "fmt"
import "time"

func counter(out chan<- int) {
	for i := 0; i <= 5; i++ {
		out <- i
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func square(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for s := range in {
		fmt.Println(s)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go square(squares, naturals)
	printer(squares)
}