





//public class GrammarSolver {
//	private SortedMap<String, String[]> sentenceTree; // Stores the nonterminal and terminal grammars
//	
//	//  pre: If the list of rules passed in is actually empty, or a rule appears more than
//	//			once in the parameter, an IllegalArgumentException is thrown. This program
//	// 		requires that the inputted list of rules is in the Backus-Naur Form (BNF) format
//	// post: Constructs an object that will hold all of the nonterminal rules pointing at their
//	// 		repsective rules
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
//
//	// post: Returns if the symbol is a nonterminal grammar rule; it is case sensitive
//	public boolean grammarContains(String symbol) {
//		return sentenceTree.containsKey(symbol);
//	}
//	
//	//  pre: If the nonterminal rule being asked to be generated is not defined, or if the number
//	// 		of times it was asked to be generated is less than zero, an IllegalArgumentException
//	// 		is thrown
//	// post: Builds the series of randomly generated grammars with the help of a private method
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
//
//	// post: Helper for the generate method that builds each individual random grammar 
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
//	
//	// post: Returns a list of all of the nonterminal grammar rules that are available for the
//	// 		user to call to create new grammars
//	public String getSymbols() {
//		return sentenceTree.keySet().toString();
//	}
//}
