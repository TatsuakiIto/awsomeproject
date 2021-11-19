package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type CenterCoordinate struct {
	x int
	y int
}

type PersonCoordinate struct {
	x int
	y int
}

func main() {
	// 読み込む値
	reader := strings.NewReader("0 0 1 2\n3\n0 0 \n1 1 \n4 2")

	// 本番ではこっちにする
	// scanner := bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()

	// 座標を取得
	coordinate := scanner.Text()
	coordinateSlice := strings.Split(coordinate, " ")

	centerCoordinate := getCenterCoordinate(coordinateSlice[0], coordinateSlice[1])
	minRadius, _ := strconv.Atoi(coordinateSlice[2])
	maxRadius, _ := strconv.Atoi(coordinateSlice[3])

	minCoordinate := minRadius * minRadius
	maxCoordinate := maxRadius * maxRadius

	// 人数の取得
	scanner.Scan()
	people, _ := strconv.Atoi(scanner.Text())

	fmt.Println(people)

	// 人数分回す
	for i := 0; i < people; i++ {
		scanner.Scan()
		personCoordinateSlice := strings.Split(scanner.Text(), " ")
		personCoordinateStruct := getPersonCoordinate(personCoordinateSlice[0], personCoordinateSlice[1])

		personCoordinate := (centerCoordinate.x-personCoordinateStruct.x)*(centerCoordinate.x-personCoordinateStruct.x) + (centerCoordinate.y-personCoordinateStruct.y)*(centerCoordinate.y-personCoordinateStruct.y)

		if minCoordinate <= personCoordinate && personCoordinate <= maxCoordinate {
			fmt.Println("yes")
		} else {
			fmt.Println("No")
		}
	}
}

// 中心座標の取得
func getCenterCoordinate(x, y string) CenterCoordinate {
	xc, _ := strconv.Atoi(x)
	yc, _ := strconv.Atoi(y)
	return CenterCoordinate{x: xc, y: yc}
}

// 個人の座標取得
func getPersonCoordinate(x, y string) PersonCoordinate {
	xc, _ := strconv.Atoi(x)
	yc, _ := strconv.Atoi(y)
	return PersonCoordinate{x: xc, y: yc}
}
