package main

import "fmt"

func main() {
	// x := 36
	// fmt.Println(&x)
	// y := &x
	// fmt.Println(y)
	// fmt.Println(x)
	// fmt.Println(*y)
	// *y = 25
	// fmt.Println(x)
	x := 36
	changeValueByReference(&x)
	fmt.Println(x)
}

func changeValue(x int) {
	fmt.Println("Changing Value for X")
	fmt.Println(&x)
	x = 25
}

func changeValueByReference(x *int) {
	fmt.Println("Updating value of X")
	*x = 10
}
