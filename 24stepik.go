package main

import "fmt"

func main(){
	a := [5]int{5, 10, 15, 20, 25}
	fmt.Println(a) // [5 10 15 20 25]

	for _, elem := range a {
  	  elem = 100
    	fmt.Println(elem)
    	// 100
    	// 100
    	// 100
    	// 100
    	// 100
	}
	fmt.Println(a) // [5 10 15 20 25] - массив не изменился!

	for idx := range a {
  	  a[idx] = 100
    	fmt.Println(a[idx])
    	// 100
    	// 100
    	// 100
    	// 100
    	// 100
	}
	fmt.Println(a) // [100 100 100 100 100] - теперь массив изменился
}