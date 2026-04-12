type TrieNode struct {
	LastChar bool
	Children map[byte]*TrieNode
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
    return PrefixTree{
		root: &TrieNode{
			Children: make(map[byte]*TrieNode),
		},
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this.root
	for ix := range word {
		_, ok := curr.Children[word[ix]]
		if !ok {
			curr.Children[word[ix]] = &TrieNode{
                Children: make(map[byte]*TrieNode),
            }
		}
		curr = curr.Children[word[ix]]
	}
	curr.LastChar = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this.root
	for ix := range word {
		if _, ok := curr.Children[word[ix]]; !ok {
			return false
		}
		curr = curr.Children[word[ix]]
	}
	return curr.LastChar
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this.root
	for ix := range prefix {
		if _, ok := curr.Children[prefix[ix]]; !ok {
			return false
		}
		curr = curr.Children[prefix[ix]]
	}
	return true
}
