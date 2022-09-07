package main

import "fmt"

func main() {
	var num int
	var count []int
	fmt.Scan(&num)
	for i := 10000; i > 0; i = i / 10 {
		count = append(count, num/i%10)
	}
}
