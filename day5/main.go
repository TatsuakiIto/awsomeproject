package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	
	arr := [] string{"a", "i", "u", "e", "o"}
	for i := 0; i < len(arr); i++ {
	    str = strings.Replace(str, arr[i], "", -1)
	    str = strings.Replace(str, strings.ToUpper(arr[i]), "", -1)
	}
	
	fmt.Println(str)
}
