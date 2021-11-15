package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 単語ごとに分けて読み込むデフォルトであるSplit
	reader := strings.NewReader("abc efg\nhijk lmn")
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// 数値を読み込む
	scanner2 := bufio.NewScanner(strings.NewReader("1234 5678 1234567891234567890"))

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner2.Split(split)

	for scanner2.Scan() {
		fmt.Println(scanner2.Text())
	}

	if err := scanner2.Err(); err != nil {
		fmt.Println(err)
	}

	// カンマ区切りの単語を読み込む
	scanner3 := bufio.NewScanner(strings.NewReader("1,2,3,4"))
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner3.Split(onComma)
	for scanner3.Scan() {
		fmt.Println(scanner3.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
