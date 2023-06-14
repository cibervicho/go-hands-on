package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const COUNT = 50

func main() {
	fmt.Println("======== PART 1 ========")

	for i := 0; i < COUNT; i++ {
		if (i != 0 && i%10 == 0) || i == COUNT-1 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%v, ", i)
		}
	}

	fmt.Println("\n======== PART 2 ========")

	for i := 0; i < COUNT; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)

		fmt.Printf("Iteration: [%v]\n\tx = %v\n\ty = %v\n", i+1, x, y)

		switch {
		case x < 4 && y < 4:
			fmt.Printf("\tx and y are both less than 4\n")
		case x > 6 && y > 6:
			fmt.Printf("\tx and y are both greater than 6\n")
		case x >= 4 && x <= 6:
			fmt.Printf("\tx is greater than or equal to 4 and less than or equal to 6\n")
		case y != 5:
			fmt.Printf("\ty is not 5\n")
		default:
			fmt.Printf("\tnone of the previous cases were met\n")
		}
	}

}
