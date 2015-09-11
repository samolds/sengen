#ifndef GRAMMARSOLVER_H
#define GRAMMARSOLVER_H

#include <string>
#include <map>

#define MAX_TOKENS 64

class GrammarSolver {
  public:
    typedef std::map<std::string, std::string *> sentenceTree_t;
    typedef std::map<std::string, int> availGrams_t;

	  GrammarSolver(std::string * grammar, int grammarSize);
	  ~GrammarSolver();

    bool grammarContains(std::string symbol);
    void generate(std::string symbol, std::string * phrases, int times);
    std::string getSymbols();

  private:
    sentenceTree_t sentenceTree;
    int sentenceTreeLength;
    availGrams_t availGrams;

    std::string generatePhrase(std::string phrase, std::string symbol);
    int parse(std::string toParse, std::string delimeter, std::string * tokens, int numTokens);
};

#endif
