package main

import "fmt"

type Mytype struct {
	id int
	name string
}

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	} {
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	t := []Mytype{
		{1, "name-0"},
		{2, "name-1"},
		{3, "name-2"},
		{4, "name-3"},
	}
	fmt.Println(t)



}