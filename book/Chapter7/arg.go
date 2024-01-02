package main

import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	// sx := []int{1, 2, 3}
	// fmt.Println((add(sx)))
	fmt.Println((add(2, 3, 4, 5)))
}

// func Println(a ...interface{}) (n int, err error)
