package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	textByte := []byte(text) // Или Rune
	dots := []byte("...")
	res := text

	if len(textByte) > width {
		resByte := textByte[:width]
		resByte = append(resByte, dots...)
		res = string(resByte)
	}

	res = text // Простой вариант, только текст
	if len(text) > width {
		res = text[:width] + "..."
	}

	if len(text) > width { // Вариант через Sprintf
		res = fmt.Sprintf("%v...", text[:width])
	}

	fmt.Println(res)
}
