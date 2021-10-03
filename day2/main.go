package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	fmt.Println(Converter(str))
}

func Converter(str string) int {
	i, _ := strconv.Atoi(str)
	i--
	return i
}
