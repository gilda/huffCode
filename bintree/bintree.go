package bintree

import (
	"fmt"
)

// Node represents a binary tree's node
type Node struct {
	Parent *Node
	Zero   *Node
	One    *Node
	Char   byte
	Dist   float32
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
	var ret []byte = nil
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

// GenerateTree generates the binary tree according to the formatted array
func GenerateTree(str string) *Node {
	// get the distribution of chars in the string
	formChar := FormatChar(str)
	cChar := CountChar(str)

	// generate all character nodes with their distributaions
	var nodes []Node
	for _, v := range formChar {
		nodes = append(nodes, Node{Char: v, Dist: float32(cChar[v]) / float32(len(str))})
	}
	addNode(nodes)
	fmt.Println(nodes)

	return nil
}

// addTreeLayer adds a layer of aprent nodes with the correct encoding
func addNode(base []Node) {
	// finds the two parent less nodes with lowest dist
	lowest, secLowest := findLowestDist(base)
	// create new node with connections to them
	added := Node{Zero: secLowest, One: lowest, Dist: lowest.Dist + secLowest.Dist}

	// assign parents
	lowest.Parent = &added
	secLowest.Parent = &added
}

// TODO figure out why parents are not added probably pointer issue
// findLowestDist finds the node with the lowest dist in the tree
func findLowestDist(base []Node) (*Node, *Node) {
	var noParent []Node
	var s Node

	// get all nodes with no parent
	for _, v := range base {
		if v.Parent == nil {
			noParent = append(noParent, v)
			continue
		}
		for v.Parent != nil {
			s = *v.Parent
		}
		// append new node
		noParent = append(noParent, s)
	}

	// this is the last node, we finished building binary tree
	if len(noParent) == 1 {
		return &noParent[0], nil
	}

	var lowest Node
	var secLowest Node
	// find two parentless nodes with lowest distribution
	for i, v := range noParent {
		if i == 0 {
			lowest = v
		} else if i == 1 {
			secLowest = v
			// find the lowest distribution
		} else if v.Dist < lowest.Dist || v.Dist < secLowest.Dist {
			if v.Dist < lowest.Dist {
				lowest = v
			} else {
				secLowest = v
			}
		}
	}

	return &lowest, &secLowest
}
