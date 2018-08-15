package main

import (
	"fmt"

	"./bintree"
)

func main() {
	var comp = "gildaisverynice!!!!"
	
	fmt.Printf("commpressed = %s\nlen = %d\n\n", comp, len(comp))
	fmt.Printf("Alphabet = %s\nlen = %d\n\n", bintree.Alphabet, len(bintree.Alphabet))

	fmt.Printf("%d", bintree.CountChar(comp))

	

}
