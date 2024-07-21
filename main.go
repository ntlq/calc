package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

func printNumber(v int, flag bool) {
	if flag {
		fmt.Printf("%v\n", arabicToRoman(v))
	} else {
		fmt.Printf("%v\n", v)
	}

}

func arabicToRoman(n int) string {
	// 100 достаточно, потому что по условию не больше 10 * 10
	var arabics = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var result string
	var index = len(arabics) - 1
	for n > 0 {
		for n >= arabics[index] {
			n -= arabics[index]
			result += romans[index]
		}
		index -= 1
	}
	return result
}

func validateNumber(s string) (int, bool) {
	var isRoman = false
	v, ok := romanToInt[s]
	if ok {
		isRoman = true
	} else {
		vInt, err := strconv.Atoi(s)
		if (err != nil) || (vInt < 1) || (vInt > 10) {
			panic("Что-то не так с числом.")
		}
		return vInt, isRoman
	}
	return v, isRoman

}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		operands := strings.Fields(input)
		if len(operands) != 3 {
			panic("Только 2 числа и оператор")
		}
		num1, romanFlag1 := validateNumber(operands[0])
		op := operands[1]
		num2, romanFlag2 := validateNumber(operands[2])
		if romanFlag2 != romanFlag1 {
			panic("Oба римские, либо оба арабские")
		}

		var result int
		switch op {
		case "+":
			result = num1 + num2
			printNumber(result, romanFlag1)
		case "-":
			result = num1 - num2
			if romanFlag1 {
				if result < 1 {
					panic("Получилось что-то не римское")
				} else {
					printNumber(result, romanFlag1)
				}
			} else {
				printNumber(result, romanFlag1)
			}
		case "*":
			result = num1 * num2
			printNumber(result, romanFlag1)
		case "/":
			result = num1 / num2
			if romanFlag1 {
				if result < 1 {
					panic("Получилось что-то не римское")
				} else {
					printNumber(result, romanFlag1)
				}
			} else {
				printNumber(result, romanFlag1)
			}
		default:
			panic("Такого оператора нам не надо.")
		}
		continue
	}
}
