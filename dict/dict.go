package dict

import (
	"bufio"
	"container/list"
	"os"
	"strings"
	"wwf/solveresult"
	"wwf/trie"
)

var SCORES map[rune]int = map[rune]int{
	'a': 1, 'b': 4, 'c': 4, 'd': 2, 'e': 1, 'f': 4, 'g': 3, 'h': 3,
	'i': 1, 'j': 10, 'k': 5, 'l': 2, 'm': 4, 'n': 2, 'o': 1, 'p': 4,
	'q': 5, 'r': 1, 's': 1, 't': 1, 'u': 2, 'v': 5, 'w': 4, 'x': 8,
	'y': 3, 'z': 10}

type Dict struct {
	wordList *trie.Trie
}

func getNode(word string, node *trie.Trie) *trie.Trie {
	if node == nil {
		return nil
	}
	if word == "" {
		return node
	}
	return getNode(word[1:len(word)], node.Next(rune(word[0])))
}

func Score(word string) (score int) {
	score = 0
	for _, c := range word {
		score += SCORES[c]
	}
	return
}

// TODO: Use goroutines and channels for our recursion
func (dict *Dict) wordSearch(word, contains, end string, score, numWildcards int, results *list.List, node *trie.Trie) (count int) {
	count = 0
	if node == nil {
		return
	}

	endNode := getNode(end, node)

	if endNode != nil && contains == "" && endNode.GetWord() != "" {
		results.PushBack(solveresult.New(endNode.GetWord(), score+Score(end)))
	}

	if numWildcards > 0 {
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('a'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('b'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('c'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('d'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('e'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('f'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('g'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('h'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('i'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('j'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('k'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('l'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('m'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('n'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('o'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('p'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('q'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('r'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('s'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('t'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('u'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('v'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('w'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('x'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('y'))
		count += dict.wordSearch(word, contains, end, score, numWildcards-1, results, node.Next('z'))
	}

	for _, c := range word {
		newWord := strings.Join(strings.SplitN(word, string(c), 1), "")
		count += dict.wordSearch(newWord, contains, end, score+Score(string(c)), numWildcards, results, node.Next(c))
	}

	if contains != "" {
		count += dict.wordSearch(word, "", end, score+Score(string(contains)), numWildcards, results, getNode(contains, node))
	}

	count++
	return
}

func (dict *Dict) Solve(word, start, contains, end string, results *list.List) {
	word = strings.ToLower(word)
	start = strings.ToLower(start)
	contains = strings.ToLower(contains)
	end = strings.ToLower(end)

	node := getNode(start, dict.wordList)
	if node == nil {
		return
	}

	numWildcards := strings.Count(word, "*")
	word = strings.Replace(word, "*", "", -1) // Remove all wildcards
	dict.wordSearch(word, contains, end, Score(start), numWildcards, results, node)
}

func (dict *Dict) IsWord(word string) bool {
	result := getNode(strings.ToLower(word), dict.wordList)
	return result != nil && result.GetWord() != ""
}

func New(file string) *Dict {
	wordList := trie.New()
	fin, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			wordList.AddWord(word)
		}
	}
	return &Dict{wordList: wordList}
}
