package main

import "fmt"

func main() {
	c1 := incrementor("Foo:")
	c2 := incrementor("Bar:")
	c3 := puller(c1)
	c4 := puller(c2)
	fmt.Println("FInal sum", <-c3+<-c4)
}
func incrementor(s string) chan int {
	k := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			k <- i
			fmt.Println(s, i)
		}
		close(k)
	}()
	return k
}
func puller(m chan int) chan int {
	l := make(chan int)
	go func() {
		var sum int
		for x := range m {
			sum += x
		}
		l <- sum
		close(l)
	}()
	return l
}
