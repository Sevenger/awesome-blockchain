package trie

type MapTreeNode struct {
	Child  map[rune]*MapTreeNode
	IsWord bool
}

func NewMapTrieTree() *MapTreeNode {
	return &MapTreeNode{
		Child: make(map[rune]*MapTreeNode),
	}
}

func (root *MapTreeNode) Insert(word string) {
	cur := root
	for _, s := range word {
		if _, ok := cur.Child[s]; !ok {
			cur.Child[s] = NewMapTrieTree()
		}
		cur = cur.Child[s]
	}
	cur.IsWord = true
}

func (root *MapTreeNode) Search(word string) bool {
	tail := root.getTailNode(word)
	return tail != nil && tail.IsWord
}

func (root *MapTreeNode) StartWith(prefix string) bool {
	return root.getTailNode(prefix) != nil
}

func (root *MapTreeNode) getTailNode(word string) *MapTreeNode {
	cur := root
	for _, s := range word {
		if _, ok := cur.Child[s]; !ok {
			return nil
		}
		cur = cur.Child[s]
	}
	return cur
}
