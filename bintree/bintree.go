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
	var arr [256]byte
	for i := 0; i < 256; i++ {
		arr[i] = uint8(i)
	}
	return arr
}

// keeps count of all the characters in the string
func CountChar(str string) [256]int {
	// convert input string to byte array
	bArray := []byte(str)
	var ret [256]int

	// zero out all of array
	for i := range ret {
		ret[i] = 0
	}

	// keep count
	for i := 0; i < len(bArray); i++ {
		ret[bArray[i]] += 1
	}
	return ret
}

// check wether or not the a array contains the s byte
func ContainsByte(a []byte, s byte) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// returns the number of unique characters
func UniqueChars(str string) []byte {
	var found []byte

	//first character is always unique
	found = append(found, str[0])

	// loop over all characters and find uniques
	for i := 1; i < len(str); i++ {
		if !ContainsByte(found, str[i]) {
			// append if it wasnt found yet
			found = append(found, str[i])
		}
	}
	// return all the unique characters and the length
	return found
}

// TODO
// sort and cut the CountChar output to start creating the binary tree
func FormatChar(str string) []byte {
	return nil
}
