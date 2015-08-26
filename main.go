package main

// Mathias Monnerville <mathias@monnerville.com>

import (
	"bufio"
	"container/list"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

// Node holding a letter and possible children. Boundary
// is true when word (root node to current node) is a valid word
// from the input list.
type Node struct {
	letter   string
	children map[string]*Node
	boundary bool // Marks a word boundary
}

func newNode(letter string) *Node {
	n := &Node{
		letter:   letter,
		children: make(map[string]*Node),
		boundary: false,
	}
	return n
}

// Prefix tree
type Tree struct {
	root *Node
}

// newTree creates a tree with a valid root node.
func newTree() *Tree {
	return &Tree{newNode("")}
}

// addWord inserts a word in the tree data structure.
func (t *Tree) addWord(word string) {
	curnode := t.root // Current node
	for _, code := range word {
		letter := fmt.Sprintf("%c", code)
		if _, ok := curnode.children[letter]; !ok {
			curnode.children[letter] = newNode(letter)
		}
		curnode = curnode.children[letter]
	}
	curnode.boundary = true // End of word
}

// findWordPrefixes looks for all prefixes of a given word.
func (t *Tree) findWordPrefixes(word string) []string {
	prefixes := make([]string, 0)
	var prefix string
	curnode := t.root

	for _, code := range word {
		letter := fmt.Sprintf("%c", code)
		if _, ok := curnode.children[letter]; !ok {
			return prefixes
		}
		prefix += letter
		curnode = curnode.children[letter]
		if curnode.boundary {
			prefixes = append(prefixes, prefix)
		}
	}
	return prefixes
}

// exists looks for the word in the tree.
func (t *Tree) exists(word string) bool {
	curnode := t.root
	for _, code := range word {
		letter := fmt.Sprintf("%c", code)
		if _, ok := curnode.children[letter]; !ok {
			return false
		}
		curnode = curnode.children[letter]
	}
	return curnode.boundary
}

type pair struct {
	word, suffix string
}

func findLongestCompoundWord(r io.Reader) string {
	tree := newTree()
	pq := list.New() // Processing queue
	scanner := bufio.NewScanner(r)

	// Insert words into the tree data structure
	count := 0
	for scanner.Scan() {
		word := scanner.Text()
		prefixes := tree.findWordPrefixes(word)
		for _, prefix := range prefixes {
			pq.PushBack(&pair{word, word[len(prefix):]})
		}
		tree.addWord(word)
		count++
	}

	// Now process pairs in the queue
	var match string
	maxLen := 0
	for pq.Len() > 0 {
		p := pq.Remove(pq.Front()).(*pair)
		if tree.exists(p.suffix) && len(p.word) > maxLen {
			maxLen = len(p.word)
			match = p.word
		} else {
			prefixes := tree.findWordPrefixes(p.suffix)
			for _, prefix := range prefixes {
				pq.PushBack(&pair{p.word, p.suffix[len(prefix):]})
			}
		}
	}
	return match
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, fmt.Sprintf("Usage: %s wordfile\n", os.Args[0]))
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Fprint(os.Stderr, "word file is missing.\n\n")
		flag.Usage()
	}
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(findLongestCompoundWord(f))
}
