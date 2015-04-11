package markovian_test

import (
	"github.com/pspeter3/markovian"
	"strings"
	"testing"
)

func TestChooseReturnsRandomWord(t *testing.T) {
	chain := markovian.NewMarkovChain(map[string][]string{
		"test": {"foo", "bar"},
	}, 1)
	choice, err := chain.Choose(markovian.NewNgram("test"))
	if err != nil {
		t.Fatal(err)
	}
	if choice != "foo" && choice != "bar" {
		t.Fatalf("Invalid choice %s", choice)
	}
}

func TestChooseReturnsErrorIfKeyDoesNotExist(t *testing.T) {
	chain := markovian.NewMarkovChain(map[string][]string{
		"test": {"foo", "bar"},
	}, 1)
	_, err := chain.Choose(markovian.NewNgram("DNE"))
	if err == nil {
		t.Fatal(err)
	}
}

func TestParseNgrams(t *testing.T) {
	keys := []string{"hello", "world", "foo", "bar"}
	reader := strings.NewReader(strings.Join(keys, " "))
	graph := markovian.ParseNgrams(reader, 1)
	for index, key := range keys[:len(keys)-1] {
		value, ok := graph[key]
		if !ok {
			t.Fatalf("Key %s not in graph", key)
		}
		if value[0] != keys[index+1] {
			t.Fatalf("Found wrong value for key. Expected: %s, Actual: %s",
				key[index+1], value[0])
		}
	}
}

func TestGenerateWithEnoughWords(t *testing.T) {
	chain := markovian.ParseMarkovChain(strings.NewReader("hello world"), 1)
	text := chain.Generate(1)
	if text != "hello" {
		t.Fatalf("Text was %s", text)
	}
}

func TestGenerateWithNotEnoughWords(t *testing.T) {
	chain := markovian.ParseMarkovChain(strings.NewReader("hello world"), 1)
	text := chain.Generate(3)
	if text != "hello world" {
		t.Fatalf("Text was %s", text)
	}
}
