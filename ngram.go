package markovian

import (
	"strings"
)

type Ngram []string

func (n Ngram) ToKey() string {
	return strings.Join(n, " ")
}

func NewNgram(text string) Ngram {
	return strings.Split(text, " ")
}
