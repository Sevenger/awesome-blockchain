package trie

import (
	"fmt"
	"testing"
)

func TestTrieTree(t *testing.T) {
	var trieTree TrieTree

	trieTree = NewArrayTrieTree()
	trieTree.Insert("123abc")
	assert(t, trieTree.Search("123abc"), true)
	assert(t, trieTree.Search("123"), false)
	assert(t, trieTree.StartWith("123"), true)
	assert(t, trieTree.StartWith("123abc123"), false)

	trieTree = NewMapTrieTree()
	trieTree.Insert("大波美人鱼人美波大")
	assert(t, trieTree.Search("大波美人鱼人美波大"), true)
	assert(t, trieTree.Search("大波美人鱼"), false)
	assert(t, trieTree.StartWith("大波美人鱼"), true)
	assert(t, trieTree.StartWith("小波美人鱼"), false)
}

func assert(t *testing.T, res, except interface{}) {
	s := fmt.Sprintf("res: %v, except: %v", res, except)
	if res != except {
		t.Logf("faild, %s", s)
		return
	}
	t.Logf("success, %s", s)
}
