package markovian

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"strings"
)

type MarkovChain struct {
	graph     map[string][]string
	ngramSize int
}

func (m *MarkovChain) Choose(ngram Ngram) (error, string) {
	choices := m.graph[ngram.ToKey()]
	size := len(choices)
	if size == 0 {
		return errors.New("No more choices"), ""
	}
	return nil, choices[rand.Intn(size)]
}

func (m *MarkovChain) Generate(numWords int) string {
	ngram := make(Ngram, m.ngramSize)
	var words []string
	for i := 0; i < numWords; i++ {
		err, nextWord := m.Choose(ngram)
		if err != nil {
			break
		}
		words = append(words, nextWord)
		ngram = ngram.Next(nextWord)
	}
	return strings.Join(words, " ")
}

func NewMarkovChain(graph map[string][]string, ngramSize int) *MarkovChain {
	return &MarkovChain{graph, ngramSize}
}

func ParseNgrams(reader io.Reader, ngramSize int) map[string][]string {
	buffer := bufio.NewReader(reader)
	ngram := make(Ngram, ngramSize)
	graph := make(map[string][]string)
	for {
		var nextWord string
		if _, err := fmt.Fscan(buffer, &nextWord); err != nil {
			break
		}
		key := ngram.ToKey()
		graph[key] = append(graph[key], nextWord)
		ngram = ngram.Next(nextWord)
	}
	return graph
}

func ParseMarkovChain(reader io.Reader, ngramSize int) *MarkovChain {
	return NewMarkovChain(ParseNgrams(reader, ngramSize), ngramSize)
}
