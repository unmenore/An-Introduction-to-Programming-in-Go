package main

import "fmt"

var x = []int{1, 2, 3}

func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return x[2]
}

/*Каждая вызываемая функция помещается
в стек вызовов, каждый возврат из функции возвращает нас к
предыдущей приостановленной подпрограмме */
