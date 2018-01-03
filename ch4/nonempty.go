package main

import "fmt"

func main() {
	a := []string{"one", "", "two", ""}
	fmt.Println(a)
	b := nonempty(a)
	fmt.Println(b)
}

func nonempty(strings []string) []string {
    i := 0
    for _, s := range strings {
        if s != "" {
            strings[i] = s
            i++
        }
    }
    return strings[:i]
}