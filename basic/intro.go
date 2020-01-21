package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World!!")
	//Declaring a variable
	var x int
	x = 14
	fmt.Println(x)
	x = 25
	fmt.Println(x)
	//x = 25.7
	//fmt.Println(x)
	//var y string
	//y = "Hello World"
	y := "Hello World"
	fmt.Println(y)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()
	fmt.Println(name)
}
