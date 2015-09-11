#include "GrammarSolver.h"
#include <stdlib.h>

using namespace std;


/**
 * Description: Constructor that builds up a map of strings to a list of strings
 *
 *  Parameters: grammar:
 *              grammarSize:
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
 * Description: Deconstructor
 */
GrammarSolver::~GrammarSolver() {
  for (sentenceTree_t::iterator iter(sentenceTree.begin()); iter != sentenceTree.end(); ++iter) {
    delete [] iter->second; // grammarMatch
  }
}


/**
 * Description: 
 *
 *  Parameters: symbol:
 */
bool GrammarSolver::grammarContains(string symbol) {
  return sentenceTree.count(symbol) != 0;
}


/**
 * Description: 
 *
 *  Parameters: symbol:
 *              phrases:
 *              times:
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
 * Description: 
 *
 *  Parameters: phrase:
 *              symbole:
 */
string GrammarSolver::generatePhrase(string phrase, string symbol) {
  int i, randnum, numSymbols = 0;
  if (!grammarContains(symbol)) {
    phrase = phrase + " " + symbol;
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

  if (phrase.size() > 0) {
    phrase.erase(phrase.size() - 1);
  }
  return phrase;
}


/**
 * Description: 
 *
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
 * Description: 
 *
 *  Parameters: toParse:
 *              delimeter:
 *              tokens:
 *              numTokens:
 */
int GrammarSolver::parse(string toParse, string delimiter, string * tokens, int numTokens) {
  size_t pos = 0;
  int len = 0;

  while ((pos = toParse.find(delimiter)) != std::string::npos) {
    tokens[len++] = toParse.substr(0, pos);
    toParse.erase(0, pos + delimiter.length());
    if (len >= numTokens) {
      throw; // "Too many tokens!"
    }
  }
  tokens[len++] = toParse;

  return len;
}
