package main

import (
	"fmt"
	"strings"
)

func Repeat(s string, count int) string {
	return strings.Repeat(s, 4)
}

func main() {
	fmt.Println(Repeat("a", 4))
}
