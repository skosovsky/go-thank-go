package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	var result []int
	for _, el := range iterable {
		if predicate(el) {
			result = append(result, el)
		}
	}
	return result
}

func filterGeneric[T any](predicate func(T) bool, iterable []T) []T { // Вариант с дженериками
	var filtered []T
	for _, item := range iterable {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func main() {
	src := readInput()
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := filter(...)

	res := filter(func(el int) bool { return el%2 == 0 }, src)
	fmt.Println(res)

	f := func(n int) bool { // Предикт наружний, и нечетность через биты
		return n&1 == 0
	}

	res2 := filter(f, src)
	fmt.Println(res2)

	// Заодно проверил разницу в скорости обработки при делении и битовой операции
	fmt.Println(isEvenBit(5), isEvenDiv(5))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}

func isEvenDiv(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func isEvenBit(num int) bool {
	if num&1 == 0 {
		return true
	}
	return false
}
