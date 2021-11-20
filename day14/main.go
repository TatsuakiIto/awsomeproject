package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type ConstructionSite struct{
	x int // 工事現場のx座標
	y int // 工事現場のy座標
	r int // 工事現場の影響半径
}

type Spot struct{
	x int
	y int
}

func main() {
	// 読み込む値
	reader := strings.NewReader("20 10 10\n3\n25 10\n20 15\n70 70")

	// 本番ではこっちにする
	// scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	// 座標を取得
	constructionSiteLine := scanner.Text()
	constructionSiteSlice := strings.Split(constructionSiteLine, " ")
	constructionSite := changeLineToConstructionSiteStruct(constructionSiteSlice)

	// 場所の数の取得
	scanner.Scan()
	spots, _ := strconv.Atoi(scanner.Text())

	// 場所の数分回す
	for i := 0; i < spots; i++ {
		scanner.Scan()
		spotLine := strings.Split(scanner.Text(), " ")
		spot := getSpot(spotLine)

		judgeSpot(constructionSite, spot)
	}
}

// 工事現場座標のstringsを構造体に変換
func changeLineToConstructionSiteStruct(s []string) ConstructionSite {
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	r, _ := strconv.Atoi(s[2])
	return ConstructionSite{x: x, y: y, r: r}
}

// 場所の座標取得
func getSpot(s []string) Spot {
	x, _ := strconv.Atoi(s[0])
	y, _ := strconv.Atoi(s[1])
	return Spot{x: x, y: y}
}

// 判定処理
func judgeSpot(c ConstructionSite, s Spot) {
	if (s.x - c.x) * (s.x - c.x) + (s.y - c.y) * (s.y - c.y) >= c.r * c.r{
		fmt.Println("silent")
	} else {
		fmt.Println("noisy")
	}
}
