package main

import (
	"fmt"
)

type Path struct {
	x int
	y int
}

func (path *Path) Get(field string) int {
	if field == "x" {
		return path.x
	}
	return path.y
}

func main() {
	temp := Path{1, 2}
	val := temp.Get("x")
	fmt.Println(val)
}