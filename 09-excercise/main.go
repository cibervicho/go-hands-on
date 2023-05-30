package main

import "fmt"

// variable at package level
var v1 string = "This is my variable v1"

const pi float64 = 3.1416

func main() {
	//fmt.Println("vim-go")
	v2 := "David M"

	fmt.Printf("v1 = %v\n", v1)
	fmt.Printf("pi = %v\n", pi)
	fmt.Printf("v2 = %v\n", v2)
}
