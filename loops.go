package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 4, 0}
	for _, value := range nums {
		if value == 3 {
			continue //break
		}
		fmt.Println(value)
	}
	//list:=range(2)
}
