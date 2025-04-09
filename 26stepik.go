package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	array := make([]int, num, num*2)

	for i := 0; i < num; i++{
		fmt.Scan(&array[i])
	}
	fmt.Println(array[3])
}