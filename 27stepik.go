package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	array := make([]int, num, num)

	for i := 0; i < num; i ++{
		fmt.Scan(&array[i])
	}
	for i:= 0; i < len(array); i += 2{
		fmt.Printf("%d ", array[i])
	}
}