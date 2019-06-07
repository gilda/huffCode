package bintree

import "fmt"

// Node represents a binary tree's node
type Node struct {
	Parent *Node
	Zero   *Node
	One    *Node
	Char   byte
	Dist   float32
}

func (n *Node) String() string {
	return fmt.Sprintf("%p -> {%p %p %p %d %f}", n, n.Parent, n.Zero, n.One, n.Char, n.Dist)
}

var alphabet = fillAlphabet()

// fills the alphabet with all the ascii values
func fillAlphabet() [256]byte {
	var arr [256]byte
	for i := 0; i < 256; i++ {
		arr[i] = uint8(i)
	}
	return arr
}

// CountChar keeps count of all the characters in the string
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
		ret[bArray[i]]++
	}
	return ret
}

// check wether or not the a array contains the s byte
func containsByte(a []byte, s byte) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

// UniqueChars returns the number of unique characters
func UniqueChars(str string) []byte {
	var found []byte

	//first character is always unique
	found = append(found, str[0])

	// loop over all characters and find uniques
	for i := 1; i < len(str); i++ {
		if !containsByte(found, str[i]) {
			// append if it wasnt found yet
			found = append(found, str[i])
		}
	}
	// return all the unique characters and the length
	return found
}

// FormatChar sort and cut the CountChar output to start creating the binary tree
func FormatChar(str string) []byte {
	var ret []byte
	cChar := CountChar(str)
	notSorted := UniqueChars(str)

	// assign correct size for the returning slice
	for range notSorted {
		ret = append(ret, byte(0))
	}

	// sort array by the number of all occurences in str
	for i := range ret {
		// search for the largest number of occurences
		var largest int
		var char byte
		for _, v := range notSorted {
			// if found largest note
			if cChar[v] > largest {
				largest = cChar[v]
				char = v
			}
		}
		// make sure not passing by again on the largest
		cChar[char] = -1
		// append to the returned array
		ret[i] = char
	}

	// return sorted array
	return ret
}

// GenTree generate the tree and return the master node
func GenTree(base []*Node) *Node {
	// endless loop to get the final tree
	for {
		// if only one node has no parents then the tree is done
		// return the master node
		if len(findNoParents(base)) == 1 {
			return findNoParents(base)[0]
		}

		// find the lowest distribution nodes and combine them together
		var lowest, secLowest = findLowestDist(base)

		// create a parent node
		var par = Node{nil, lowest, secLowest, 0, secLowest.Dist + lowest.Dist}

		// assign the two nodes their parent
		lowest.Parent = &par
		secLowest.Parent = &par
	}
}

func findNoParents(base []*Node) []*Node {
	var it = &Node{nil, nil, nil, 0, 0}
	var res []*Node
	var ret []*Node

	// iterate over all of base
	for _, v := range base {
		// if no parent add to array
		if v.Parent == nil {
			res = append(res, v)
		} else {
			// if has parent go up the chain till you find one that doesnt have a parent
			it = v.Parent
			for it.Parent != nil {
				it = it.Parent
			}
			// when found the one that did not have a parent add it
			res = append(res, it)
		}
	}

	// remove duplicates
	for _, v := range res {
		var found = false
		for _, l := range ret {
			if v == l {
				found = true
			}
		}
		if !found {
			ret = append(ret, v)
		}
	}

	// return without duplicates
	return ret
}

func findLowestDist(base []*Node) (*Node, *Node) {
	var noParent = findNoParents(base)

	// find two lowest nodes
	var lowest float32 = 1
	var lowestNode *Node
	var secLowest float32 = 1
	var secLowestNode *Node

	// find absolute lowest node
	for _, v := range noParent {
		if v.Dist < lowest {
			lowest = v.Dist
			lowestNode = v
		}
	}

	// find second lowest node
	for _, v := range noParent {
		if v.Dist < secLowest && v != lowestNode {
			secLowest = v.Dist
			secLowestNode = v
		}
	}

	// return the two nodes with the least distribution
	return lowestNode, secLowestNode
}
