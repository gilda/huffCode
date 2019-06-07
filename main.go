package main

import (
	"fmt"

	"./bintree"
	"./huffEncode"
)

func main() {
	var comp = "bibbity_bobbity"

	fmt.Println(comp)
	fmt.Println(string(bintree.FormatChar(comp)[:]))
	huffencode.WriteEncodingToFile(comp, "")

}
