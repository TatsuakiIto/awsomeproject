package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 読み込む値
	//reader := strings.NewReader("9315-35-7398")

	// 本番ではこっちにする
	scanner := bufio.NewScanner(os.Stdin)
	//scanner := bufio.NewScanner(reader)
	scanner.Scan()

	// 電話番号を取得
	phoneNumber := strings.Replace(scanner.Text(), "-", "", -1)
	var sum int
	result := Range(phoneNumber, sum)

	fmt.Println(result)
}

// 配列で何回回るか処理する
func Range(input string, sum int) int{
	for _, n := range input {
		number, err := strconv.Atoi(string(n))
		if err != nil{
			fmt.Println("strconv err: ", err)
		}

		switch {
		case number == 0:
			sum += 12
		case number == 1:
			sum += 3
		case number == 2:
			sum += 4
		case number == 3:
			sum += 5
		case number == 4:
			sum += 6
		case number == 5:
			sum += 7
		case number == 6:
			sum += 8
		case number == 7:
			sum += 9
		case number == 8:
			sum += 10
		case number == 9:
			sum += 11
		}
	}

	return sum * 2
}
