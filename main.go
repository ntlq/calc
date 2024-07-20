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

func printer(value int, flag bool) {
	if flag {
		fmt.Printf("%v\n", arabicToRoman(value))
	} else {
		fmt.Printf("%v\n", value)
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
			//fmt.Println(n)
			result += romans[index]
			//index -= 1
		}
		index -= 1
	}
	return result
}

func numberValidator(s string) (int, bool) {
	var isRoman = false
	v, ok := romanToInt[s]
	if ok {
		isRoman = true
	} else {
		vNum, err := strconv.Atoi(s)
		if (err != nil) || (vNum < 1) || (vNum > 10) {
			panic("Что-то не так с числом.")
		}
		return vNum, isRoman
	}
	return v, isRoman

}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		//input = strings.TrimSpace(input)
		operands := strings.Fields(input)
		if len(operands) != 3 {
			panic("Только 2 числа и оператор")
		}
		num1, romanFlag1 := numberValidator(operands[0])
		op := operands[1]
		num2, romanFlag2 := numberValidator(operands[2])
		if romanFlag2 != romanFlag1 {
			panic("Oба римские, либо оба арабские")
		}

		var result int
		switch op {
		case "+":
			result = num1 + num2
			printer(result, romanFlag1)
		case "-":
			result = num1 - num2
			if romanFlag1 {
				if result < 1 {
					panic("Получилось что-то не римское")
				} else {
					printer(result, romanFlag1)
				}
			}
		case "*":
			result = num1 * num2
			printer(result, romanFlag1)
		case "/":
			result = num1 / num2
			printer(result, romanFlag1)
		default:
			panic("Такого оператора нам не надо.")
		}
		continue
	}
}
