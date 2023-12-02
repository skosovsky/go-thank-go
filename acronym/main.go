package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...

	words := strings.Fields(phrase)
	var abbr []rune
	for _, word := range words {
		symbol := []rune(word)[0] // Или так: utf8.DecodeRuneInString([]rune(word)[0])
		if unicode.IsLetter(symbol) {
			abbr = append(abbr, unicode.ToUpper(symbol))
		}
	}
	fmt.Println(string(abbr))

	abbrB := new(strings.Builder) // Вариант со strings.Builder
	prevC := ' '
	for _, c := range phrase {
		if prevC == ' ' && unicode.IsLetter(c) {
			abbrB.WriteRune(unicode.ToUpper(c))
		}
		prevC = c
	}

	fmt.Println(abbrB.String())

	abbrP := ""
	newWordExpected := true
	for _, l := range phrase { // Забавный вариант с перебором всех букв
		if newWordExpected && unicode.IsLetter(l) {
			abbrP += string(unicode.ToUpper(l))
			newWordExpected = false
		} else if string(l) == " " {
			newWordExpected = true
		}
	}

	fmt.Println(abbr)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
