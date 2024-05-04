package tries

import (
	"slices"
	"testing"
)

func TestWordTrie(t *testing.T) {
	trie := NewWordsTrie()
	trie.AddWord("cat")
	trie.AddWord("cats")
	trie.AddWord("cattle")
	trie.AddWord("catastrophe")
	trie.AddWord("card")
	trie.AddWord("marc")

	expected := []string{"card", "cat", "catastrophe", "cats", "cattle"}
	if result := trie.Autocomplete("ca"); !slices.Equal(result, expected) {
		t.Fatalf("Autocompleting 'ca' expected %v but got %v", expected, result)
	}

	t.Fatal(trie.Autocomplete("ma"))
}
