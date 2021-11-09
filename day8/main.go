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

	// strを+で区切ってスライスに変換する
	stringsSlice := strings.Split(str, "+")

	// 数字の配列
	var numbers []int

	// TODO: リファクタリングできそう
	for _, s := range stringsSlice {
		onesPlace := strings.Count(s, "/")
		tensPlace := strings.Count(s, "<")
		number := onesPlace + tensPlace*10

		numbers = append(numbers, number)
	}

	sum := 0
	for _, i := range numbers{
		sum += i
	}

	fmt.Println(sum)
}
