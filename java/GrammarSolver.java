import java.util.*;

public class GrammarSolver {
	private SortedMap<String, String[]> sentenceTree;
	
	public GrammarSolver(List<String> grammar) {
		if (grammar.isEmpty()) {
			throw new IllegalArgumentException("Your list is empty.");
		}
		sentenceTree = new TreeMap<String, String[]>();
		for (int i = 0; i < grammar.size(); i++) {
			String[] line = grammar.get(i).toString().split("::=");
			String grammarRule = line[0];
			if (grammarContains(grammarRule)) {
				throw new
					IllegalArgumentException("You have the same nonterminal defined more than once.");
			}
			String[] grammarMatch = line[1].trim().split("[|]+");
			sentenceTree.put(grammarRule, grammarMatch);
		}
	}

	public boolean grammarContains(String symbol) {
		return sentenceTree.containsKey(symbol);
	}
	
	public String[] generate(String symbol, int times) {
		if (!grammarContains(symbol)) {
			throw new IllegalArgumentException("The given rule is not defined in your list.");
		}
		if (times < 0) {
			throw new IllegalArgumentException("I can't print something negative times.");
		}
		String[] phrases = new String[times];
		for (int i = 0; i < times; i++) {
			phrases[i] = generatePhrase("", symbol);
		}
		return phrases;
	}

	private String generatePhrase(String phrase, String symbol) {
		if (!grammarContains(symbol)) {
			phrase = phrase + " " + symbol;
		} else {
			int randnum = (int)(Math.random() * (sentenceTree.get(symbol).length));
			String[] symbols = sentenceTree.get(symbol)[randnum].split("[ \t]+");
			for (int i = 0; i < symbols.length; i++) {
				phrase = generatePhrase(phrase, symbols[i]);
			}
		}
		return phrase.trim();
	}
	
	public String getSymbols() {
		return sentenceTree.keySet().toString();
	}
}
