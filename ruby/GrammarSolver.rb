class GrammarSolver
  def initialize(grammar)
    @sentence_tree = Hash.new
    
    if grammar.length() == 0
      raise "Grammar must not be empty."
    end
    for gram in grammar
      line = gram.split(pattern="::=")
      grammar_rule = line[0]
      if grammar_contains(grammar_rule)
        raise "You have the same nonterminal defined  more than once."
      end
      grammar_match = line[1].strip().split(pattern=/[|]+/)
      @sentence_tree.store(grammar_rule, grammar_match)
    end
  end

  
  def grammar_contains(symbol)
    return @sentence_tree.has_key?(symbol)
  end

  
  def generate(symbol, times)
    if !grammar_contains(symbol)
      raise "The given rule is not defined in your list."
    end
    if times < 0
      raise "I can't print something negative times."
    end
    phrases = Array.new
    for i in 0..times
      phrases.push(generate_phrase("", symbol))
    end
    return phrases
  end


  def generate_phrase(phrase, symbol)
    if !grammar_contains(symbol)
      phrase = phrase + " " + symbol
    else
      randnum = Random.rand(@sentence_tree[symbol].length())
      symbols = @sentence_tree[symbol][randnum].split(pattern=/[ \t]+/)
      for symbol in symbols
        phrase = generate_phrase(phrase, symbol)
      end
    end
    return phrase.strip()
  end


  def get_symbols
    return @sentence_tree.keys().inspect()
  end
end
