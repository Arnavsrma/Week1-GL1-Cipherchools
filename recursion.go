package main

import "fmt"

func main() {

	recn(4)
}
func rec(number int) {
	if number == 0 {
		return
	}
	if number%2 == 0 {
		fmt.Println(number + 1)
		rec(number - 1)
	} else {
		fmt.Println(number + 2)
		rec(number - 1)
	}
	fmt.Println(number - 1)
	//output = 75533
}
func recsn(number int) {
	if number == 0 {
		return
	}
	if number%2 == 0 {
		recsn(number - 1)
		fmt.Println(number - 1)
	} else {
		recsn(number - 1)
		fmt.Println(number - 1)
	}
	fmt.Println(number - 1)
	//output for recsn(5)=0011223344
}
func recn(number int) {
	if number <= 0 {
		return
	}
	recn(number - 1)
	recn(number - 2)
	fmt.Println(number - 1)
	//output for recn(4)=0102013
}
func fib(number int) int {
	if number == 0 || number == 1 {
		return number
	}
	result := fib(number-1) + fib(number-2)
	return result
}
func fact(number int) int {
	if number == 1 || number == 0 {
		return 1
	}
	if number < 0 {
		fmt.Println("Invalid Number")
		return -1
	}
	result := number * fact(number-1)
	return result
}
