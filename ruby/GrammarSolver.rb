import random
import re


class GrammarSolver:
    __sentence_tree = {}

    def __init__(self, grammar):
        if len(grammar) == 0:
            raise ValueError("Grammar must not be empty.")
        for gram in grammar:
            line = gram.split("::=")
            grammar_rule = line[0]
            if self.grammar_contains(grammar_rule):
                raise ValueError("You have the same nonterminal defined more than once.")
            grammar_match = re.split("[|]+", line[1].strip())
            self.__sentence_tree.update({grammar_rule: grammar_match})

    def grammar_contains(self, symbol):
        return self.__sentence_tree.has_key(symbol)

    def generate(self, symbol, times):
        if not self.grammar_contains(symbol):
            raise ValueError("The given rule is not defined in your list.")
        if times < 0:
            raise ValueError("I can't print something negative times.")
        phrases = []
        i = 0
        while (i < times):
            phrases.append(self.__generate_phrase("", symbol))
            i = i + 1
        return phrases

    def __generate_phrase(self, phrase, symbol):
        if not self.grammar_contains(symbol):
            phrase = phrase + " " + symbol
        else:
            randnum = random.randint(0, len(self.__sentence_tree[symbol]) - 1)
            symbols = re.split("[ \t]+", self.__sentence_tree[symbol][randnum])
            for symbol in symbols:
                phrase = self.__generate_phrase(phrase, symbol)
        return phrase.strip()

    def get_symbols(self):
        return ', '.join(self.__sentence_tree.keys())
