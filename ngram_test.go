package markovian_test

import (
	"github.com/pspeter3/markovian"
	"testing"
)

func TestConstructor(t *testing.T) {
	ngram := markovian.NewNgram("hello world")
	if ngram[0] != "hello" {
		t.Fatalf("First value was %s", ngram[0])
	}
	if ngram[1] != "world" {
		t.Fatalf("Second value was %s", ngram[1])
	}
}

func TestNext(t *testing.T) {
	ngram := markovian.NewNgram("hello world")
	next := ngram.Next("today")
	if next.ToKey() != "world today" {
		t.Fatalf("%s was not world today", next.ToKey())
	}
}
func TestToKey(t *testing.T) {
	ngram := make(markovian.Ngram, 2)
	ngram[0] = "hello"
	ngram[1] = "world"
	if ngram.ToKey() != "hello world" {
		t.Fatalf("%s was not hello world", ngram.ToKey())
	}
}
