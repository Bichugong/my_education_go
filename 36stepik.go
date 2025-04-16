package main

import (
  "fmt"
  "bufio"
	"os"
	"strings"
)
func Odd_Numbers(s string) string{
	r := []rune(s)
	var r2 []rune
	for i := 1; i < len(r); i += 2{
			r2 = append(r2, r[i])
	}
	return string(r2)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%s\n", Odd_Numbers(text))
}