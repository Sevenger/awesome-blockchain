package trie

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	trieTree := NewTrieTree()
	trieTree.Insert("123abc")
	assert(t, trieTree.Query("123abc"), true)
	assert(t, trieTree.Query("123"), false)
	assert(t, trieTree.HasPrefix("123"), true)
	assert(t, trieTree.HasPrefix("123abc123"), false)
}

func assert(t *testing.T, res, except interface{}) {
	s := fmt.Sprintf("res: %v, except: %v", res, except)
	if res != except {
		t.Logf("faild, %s", s)
	}
	t.Logf("success, %s", s)
}
