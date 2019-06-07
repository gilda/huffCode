package huffencode

import (
	"fmt"

	"../bintree"
)

func traverse(master *bintree.Node, base *bintree.Node) string {
	var chain []*bintree.Node
	chain = append(chain, base)
	var it *bintree.Node = base
	var ret = ""

	for it.Parent != nil {
		chain = append(chain, it.Parent)
		it = it.Parent
	}

	for i := len(chain); i > 1; i-- {
		if chain[i-2] == chain[i-1].One {
			ret += "1"
		}
		if chain[i-2] == chain[i-1].Zero {
			ret += "0"
		}
	}

	return ret
}

func printCodes(master *bintree.Node, base []*bintree.Node) {
	for _, v := range base {
		fmt.Print(string(v.Char) + " " + traverse(master, v))
		fmt.Println()
	}
}

// WriteEncodingToFile writes the encoding to a file
func WriteEncodingToFile(str string, filePath string) {
	fmt.Println()
	//os.Create(filePath)

	// create the binary tree and base
	formChar := bintree.FormatChar(str)
	cChar := bintree.CountChar(str)
	var base []*bintree.Node
	// generateteh base of the tree
	for _, v := range formChar {
		base = append(base, &bintree.Node{Char: v, Dist: float32(cChar[v]) / float32(len(str))})
	}

	for _, v := range base {
		fmt.Println(v)
	}
	fmt.Println(bintree.GenTree(base))
	printCodes(bintree.GenTree(base), base)
}
