package main

import (
	"errors"
	"math/rand"
	"sort"
	"strings"
)

func grammarContains(symbol string, sentenceTree map[string][]string) bool {
	_, contains := sentenceTree[symbol]
	return contains
}

func generate(symbol string, times int, sentenceTree map[string][]string) ([]string, error) {
	if !grammarContains(symbol, sentenceTree) {
		return nil, errors.New("The given rule is not defined in your list.")
	}
	if times < 0 {
		return nil, errors.New("I can't print something negative times.")
	}
	phrases := make([]string, 0, times)
	for i := 0; i < times; i++ {
		phrases = append(phrases, generatePhrase("", symbol, sentenceTree))
	}
	return phrases, nil
}

func generatePhrase(phrase string, symbol string, sentenceTree map[string][]string) string {
	if !grammarContains(symbol, sentenceTree) {
		phrase = phrase + " " + symbol
	} else {
		randnum := rand.Intn(len(sentenceTree[symbol]))
		symbols := strings.Fields(sentenceTree[symbol][randnum])
		for i := 0; i < len(symbols); i++ {
			phrase = generatePhrase(phrase, symbols[i], sentenceTree)
		}
	}
	return strings.TrimSpace(phrase)
}

func getSymbols(sentenceTree map[string][]string) string {
	keys := make([]string, 0, len(sentenceTree))
	keyString := ""
	for key := range sentenceTree {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for i := 0; i < len(keys); i++ {
		keyString = keyString + " " + keys[i]
	}
	return strings.TrimSpace(keyString)
}

func buildSolver(grammar []string) (map[string][]string, error) {
	if len(grammar) == 0 {
		return nil, errors.New("Your list is empty.")
	}
	sentenceTree := make(map[string][]string)
	for _, gram := range grammar {
		line := strings.Split(gram, "::=")
		grammarRule := line[0]
		if grammarContains(grammarRule, sentenceTree) {
			return nil, errors.New("You have the same nonterminal defined more than once.")
		}
		pipeSplit := func(c rune) bool {
			return c == '|'
		}
		grammarMatch := strings.FieldsFunc(strings.TrimSpace(line[1]), pipeSplit)
		sentenceTree[grammarRule] = grammarMatch
	}
	return sentenceTree, nil
}
