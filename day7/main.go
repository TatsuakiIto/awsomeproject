package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	str := "777 80 777"
	radius := 78
	s := strings.Split(str, " ")

	boxSides := MakeIntSlice(s)

	// 比較するときに全部回さなくていいようにソートする
	sort.Ints(boxSides)

	if boxSides[0] > radius {
		fmt.Println("入るで！")
	}
}

// 取得した各辺の長さの配列を[]intに変換する処理
func MakeIntSlice(s []string) []int {
	var result []int

	for _, v := range s {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}

	return result
}


