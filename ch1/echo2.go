package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		s += arg + sep
		sep = " "
	}
	fmt.Println(s)
}