package main

import "fmt"

func

func main() {
	defer fmt.Println("WORLD")

	fmt.Println("HELLO")
}