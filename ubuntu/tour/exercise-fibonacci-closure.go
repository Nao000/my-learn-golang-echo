package main

import "fmt"

func fibonacci() func() int {
	count := -1
	v1 := 0
	v2 := 1
	return func() int {
		count++

		if count == 0 || count == 1 {
			return count
		}

		res := v1 + v2
		v1 = v2
		v2 = res

		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}