package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func main() {
	var number int
	_, err := fmt.Scanf("%d", &number)
	if err != nil {
		log.Println(err)
	}

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)

	for _, v := range strconv.Itoa(number) {
		s := string(v)
		n, _ := strconv.Atoi(s)
		counter[n]++
	}

	x := number
	for x != 0 { // Вариант с остатком от деления
		digit := x % 10
		counter[digit]++
		x = x / 10
	}

	for ; number > 0; number /= 10 { // Вариант без лишних переменных
		counter[number%10]++
	}

	for number > 0 { // Аналогично, но более читаемый вариант
		counter[number%10] += 1
		number = number / 10
	}

	for _, character := range strconv.Itoa(number) {
		characterToInt := int(character - '0') // Необычный вариант, с приведением к int
		counter[characterToInt]++
	}

	printCounter(counter)
}

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
