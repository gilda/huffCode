package huffencode

import (
	"fmt"
	//    "io"
	"os"

	"../bintree"
)

func writeEncodingToFile(str string, filePath string) {
	fmt.Println()
	os.Create(filePath)

	// create the binary tree and base
	formChar := bintree.FormatChar(str)
	cChar := bintree.CountChar(str)
	var base []*bintree.Node
	// generateteh base of the tree
	for _, v := range formChar {
		base = append(base, &bintree.Node{Char: v, Dist: float32(cChar[v]) / float32(len(str))})
	}
	// get the msater node for encoding
	master := bintree.GenerateTree(str)
	chr, encoding := bintree.GetEncoding(base, master)
}
