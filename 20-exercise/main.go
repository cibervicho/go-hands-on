package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	x := rand.Intn(400)
	fmt.Println("x =", x)

	switch {
	case x <= 100:
		fmt.Printf("%v is between 0 and 100\n", x)
	case x >= 101 && x <= 200:
		fmt.Printf("%v is between 101 and 200\n", x)
	case x >= 201 && x <= 250:
		fmt.Printf("%v is between 201 and 250\n", x)
	default:
		fmt.Printf("%v is more than 250\n", x)

	}
}
