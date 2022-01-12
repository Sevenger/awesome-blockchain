package trie

type TrieTree interface {
	Insert(string)
	Search(string) bool
	StartWith(string) bool
}
