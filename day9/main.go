package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 当たり番号取得
	scanner.Scan()
	winningNumbers := scanner.Text()
	winnigNumberList := strings.Split(winningNumbers, " ")

	// 枚数取得
	scanner.Scan()
	numberOfSheets := scanner.Text()
	intNumberOfSheets, _ := strconv.Atoi(numberOfSheets)

	// 枚数分当たり番号を取得してスライスに入れる
	var buyNumberList [][]string
	for i := 0; i < intNumberOfSheets; i++ {
		scanner.Scan()
		buyNumbers := scanner.Text()
		buyNumberSlice := strings.Split(buyNumbers, " ")
		buyNumberList = append(buyNumberList, buyNumberSlice)
	}

	for _, bn := range buyNumberList {
		winningCount := 0
		for _, b := range bn {
			for _, w := range winnigNumberList {
				if b == w {
					winningCount += 1
				}
			}
		}
		fmt.Println(winningCount)
	}
}
