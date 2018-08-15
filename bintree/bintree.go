package bintree

// Node represents a binary tree's node
type Node struct {
	Parent *Node
	Zero   *Node
	One    *Node
	Char   byte
}

var Alphabet [256]byte = fillAlphabet()

// fills the alphabet with all the ascii values
func fillAlphabet() [256]byte {
	var arr [256] byte
	for i := 0; i < 256; i++{
		arr[i] = uint8(i)
	}
	return arr
}

// keeps count of all the characters in the string
func CountChar(str string) [256]int{
	// convert input string to byte array
	bArray := []byte(str)
	var ret [256]int
	
	// zero out all of array
	for i := range ret{
		ret[i] = 0
	}
	
	// keep count
	for i := 0; i < len(bArray); i++{
		ret[bArray[i]] += 1
	}
	return ret
}
