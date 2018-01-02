package main
import (
	"fmt"
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
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:];
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}