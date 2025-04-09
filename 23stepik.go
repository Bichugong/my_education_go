package main

import "fmt"

func main(){
	a := [5]int{5, 10, 15, 20, 25}

	for idx := range a {
  	  fmt.Println(a[idx]) // Вывод значений массива через индекс
	}

	for _, elem := range a {
  	  fmt.Println(elem) // Вывод значений напрямую
	}
}