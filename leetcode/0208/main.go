package main

type Node struct {
	children map[rune]*Node
	isWord   bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children: make(map[rune]*Node),
			isWord:   false,
		},
	}
}

func (this *Trie) Insert(word string) {
	cur := this.root

	for _, char := range word {
		if cur.children[char] == nil {
			cur.children[char] = &Node{
				children: make(map[rune]*Node),
				isWord:   false,
			}
		}
		cur = cur.children[char]
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root

	for _, char := range word {
		if cur.children[char] == nil {
			return false
		}
		cur = cur.children[char]
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root

	for _, char := range prefix {
		if cur.children[char] == nil {
			return false
		}
		cur = cur.children[char]
	}
	return true
}
