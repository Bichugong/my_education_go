package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	str := []rune(text)


	wrong := true
	for _, char := range str {
		if !unicode.Is(unicode.Latin, char) && !unicode.IsDigit(char) {
			wrong = false
			break
		}
	}
	if wrong && len(str) > 4 {
		fmt.Println("Ok")
	}else {
		fmt.Println("Wrong password")
	}
}
