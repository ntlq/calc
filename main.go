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

func arabicToRoman(n int) string {
	// 100 достаточно, потому что по условию не больше 10
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
func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, _ := strconv.Atoi(input)
		fmt.Println(arabicToRoman(num))
		continue
	}
}
