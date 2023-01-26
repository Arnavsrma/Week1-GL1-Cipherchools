package main

import "fmt"

/*switch(data){
case var1:
	statement
case var2:
	statement
default:
	statement
}
*/

func main() {
	fmt.Print("Enter a Number: ")
	var input int
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println("You are even")
	} else {
		fmt.Println("You are odd")
	}

	data := 10
	switch data {
	case 10:
		data = 100
		fmt.Println(data)
		fallthrough //pickup the next case also
	case 111:
		data = 100
		fmt.Println(data)
		fallthrough
	case 112:
		data = 100
		fmt.Println(data)
	default:
		data = 10000
		fmt.Println(data)
	}
}
