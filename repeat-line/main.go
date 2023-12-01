package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string
	var times int
	// гарантируется, что значения корректные
	fmt.Scan(&source, &times)

	// возьмите строку `source` и повторите ее `times` раз
	// запишите результат в `result`
	// ...

	var result string
	for i := 1; i <= times; i++ { // или так "for ; times > 0; times-- {"
		result += source
	}

	fmt.Println(result)
	fmt.Println(strings.Repeat(source, times)) // Повторяем через strings.Repeat

	var res strings.Builder
	for i := 0; i < times; i++ {
		res.WriteString(source)
	}
	fmt.Println(res.String()) // Альтернатива через strings.Builder

	for times > 0 { // Необычная реализация for
		result += source
		times--
	}
	fmt.Println(result)
}
