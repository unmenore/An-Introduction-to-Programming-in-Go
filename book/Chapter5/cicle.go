package main

import "fmt"

// func main() {
// 	i := 1
// 	for i < 10 {
// 		fmt.Println(i)
// 		i += 1
// 	}
// }

/* new example
func main() {
	for i := 10; i > 1; i-- {
		fmt.Println(i)
	}
}
*/

// func main() {
// 	i :=
// 	if i%2 == 0 {
// 		fmt.Println("Even")
// 	} else {
// 		fmt.Println("Odd")
// 	}
// }

// func main() {
// 	i := 15
// 	if i%2 == 0 {
// 		fmt.Println("Even")
// 	} else if i%3 == 0 {
// 		fmt.Println("Odd")
// 	} else if i%5 == 0 {
// 		fmt.Println("Odd")
// 	} else {
// 		fmt.Println("Odd")

// 	}
// }

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "Buzz")
		}
	}
}
