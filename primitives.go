package main

import "fmt"

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1")
	go 

	fmt.Println("hi")
}
