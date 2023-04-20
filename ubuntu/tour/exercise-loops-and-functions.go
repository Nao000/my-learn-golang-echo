package main
import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	count := 0
	for {
		count++
		// fmt.Printf("count: %v\n", count)

		tmp := z
		z -= (z*z - x) / (2*z)

		diff := math.Abs(tmp - z)

		if diff < 0.000000001 {
			break
		}
	}

	if count > 10 {

		fmt.Println("more than 10 calc")
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2000))
}