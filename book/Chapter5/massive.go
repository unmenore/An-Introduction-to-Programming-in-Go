package main

import "fmt"

// func main() {
// 	var x [5]int
// 	x[4] = 100
// 	fmt.Println(x[4])
// }

// func main() {
// 	var x = [5]float64{98, 93, 77, 82, 83}

// 	var total float64 = 0
// 	for i := 0; i < 5; i++ {
// 		total += x[i]
// 	}
// 	fmt.Println(total / 5)
// }

// func main() {
// 	var x = [6]string{"hello", "world", "good", "morning", "good", "night"}

// 	fmt.Println(x[2:5])
// }

func main() {
	var x = []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	//minimal element i in x
	min := x[0]
	for _, num := range x {
		if num < min {
			min = num
		}
	}

	fmt.Println(min)
}
