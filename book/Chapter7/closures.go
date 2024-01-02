package main

import "fmt"

// func main() {
// 	add := func(x, y int) int {
// 		return x + y
// 	}

// 	fmt.Println(add(1, 2))
// }

/*add является локальной переменной типа func(int, int) int (функция принимает два аргумента типа int и возвращает int).
При создании локальная функция также получает доступ к локальным переменным */

// func main() {
// 	x := 0
// 	increment := func() int {
// 		x++
// 		return x
// 	}

// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

/* increment прибавляет 1 к переменной x, которая определена в рамках функции main.
Значение переменной x может быть изменено в функции increment.
Вот почему при первом вызове increment на экран выводится 1, а при втором — 2.

Функцию, использующую переменные, определенные вне этой функции, называют замыканием.
В нашем случае функция increment и переменная x образуют замыкание.
*/

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
