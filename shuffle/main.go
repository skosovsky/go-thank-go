package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// shuffle перемешивает элементы nums in-place.
func shuffle(nums []int, rnd *rand.Rand) {
	rnd.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

func main() {
	rnd := rand.New(rand.NewSource(42))

	nums := readInput()
	shuffle(nums, rnd)
	fmt.Println(nums)
}

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
