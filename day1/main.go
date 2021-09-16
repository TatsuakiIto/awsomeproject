// 問題: https://i.gyazo.com/4e271600934bac2eaeee45d4c8fbc034.png

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	str = strings.TrimSpace(str)

	Split(str)
}

func Split(input string) {
	slice := strings.Split(input, "")
	len := len(slice)
	for i := 0; i < len; i++ {
		fmt.Printf("%s\n", slice[i])
	}
}
