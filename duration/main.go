package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// считываем временной отрезок из os.Stdin
	// гарантируется, что значение корректное
	// не меняйте этот блок
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Println(err)
	}
	d, err := time.ParseDuration(s)
	if err != nil {
		log.Println(err)
	}

	// выведите исходное значение
	// и количество минут в нем
	// в формате "исходное = X min"
	// используйте метод .Minutes() объекта d
	fmt.Println(s, "=", d.Minutes(), "min")
	fmt.Printf("%s = %v min", s, d.Minutes())
	fmt.Printf("%s = %.0f min\n", s, d.Minutes())
}
