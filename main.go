package main

import (
	"fmt"

	"./huffEncode"
)

func main() {
	var comp = "bibbity_bobbity"

	fmt.Println(comp)
	huffencode.Encode(comp)

}
