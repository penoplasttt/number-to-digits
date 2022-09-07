package main

import (
	"fmt"
	"math"
)

func main() {
	var num1, num2 int
	
	fmt.Scan(&num1, &num2)

	s1 := NumberOfDigits(num1)
	s2 := NumberOfDigits(num2)

	count1 := ConvertToSlice(num1, s1)
	count2 := ConvertToSlice(num2, s2)
	//fmt.Print(count1, count2)
	Search(count1, count2)
}

func NumberOfDigits(num int) int {
	s:=0
	for num > 0 {
		num = num / 10
		s++
	}
	return s
}

func ConvertToSlice(num, s int)[]int{
	var count []int
	for i := int(math.Pow(10, float64(s-1))); i > 0; i = i / 10 {
		count = append(count, num/i%10)
	}
	return count
}

func Search(count1, count2 []int){
	for _,i:= range count1{
		for _,k:= range count2{
			if i == k{
				fmt.Printf("%d ",i)
			}
		}
	}
}
