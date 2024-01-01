package main

import "fmt"

func main() {
	const fut = 0.3048
	fmt.Print("Farinheit: ")
	var F float32
	fmt.Scanf("%f", &F)

	output := (F - 32) * 5 / 9
	mypenlenght := 10 * fut

	fmt.Println("Celceus: ", output)
	fmt.Println("Lenght: ", mypenlenght)
}
