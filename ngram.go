package markovian

import (
	"strings"
)

type Ngram []string

func (n Ngram) ToKey() string {
	return strings.Join(n, " ")
}

func (n Ngram) Next(nextWord string) Ngram {
	size := len(n)
	ngram := make(Ngram, size)
	copy(ngram, n[1:])
	ngram[size-1] = nextWord
	return ngram
}

func NewNgram(text string) Ngram {
	return strings.Split(text, " ")
}
