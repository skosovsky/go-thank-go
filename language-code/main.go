package main

import (
	"fmt"
	"log"
)

func main() {
	var code string
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Println(err)
	}

	// определите полное название языка по его коду
	// и запишите его в переменную `lang`
	// ...

	var lang string
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}

	fmt.Println(lang)
}
