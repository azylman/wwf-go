package solveresult

type SolveResult struct {
	word   string
	score  int
	length int
}

func (r SolveResult) GetWord() string {
	return r.word
}

func (r SolveResult) GetScore() int {
	return r.score
}

func (r SolveResult) GetLength() int {
	return r.length
}

func New(word string, score int) SolveResult {
	return SolveResult{word: word, score: score, length: len(word)}
}
