package main

import "fmt"

func main() {
	c := make(chan int)
	d := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		d <- true
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		d <- true
	}()
	go func() {
		<-d
		<-d
		close(c)
	}()
	for x := range c {
		fmt.Println(x)
	}
}
