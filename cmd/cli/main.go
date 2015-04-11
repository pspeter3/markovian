package main

import (
	"flag"
	"fmt"
	"github.com/pspeter3/markovian"
	"math/rand"
	"os"
	"time"
)

var ngramSize int
var numWords int

func init() {
	flag.IntVar(&ngramSize, "ngram", 2, "The ngram size for the word")
	flag.IntVar(&numWords, "words", 100, "The number or words to print")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
}

func main() {
	chain := markovian.ParseMarkovChain(os.Stdin, ngramSize)
	fmt.Println(chain.Generate(numWords))
}
