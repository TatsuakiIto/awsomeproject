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

		// 金額を数値にしておく
		paid, _ := strconv.Atoi(info[1])

		point = AddingToPoint(paid, info[1], point)
	}
	fmt.Println(point)

	// 最初の行以外が決まっているのであれば、最初の行はどうする
	// 次の行からはどうするみたいな処理をまとめて書けるようになりたい
}

// 日付で条件を分けてポイントを算出する
func AddingToPoint(paid int, dayStr string, point int) int {
	switch  {
		case strings.Contains(dayStr, "3"):
			point += int(float64(paid) * 0.03)
		case strings.Contains(dayStr, "5"):
			point += int(float64(paid) * 0.05)
		default:
			point += int(float64(paid) * 0.01)
	}
	return point
}
