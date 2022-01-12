package trie

const ASCII = 256

type Node struct {
	Child  [ASCII]*Node
	IsWord bool
}

func NewTrieTree() *Node {
	return &Node{
		Child: [ASCII]*Node{},
	}
}

func (root *Node) Insert(word string) {
	cur := root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if cur.Child[c] == nil {
			cur.Child[c] = NewTrieTree()
		}
		cur = cur.Child[c]
	}
	cur.IsWord = true
}

func (root *Node) Query(word string) bool {
	tail := root.getTailNode(word)
	return tail != nil && tail.IsWord
}

func (root *Node) HasPrefix(prefix string) bool {
	return root.getTailNode(prefix) != nil
}

func (root *Node) getTailNode(word string) *Node {
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
