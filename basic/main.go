package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 42
	var j string
	j = strconv.Itoa(i)
	fmt.Println(i, j+"1")
}
