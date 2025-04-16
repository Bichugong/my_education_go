package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	str := []rune(text)

	result := make([]rune, 0)

	for i, char := range str {
		flag := true
		for j, otherChar := range str {
			if char == otherChar && i != j {
				flag = false
				break
			}
		}
		if flag {
			result = append(result, char)
		} 
	}
	fmt.Printf("%s\n", string(result))
}