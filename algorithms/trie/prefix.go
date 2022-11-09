const ALPHABET_SIZE = 26

type Trie struct {
    children []*Trie
    isEndOfWord bool
}

func Constructor() Trie {
    return Trie{children: make([]*Trie, ALPHABET_SIZE)}
}

func (this *Trie) Insert(word string) {
    prev := this
    for i := 0; i < len(word); i++ {
        id := this.id(word[i])
        trie := &Trie{children: make([]*Trie, ALPHABET_SIZE)}
        if prev.children[id] == nil {
            prev.children[id] = trie
        }
        prev = prev.children[id]
    }
    prev.isEndOfWord = true
}

func (this *Trie) Search(word string) bool {
    prev := this
    for i := 0; i < len(word); i++ {
        if prev = prev.children[this.id(word[i])]; prev == nil {
            return false
        }
    }
    return prev.isEndOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
    prev := this
    for i := 0; i < len(prefix); i++ {
        if prev = prev.children[this.id(prefix[i])]; prev == nil {
            return false
        }
    }
    return true
}

func (this *Trie) id(c byte) int {
    return int(c) % 32 - 1
}
