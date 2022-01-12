package trie

const ASCII = 256

type ArrayTreeNode struct {
	Child  [ASCII]*ArrayTreeNode
	IsWord bool
}

func NewArrayTrieTree() *ArrayTreeNode {
	return &ArrayTreeNode{}
}

func (root *ArrayTreeNode) Insert(word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.Child[c] == nil {
			cur.Child[c] = NewArrayTrieTree()
		}
		cur = cur.Child[c]
	}
	cur.IsWord = true
}

func (root *ArrayTreeNode) Search(word string) bool {
	tail := root.getTailNode(word)
	return tail != nil && tail.IsWord
}

func (root *ArrayTreeNode) StartWith(prefix string) bool {
	return root.getTailNode(prefix) != nil
}

func (root *ArrayTreeNode) getTailNode(word string) *ArrayTreeNode {
	cur := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.Child[c] == nil {
			return nil
		}
		cur = cur.Child[c]
	}
	return cur
}
