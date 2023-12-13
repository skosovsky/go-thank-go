package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// result представляет результат матча
type result byte

// Возможные результаты матча
const (
	win  result = 'W'
	draw result = 'D'
	loss result = 'L'
)

// team представляет команду
type team byte

// match представляет матч
// содержит три поля:
// - first (первая команда)
// - second (вторая команда)
// - result (результат матча)
// например, строке BAW соответствует
// first=B, second=A, result=W
type match struct {
	first  team
	second team
	result
}

// points возвращает количество очков,
// которые получили за матч первая и вторая команда
func (m *match) points() (int, int) {
	switch m.result {
	case win:
		return 3, 0
	case loss:
		return 0, 3
	default:
		return 1, 1
	}
}

// rating представляет турнирный рейтинг команд -
// количество набранных очков по каждой команде
type rating map[team]int

// tournament представляет турнир
type tournament []match

// calcRating считает и возвращает рейтинг турнира без запроса очков
func (trn *tournament) calcRating() rating {
	rate := rating{}
	for _, el := range *trn {
		switch el.result {
		case win:
			rate[el.first] += 3
			rate[el.second] += 0
		case draw:
			rate[el.first] += 1
			rate[el.second] += 1
		case loss:
			rate[el.first] += 0
			rate[el.second] += 3
		}
	}
	return rate
}

// calcRating запрашивает очки и возвращает рейтинг турнира
func (trn *tournament) calcPointRating() rating {
	rt := rating(make(map[team]int))
	for _, mat := range *trn {
		firstPts, secondPts := mat.points()
		rt[mat.first] += firstPts
		rt[mat.second] += secondPts
	}
	return rt
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// Код, который парсит результаты игр и печатает рейтинг, уже реализован,
// ваша задача - реализовать недостающие структуры и методы выше
func main() {
	src := readString()
	trn := parseTournament(src)
	rt := trn.calcRating()
	rt.print()
}

// readString считывает и возвращает строку из os.Stdin
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}

// parseTournament парсит турнир из исходной строки
func parseTournament(s string) tournament {
	descriptions := strings.Split(s, " ")
	trn := tournament{}
	for _, descr := range descriptions {
		m := parseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

// parseMatch парсит матч из фрагмента исходной строки
func parseMatch(s string) match {
	return match{
		first:  team(s[0]),
		second: team(s[1]),
		result: result(s[2]),
	}
}

// addMatch добавляет матч к турниру
func (t *tournament) addMatch(m match) {
	*t = append(*t, m)
}

// print печатает результаты турнира
// в формате Aw Bx Cy Dz
func (r *rating) print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}
