#include "GrammarSolver.h"
#include <stdlib.h>

using namespace std;


/**
 * Description: Constructor that builds up a map from a list of lines matching the
 *              syntax in the sentence.txt and formula.txt files. Maps a string
 *              to a list of strings
 *
 *  Parameters: grammar: List of lines from the grammar file
 *              grammarSize: Number of lines
 */
GrammarSolver::GrammarSolver(string grammar[], int grammarSize) {
  if (grammarSize == 0) {
    throw; // "Your list is empty"
  }

  sentenceTreeLength = grammarSize;
  int i, numElements = 0;
  for (i = 0; i < grammarSize; i++) {
    string line = grammar[i];
    string tokens[2];
    parse(line, "::=", tokens, 2);
    string grammarRule = tokens[0];
    if (grammarContains(grammarRule)) {
      throw; // "You have the same nonterminal defined more than once."
    }
    string * grammarMatch = new string[MAX_TOKENS];
    numElements = parse(tokens[1], "|", grammarMatch, MAX_TOKENS);
    availGrams[grammarRule] = numElements;
    sentenceTree[grammarRule] = grammarMatch;
  }
}


/**
 * Description: Deconstructor that deletes list of strings from sentenceTree
 */
GrammarSolver::~GrammarSolver() {
  for (sentenceTree_t::iterator iter(sentenceTree.begin()); iter != sentenceTree.end(); ++iter) {
    delete [] iter->second; // grammarMatch
  }
}


/**
 * Description: Evaluates if a specific symbol is specified in a grammar
 *
 *  Parameters: symbol: the symbol to be searched for
 */
bool GrammarSolver::grammarContains(string symbol) {
  return sentenceTree.count(symbol) != 0;
}


/**
 * Description: Generates a certain number of phrases according to the grammar rules
 *              specified.
 *
 *  Parameters: symbol: The type of phrase to generate
 *              phrases: The list that all of the generated phrases will be added to
 *              times: The number of phrases to generate
 */
void GrammarSolver::generate(string symbol, string * phrases, int times) {
  if (!grammarContains(symbol)) {
    throw; // "The given rule is not defined in your list."
  }
  if (times < 0) {
    throw; // "I can't print something negative times."
  }

  int i = 0;
  for (i = 0; i < times; i++) {
    phrases[i] = generatePhrase("", symbol);
  }
}


/**
 * Description: Returns if a character is between '!' and '~' in the ascii table
 *
 *  Parameters: c: The character to evaluate
 */
bool invalidChar(char c) {
  return !(c >= 33 && c <= 126);
}


/**
 * Description: Removes all invalid characters from a string
 *
 *  Parameters: str: The string to remove characters from
 */
void strip(string & str) {
  str.erase(remove_if(str.begin(), str.end(), invalidChar), str.end());
}


/**
 * Description: Recursively builds up a random phrase
 *
 *  Parameters: phrase: The phrase being built
 *              symbol: The type of phrase to generate
 */
string GrammarSolver::generatePhrase(string phrase, string symbol) {
  int i, randnum, numSymbols = 0;
  strip(symbol);
  if (!grammarContains(symbol)) {
    phrase = phrase + symbol + " ";
  } else {
    i = 0;
    randnum = rand() % availGrams[symbol];
    string * symbols = new string[MAX_TOKENS];
    numSymbols = parse(sentenceTree[symbol][randnum], " ", symbols, MAX_TOKENS);
    for (i = 0; i < numSymbols; i++) {
      phrase = generatePhrase(phrase, symbols[i]);
    }
    delete [] symbols;
  }

  return phrase;
}


/**
 * Description: Generates the list of keys from the sentenceTree map of available
 *              types of phrases to construct.
 */
string GrammarSolver::getSymbols() {
  string symbols;
  for (sentenceTree_t::iterator iter(sentenceTree.begin()); iter != sentenceTree.end(); ++iter) {
    symbols = symbols + iter->first + " ";
  }
  symbols.erase(symbols.size() - 1);
  return symbols;
}


/**
 * Description: Parses a string into a list of tokens with a delimiter.
 *
 *  Parameters: toParse: The string to parse
 *              delimeter: The delimiter to parse with
 *              tokens: The list of strings being found
 *              maxTokens: The maximum allowed number of found tokens
 */
int GrammarSolver::parse(string toParse, string delimiter, string * tokens, int maxTokens) {
  size_t pos = 0;
  int len = 0;

  while ((pos = toParse.find(delimiter)) != std::string::npos) {
    tokens[len++] = toParse.substr(0, pos);
    toParse.erase(0, pos + delimiter.length());
    if (len >= maxTokens) {
      throw; // "Too many tokens!"
    }
  }
  tokens[len++] = toParse;

  return len;
}
