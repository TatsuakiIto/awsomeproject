package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	str_count := utf8.RuneCountInString(str)
	fmt.Println(DecorateSurroundings(str_count))
	fmt.Println(DecorateString(str))
	fmt.Println(DecorateSurroundings(str_count))
}

func DecorateSurroundings(c int) string {
	var bar string

	for i := 0; i < c+2; i++ {
		bar += "+"
	}
	return bar
}

func DecorateString(s string) string {
	str := "+" + s + "+"
	return str
}
