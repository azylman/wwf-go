package main

import (
	"container/list"
	"log"
	"wwf/dict"
	"wwf/solveresult"
)

func main() {
	aDict := dict.New("dict.txt")
	var wordsToTest []string = []string{"aaaa", "god", "dagoba"}
	for _, word := range wordsToTest {
		if aDict.IsWord(word) {
			log.Printf("%s is a word with a score of %d", word, dict.Score(word))
		} else {
			log.Printf("%s is not a word", word)
		}
	}
	results := list.New()
	aDict.Solve("godasd*", "", "", "", results)
	for result := results.Front(); result != nil; result = result.Next() {
		theResult := result.Value.(solveresult.SolveResult)
		log.Printf("Got result %s with score %d and length %d", theResult.GetWord(), theResult.GetScore(), theResult.GetLength())
	}
}
