package trie

type Trie struct {
	word     string
	size     int
	children map[rune]*Trie
}

func (t *Trie) GetWord() string {
	return t.word
}

func (t *Trie) GetSize() int {
	return t.size
}

func (t *Trie) Next(c rune) *Trie {
	return t.children[c]
}

func (t *Trie) AddWord(word string) {
	t.size++
	temp := t
	for _, c := range word {
		temp = temp.addBranch(c)
	}
	temp.word = word
}

func (t *Trie) addBranch(c rune) *Trie {
	if t.Next(c) == nil {
		t.children[c] = New()
	}
	return t.children[c]
}

func New() *Trie {
	return &Trie{size: 0, children: make(map[rune]*Trie), word: ""}
}
