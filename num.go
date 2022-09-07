package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2, s1, s2 int
	var count1 []int
	var count2 []int

	fmt.Scan(&num1, &num2)

	s1 = NumberOfDigits(num1)
	s2 = NumberOfDigits(num2)

	for i := int(math.Pow(10, float64(s1-1))); i > 0; i = i / 10 {
		count1 = append(count1, num1/i%10)
	}

	for i := int(math.Pow(10, float64(s2-1))); i > 0; i = i / 10 {

		count2 = append(count2, num2/i%10)
	}

	fmt.Print(count1, count2)
}

func NumberOfDigits(num int) int {
	s:=0
	for num > 0 {
		num = num / 10
		s++
	}
	return s
}
