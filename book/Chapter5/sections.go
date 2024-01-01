package main

import "fmt"

// func main() {
// 	slice1 := []int{1, 2, 3}
// 	//slice2 := append(slice1, 4, 5)
// 	slice2 := make([]int, 2)
// 	copy(slice2, slice1)

// 	fmt.Println(slice1, slice2)
// }

func main() {
	n := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := n[3:9]
	fmt.Println(x)
}
