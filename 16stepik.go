package main

import "fmt"

func main() {
	var num int
	max := -1
	volume := 1
	for {
		fmt.Scan(&num)
		if num == 0{
			break
		} else if max < num{
			max = num
			volume = 1
		} else if max == num{
			volume += 1
		}
	}
	fmt.Println(volume)
}