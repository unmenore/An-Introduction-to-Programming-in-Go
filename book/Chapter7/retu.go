package main

import "fmt"

func f() (int, string) {
	return 4, "example"
}

func main() {
	x, y := f()
	fmt.Println(x, y)
}
