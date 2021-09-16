package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "hello world"
	fmt.Println(s)

	s = strings.Replace(s, "h", "x", 1)
	fmt.Println(strings.Contains(s, "xello"))
	fmt.Println(s)

	fmt.Println(`test
	test`)
	fmt.Println(`"`)

	t, f := true, false
	fmt.Printf("%T, %v\n", t, t)
	fmt.Printf("%T, %v\n", f, f)

	x := 1
	xx := float64(x)

	fmt.Printf("%T, %v, %f\n", xx, xx, xx)

	y := 1.2
	yy := int(y)
	fmt.Printf("%T, %v, %d\n", yy, yy, yy)

	str := "14"
	i, _ := strconv.Atoi(str)

	fmt.Printf("%T, %v, %d\n", i, i, i)
}
