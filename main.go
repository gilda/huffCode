package main

import (
	"fmt"

	"./bintree"
)

func main() {
	var comp = "gggggiiildda!!!!"

	fmt.Println(comp)
	fmt.Println(string(bintree.FormatChar(comp)[:]))
	fmt.Println(bintree.GenerateTree(comp))
}
