package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labls": Vertex{
		40, -74,
	},
	"Google": Vertex{
		30, -120,
	},
}

func main() {
	fmt.Println(m)
}