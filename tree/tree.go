package tree

import (
	"errors"
	"fmt"
)

/* This node contains a char that is a character stored,
End is to denote the end of a word in the tree. */
type Node struct {
	char rune
	end  bool
	next []*Node
}

func newNode(c rune, e bool) *Node {
	node := &Node{char: c, end: e}
	node.next = make([]*Node, 0)
	return node
}

func (n Node) String() string {
	s := fmt.Sprintln("Char:", string(n.char), " End:", n.end)
	for _, n := range n.next {
		s += fmt.Sprint(n)
	}
	return s
}

/* Tree struct to build char tree for the words beginning with a
particular char */
type Tree struct {
	char  rune
	start *Node
}

/* Adds the word to the tree. Adds a char to the node and then the next char
to a new node that is pointed by the "next" of previous node. The node for
the last char in the word is marked as "end" */
func (t *Tree) AddWord(word string) error {
	if len(word) == 0 {
		return errors.New("empty word")
	} else if rune(word[0]) != t.char {
		return errors.New("incorrect tree")
	}
	node := t.start
	lastCharIndex := len(word) - 1
	for i, c := range word {
		end := false
		if i == lastCharIndex {
			end = true
		}
		if node == nil {
			node = newNode(c, end)
			t.start = node
		}
		// Do not go further if we are in the last char
		if i == lastCharIndex {
			break
		}
		nextEnd := false
		nextChar := rune(word[i+1])
		if i+1 == lastCharIndex {
			nextEnd = true
		}
		// Check the next node for the char
		nextFound := false
		for _, n := range node.next {
			if n.char == nextChar {
				nextFound = true
				node = n
				if nextEnd {
					n.end = true
				}
				break
			}
		}
		// Create a node if nothing matches
		if !nextFound {
			newNode := newNode(nextChar, nextEnd)
			node.next = append(node.next, newNode)
			node = newNode
		}

	}
	return nil
}

/* Runs through the tree to fetch the word */
func (t *Tree) WordExists(word string) bool {
	if t.start == nil {
		return false
	}
	//check for a single char word
	if len(word) == 1 && t.start.char == rune(word[0]) && t.start.end == true {
		return true
	}
	node := t.start
	lastCharIndex := len(word) - 1
	for i, c := range word {
		if node.char == c && node.end == true && i == lastCharIndex {
			return true
		} else if i == lastCharIndex {
			return false
		} else {
			// Check the next node for the char
			nextChar := rune(word[i+1])
			nextFound := false
			for _, n := range node.next {
				if n.char == nextChar {
					nextFound = true
					node = n
					break
				}
			}
			if !nextFound {
				return false
			}
		}
	}
	return false
}

func (t Tree) String() string {
	s := fmt.Sprintln("Tree (beginning with char): ", string(t.char), "\n", t.start)
	return s
}

/* TreeCollection holds all the trees */
type TreeCollection struct {
	Trees map[rune]*Tree
}

/* Creates new trees if they don't exist and adds the word to it */
func (tc *TreeCollection) AddWord(word string) error {
	if len(word) == 0 {
		return errors.New("empty word")
	}
	firstWord := rune(word[0])
	if tc.Trees == nil {
		tc.Trees = make(map[rune]*Tree)
	}
	if tc.Trees[firstWord] == nil {
		tc.Trees[firstWord] = &Tree{char: firstWord}
	}
	return tc.Trees[firstWord].AddWord(word)
}

func (tc *TreeCollection) WordExists(word string) bool {
	if len(word) == 0 {
		return false
	}
	firstWord := rune(word[0])
	if tc.Trees[firstWord] == nil {
		return false
	}
	return tc.Trees[firstWord].WordExists(word)
}

/* Exposed method to determine if the word is compound */
func (tc *TreeCollection) IsCompound(word string) bool {
	return tc.isCompound(word, word)
}

/* A method called recursively to  determine compound word */
func (tc *TreeCollection) isCompound(word, subword string) bool {
	if len(word) != len(subword) && tc.WordExists(subword) {
		return true
	}
	for i := 1; i < (len(subword) - 1); i++ {
		if tc.WordExists(subword[:i]) && tc.isCompound(word, subword[i:]) {
			return true
		}
	}
	return false
}
