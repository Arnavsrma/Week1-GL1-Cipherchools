package main

import "fmt"

func main() {
	result, x := c(10, "go")
	fmt.Println(result)
	fmt.Println(x)
	//b(1, 2, 3, 4, 5, 6, 7)
}

func a() (int, string) {
	return 122, "afdfs"
}
func b(args ...int) (bool, int) {
	for _, v := range args {
		fmt.Print(v)
	}
	return true, 11
}
func c(w int, name string) (i int, j string) {
	i = 10
	j = "rahul"
	return
}
