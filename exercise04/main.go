package main

import (
	"fmt"
)

func main() {
	var largeNumber int
	var shortNumber int
	fmt.Println("Please enter a large number :")
	fmt.Scanln(&largeNumber)
	fmt.Println("Please enter a short number :")
	fmt.Scanln(&shortNumber)
	remainder := remainder(largeNumber, shortNumber)
	fmt.Println("The remainder between those is:", remainder)

}

func remainder(larger int, short int) int {
	return larger % short
}
