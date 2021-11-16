package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 読み込む値
	reader := strings.NewReader("3\n1 1024\n11 2048\n21 4196")

	// 本番ではこっちにする
	// scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	// レシートの枚数を取得
	receiptCount := scanner.Text()
	receiptCountInt, _ := strconv.Atoi(receiptCount)

	// ポイントを定義
	point := 0

	// レシートの枚数分日付と金額を取得
	for i := 0; i < receiptCountInt; i++ {
		scanner.Scan()
		info := strings.Split(scanner.Text(), " ")

		// fmt.Println(info[0])
		// 3がつく, 5がつく判定
		if strings.Contains(info[0], "3") {
			paid, _ := strconv.Atoi(info[1])
			dayPoint := float64(paid) * 0.03
			point += int(dayPoint)
		} else if strings.Contains(info[0], "5") {
			paid, _ := strconv.Atoi(info[1])
			dayPoint := float64(paid) * 0.05
			point += int(dayPoint)
		} else {
			paid, _ := strconv.Atoi(info[1])
			dayPoint := float64(paid) * 0.01
			point += int(dayPoint)
		}
	}

	fmt.Println(point)
}
