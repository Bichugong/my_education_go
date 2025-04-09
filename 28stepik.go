package main

import "fmt"

func main() {
  var num int
  k := 0
  fmt.Scan(&num)
  array := make([]int, num, num)
    
  for i := 0; i < num; i++{
    fmt.Scan(&array[i])
   if array[i] > 0{
  	 k += 1
  	}
	}
 fmt.Println(k)
}