package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "math/rand"
  "sort"
)


func check(e error) {
  if e != nil {
    panic(e)
  }
}


//	public boolean grammarContains(String symbol) {
//		return sentenceTree.containsKey(symbol);
//	}
func grammar_contains(symbol string, sentence_tree map[string][]string) bool {
  _, contains := sentence_tree[symbol]
  return contains
}


//	public String[] generate(String symbol, int times) {
//		if (!grammarContains(symbol)) {
//			throw new IllegalArgumentException("The given rule is not defined in your list.");
//		}
//		if (times < 0) {
//			throw new IllegalArgumentException("I can't print something negative times.");
//		}
//		String[] phrases = new String[times];
//		for (int i = 0; i < times; i++) { // building the array of random grammars
//			phrases[i] = generatePhrase("", symbol);
//		}
//		return phrases;
//	}
func generate(symbol string, times int, sentence_tree map[string][]string) []string {
  if !grammar_contains(symbol, sentence_tree) {
    panic("The given rule is not defined in your list.")
  }
  if times < 0 {
    panic("I can't print something negative times.")
  }
  phrases := make([]string, 0, times)
  for i := 0; i < times; i++ {
    phrases = append(phrases, generate_phrase("", symbol, sentence_tree))
  }
  return phrases
}


//	private String generatePhrase(String phrase, String symbol) {
//		if (!grammarContains(symbol)) { // if the symbol isn't in the keyset (base case)
//			phrase = phrase + " " + symbol;
//		} else { // if the symbol is in the keyset
//			int randnum = (int)(Math.random() * (sentenceTree.get(symbol).length)); // grabs random one
//			String[] symbols = sentenceTree.get(symbol)[randnum].split("[ \t]+"); 	// and splits it up
//			for (int i = 0; i < symbols.length; i++) { // loops through each thing separated by space
//				phrase = generatePhrase(phrase, symbols[i]);
//			}
//		}
//		return phrase.trim();
//	}
func generate_phrase(phrase string, symbol string, sentence_tree map[string][]string) string {
  if !grammar_contains(symbol, sentence_tree) {
    phrase = phrase + " " + symbol
  } else {
    randnum := rand.Intn(len(sentence_tree[symbol]))
    tab_split := func(c rune) bool {
      return c == '\t'
    }
    symbols := strings.FieldsFunc(sentence_tree[symbol][randnum], tab_split)
    for i := 0; i < len(symbols); i++ {
      phrase = generate_phrase(phrase, symbols[i], sentence_tree)
    }
  }
  return strings.TrimSpace(phrase)
}


//	public String getSymbols() {
//		return sentenceTree.keySet().toString();
//	}
func get_symbols(sentence_tree map[string][]string) string {
  keys := make([]string, 0, len(sentence_tree))
  key_string := ""
  for key := range sentence_tree {
    keys = append(keys, key)
  }
  sort.Strings(keys)
  for i := 0; i < len(keys); i++ {
    key_string = key_string + " " + keys[i]
  }
  return strings.TrimSpace(key_string)
}


func show_results(solver map[string][]string) {
  done := false
  for !done {
    fmt.Println("\nAvailable symbols to generate are:")
    fmt.Println(get_symbols(solver))
    var target string
    fmt.Print("What do you want generated (return to quit)? ")
    fmt.Scanf("%s", &target)
    if len(target) == 0 {
      done = true;
    } else if !grammar_contains(target, solver) {
      fmt.Println("Illegal symbol")
    } else {
      var number int
      fmt.Print("How many do you want me to generate? ")
      n, err := fmt.Scanf("%d", &number)
      if err != nil || n != 1 {
        fmt.Println("that's not an integer")
      } else {
        if number < 0 {
          fmt.Println("no negatives allowed")
        } else {
          answers := generate(target, number, solver)
          for _, answer := range answers {
            fmt.Println(answer)
          }
        }
      }
    }
  }
}


//	public GrammarSolver(List<String> grammar) {
//		if (grammar.isEmpty()) {
//			throw new IllegalArgumentException("Your list is empty.");
//		}
//		sentenceTree = new TreeMap<String, String[]>();
//		for (int i = 0; i < grammar.size(); i++) {
//			String[] line = grammar.get(i).toString().split("::=");
//			String grammarRule = line[0];
//			if (grammarContains(grammarRule)) {
//				throw new
//					IllegalArgumentException("You have the same nonterminal defined more than once.");
//			}
//			String[] grammarMatch = line[1].trim().split("[|]+");
//			sentenceTree.put(grammarRule, grammarMatch);
//		}
//	}
func build_solver(grammar []string) map[string][]string {
  if len(grammar) == 0 {
    panic("Your list is empty.")
  }
  sentence_tree := make(map[string][]string)
  for _, gram := range grammar {
    line := strings.Split(gram, "::=")
    grammar_rule := line[0]
    if grammar_contains(grammar_rule, sentence_tree) {
      panic("You have the same nonterminal defined more than once.")
    }
    pipe_split := func(c rune) bool {
      return c == '|'
    }
    grammar_match := strings.FieldsFunc(strings.TrimSpace(line[1]), pipe_split)
    sentence_tree[grammar_rule] = grammar_match
  }
  return sentence_tree
}


func main() {
  fmt.Println("Welcome to the cse143 random sentence generator.\n")

  var filename string
  fmt.Print("What is the name of the grammar file? ")
  fmt.Scanf("%s", &filename)

  file, err := os.Open(filename)
  check(err)
  var grammar []string
  input := bufio.NewScanner(file)
  for input.Scan() {
    grammar = append(grammar, input.Text())
  }
  file.Close()

  solver := build_solver(grammar)
  show_results(solver)
}
