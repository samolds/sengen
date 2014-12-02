package main

import (
  "strings"
  "math/rand"
  "sort"
)


func grammarContains(symbol string, sentenceTree map[string][]string) bool {
  _, contains := sentenceTree[symbol]
  return contains
}


func generate(symbol string, times int, sentenceTree map[string][]string) []string {
  if !grammarContains(symbol, sentenceTree) {
    panic("The given rule is not defined in your list.")
  }
  if times < 0 {
    panic("I can't print something negative times.")
  }
  phrases := make([]string, 0, times)
  for i := 0; i < times; i++ {
    phrases = append(phrases, generatePhrase("", symbol, sentenceTree))
  }
  return phrases
}


func generatePhrase(phrase string, symbol string, sentenceTree map[string][]string) string {
  if !grammarContains(symbol, sentenceTree) {
    phrase = phrase + " " + symbol
  } else {
    randnum := rand.Intn(len(sentenceTree[symbol]))
    tabSplit := func(c rune) bool {
      return c == '\t' || c == ' '
    }
    symbols := strings.FieldsFunc(sentenceTree[symbol][randnum], tabSplit)
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


func buildSolver(grammar []string) map[string][]string {
  if len(grammar) == 0 {
    panic("Your list is empty.")
  }
  sentenceTree := make(map[string][]string)
  for _, gram := range grammar {
    line := strings.Split(gram, "::=")
    grammarRule := line[0]
    if grammarContains(grammarRule, sentenceTree) {
      panic("You have the same nonterminal defined more than once.")
    }
    pipeSplit := func(c rune) bool {
      return c == '|'
    }
    grammarMatch := strings.FieldsFunc(strings.TrimSpace(line[1]), pipeSplit)
    sentenceTree[grammarRule] = grammarMatch
  }
  return sentenceTree
}
