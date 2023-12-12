package main

import (
	"fmt"
	"strings"
)

func main() {
	roman := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	numeral := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	var a, s, b string
	fmt.Scanln(&a, &s, &b)

	a = strings.ToUpper(a)
	b = strings.ToUpper(b)

	if contains(roman, a) && contains(roman, b) {
		switch s {
		case "+":
			otvet := indexOf(roman, a) + indexOf(roman, b) + 1
			fmt.Println(roman[otvet])
		case "-":
			otvet := indexOf(roman, a) - (indexOf(roman, b) + 1)
			if otvet == 0 || otvet > 0 {
				fmt.Println(roman[otvet])
			} else {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			}
		case "*":
			otvet := (indexOf(roman, a)+1)*(indexOf(roman, b)+1) - 1
			fmt.Println(roman[otvet])
		case "/":
			otvet := (indexOf(roman, a)+1)/(indexOf(roman, b)+1) - 1
			fmt.Println(roman[otvet])
		}
	} else if contains(numeral, a) && contains(numeral, b) {
		switch s {
		case "+":
			otvet := indexOf(numeral, a) + indexOf(numeral, b)
			fmt.Println(otvet)
		case "-":
			otvet := indexOf(numeral, a) - indexOf(numeral, b)
			fmt.Println(otvet)
		case "*":
			otvet := indexOf(numeral, a) * indexOf(numeral, b)
			fmt.Println(otvet)
		case "/":
			otvet := indexOf(numeral, a) / indexOf(numeral, b)
			fmt.Println(otvet)
		}
	} else if contains(numeral, a) && contains(roman, b) {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else if contains(numeral, b) && contains(roman, a) {
		fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
	} else {
		var z string
		fmt.Scan(&z)
		fmt.Print("Вывод ошибки, так как строка не является математической операцией.")
	}
}

func indexOf(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}

func contains(slice []string, val string) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}
