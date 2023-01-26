package main

import "fmt"

func main() {
	number := 10
	fmt.Println(number)

	var getInt func(int) int
	getInt = func(x int) int {
		fmt.Println("In a 1st function")
		return 20 + x
	}
	g(getInt)
	getInt = func(x int) int {
		fmt.Println("In a 2nd function")
		return 10 + x
	}
	j := (19)
	fmt.Println(j)
}
func g(getInt func(int) int) {
	getInt(78)
}
