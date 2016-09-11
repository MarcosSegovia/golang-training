package main

import "fmt"

func main() {
	var userName string
	fmt.Println("Could you please tell me which is your name?")
	fmt.Scanf("%s", &userName)
	fmt.Println("Hello", userName)
}
