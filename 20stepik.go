package main

import "fmt"

func main() {
	var first, second string
	fmt.Scan(&first, &second)

	for i := 0; i < len(first); i++{
		for j:= 0; j < len(second); j++{
			if first[i] == second[j]{
				fmt.Printf("%c", first[i])
				fmt.Print(" ")
			}
		}
	}
	fmt.Printf("\n")
}