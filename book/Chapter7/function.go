package main

import "fmt"

/*func average(sx []float64) float64 {
	total := 0.0
	for _, v := range sx {
		total += v
	}
	return total / float64(len(sx))
}

func main() {
	someOtherName := []float64{11, 42, 53, 84, 15}
	// kx := []string{"hello", "world", "good", "morning", "good", "night"}

	fmt.Println((average(someOtherName)))
} */

// not work example
// функции не имеют доступа к области видимости родительской функции, то есть это не сработает
/*func f() {
	fmt.Println(x)
}

func main() {
	x := 1
	f()
} */

// work example
func f(x int) {
	fmt.Println(x)
}

func main() {
	x := 1
	f(x)
}
