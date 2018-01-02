package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		name := basename(arg)
		fmt.Printf("path:%s\tbasename:%s\n", arg, name)
	}
}

func basename(s string) string {
	if slash := strings.LastIndex(s, "/"); slash >= 0 {
		s = s[slash + 1:]
	}
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}