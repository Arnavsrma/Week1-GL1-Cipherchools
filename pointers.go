package main

import "fmt"

func main() {
	var i int
	i = 10
	var j *int
	j = &i
	*j = 100

	name := new(string)
	*name = "GoLang"
	fmt.Println(*name)
}
