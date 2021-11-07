package refact

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "1 80 40"

	// 配列にする
	// TODO: ここも配列にするのが本当に正しいかの検討
	s := strings.Split(str, " ")

	// TODO: ここのリファクタリング
	mInt, _ := strconv.Atoi(s[0])
	pInt, _ := strconv.Atoi(s[1])
	qInt, _ := strconv.Atoi(s[2])

	m := float32(mInt)
	p := MakePercentage(pInt)
	q := MakePercentage(qInt)

	fmt.Println(p, q)

	// 残りを計算する
	u := m - (m * p)
	u = u - (u * q)
	fmt.Println(u)
}

func MakePercentage(i int) float32 {
	castI := float32(i)
	f := castI / 100
	return f
}
