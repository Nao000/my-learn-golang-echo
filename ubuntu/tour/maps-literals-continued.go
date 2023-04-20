package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"B": {40, 12},
	"G": {20, 12},
}

func main() {
	fmt.Println(m)
}