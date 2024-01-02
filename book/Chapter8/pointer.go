package main

import "fmt"

//Когда мы вызываем функцию с аргументами, аргументы копируются в функцию:
/*
func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)
}
*/

/*
func zero(xPtr *int) {
	*xPtr = 1
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
*/

/*Указатели указываютна участок в памяти, где хранится значение.
Используя указатель (*int) в функции zero,мы можем изменить значение оригинальной переменной
*/

//Оператор new

func one(xPtr *int) {
	*xPtr = 55
}
func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
