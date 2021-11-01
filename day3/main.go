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

	fmt.Println(CentitoMili(str))
}

func CentitoMili(str string) int {
	centi, _ := strconv.Atoi(str)
	mili := centi * 10
	return mili
}
