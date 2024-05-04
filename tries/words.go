package tries

import (
	"strings"
)

type WordsTrie struct {
	isWord   bool
	children []*WordsTrie
}

func NewWordsTrie() WordsTrie {
	return WordsTrie{
		isWord:   false,
		children: make([]*WordsTrie, 26),
	}
}

func (t WordsTrie) getLetterIndex(letter rune) int {
	return int(letter - 'a')
}

func (t *WordsTrie) getOrAddLetter(letter rune) *WordsTrie {
	idx := t.getLetterIndex(letter)

	if t.children[idx] == nil {
		letterTrie := NewWordsTrie()
		t.children[idx] = &letterTrie
	}

	return t.children[idx]
}

func (t *WordsTrie) AddWord(word string) {
	ref := t
	for _, c := range strings.ToLower(word) {
		ref = ref.getOrAddLetter(c)
	}

	ref.isWord = true
}

func (t *WordsTrie) getWords(currWord string, words []string) []string {
	for i, w := range t.children {
		if w == nil {
			continue
		}

		letter := rune(i) + 'a'
		newCurrWord := currWord + string(letter)

		if w.isWord {
			words = append(words, newCurrWord)
		}

		words = w.getWords(newCurrWord, words)
	}

	return words
}

func (t *WordsTrie) Autocomplete(input string) []string {
	root := t
	for _, c := range input {
		idx := t.getLetterIndex(c)
		root = root.children[idx]
		if root == nil {
			return []string{}
		}
	}

	return root.getWords(input, []string{})
}
