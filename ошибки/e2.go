package main

import "fmt"

func main() {
	var ans1, ans2, ans3 string
	fmt.Print("Name: ")
	fmt.Scan(&ans1)

	fmt.Print("Fav food: ")

	fmt.Scan(&ans2)

	fmt.Print("Fav movie: ")
	fmt.Scan(&ans3)

	fmt.Println(ans1, ans2, ans3)
}
