package main

import (
	"math"
	"strconv"
)

func Index2Name(quotient int) string {
	var name []rune
	for quotient > 0 {
		remainder := (quotient - 1) % 26
		name = append([]rune{rune(int('A') + remainder)}, name...)
		quotient = (quotient - remainder) / 26

	}
	return string(name)
}

func Name2Index(name string) (result int) {
	len := len(name)
	for i, c := range []rune(name) {
		result += ((int(c) - int('A')) + 1) * int(math.Pow(26, float64(len-1-i)))
	}
	return
}

func NameSplit(name string) (column int, row int, err error) {
	for i, c := range []rune(name) {
		if int(c) <= int('9') && int(c) >= int('0') {
			column = Name2Index(string(name[:i]))
			row, err = strconv.Atoi(string(name[i:]))
			return
		}
	}
	return
}
