extern crate rand;
use self::rand::distributions::{IndependentSample, Range};
use std::collections::HashMap;


// Useful?
// https://github.com/nrc/r4cppp/blob/master/arrays.md
pub type Solver = HashMap<String, Vec<String>>;


pub fn grammar_contains<'a>(grammar: &str, sentence_tree: &Solver) -> bool {
  match sentence_tree.get(grammar) {
    Some(_) => true,
    _ => false,
  }
}


pub fn generate<'a>(target: &'a str, number: i32, sentence_tree: &Solver) -> Result<Vec<String>, &'static str> {
  if !grammar_contains(target, sentence_tree) {
    return Err("The given rule is not defined in your list.");
  }

  if number < 0 {
    return Err("I can't print something negative times.");
  }

  let mut phrases = Vec::new();
  for _ in 0..number {
    phrases.push(generate_phrase("".to_string(), target, sentence_tree));
  }

  return Ok(phrases);
}


fn pick(a: i32, b: i32) -> i32 {
  let between = Range::new(a, b);
  let mut rng = rand::thread_rng();
  return between.ind_sample(&mut rng);
}


fn generate_phrase<'a>(mut phrase: String, symbol: &'a str, sentence_tree: &Solver) -> String {
  if !grammar_contains(symbol, sentence_tree) {
    phrase.push(' ');
    phrase.push_str(symbol);
  } else {
    let randnum: i32 = pick(0, sentence_tree[symbol].len() as i32);
    let mut symbols_iter = sentence_tree[symbol][randnum as usize].split_whitespace();

    loop {
      phrase = match symbols_iter.next() {
        None => break,
        Some(sym) => generate_phrase(phrase, sym, sentence_tree),
      };
    }
  }

  return phrase.trim().to_string();
}


pub fn get_symbols<'a>(sentence_tree: &Solver) -> String {
  let mut keys = String::new();
  for key in sentence_tree.keys() {
    keys.push_str(key);
    keys.push(' ');
  }

  return keys.trim().to_string();
}


pub fn build_solver<'a>(grammar_str: &'a str) -> Result<Solver, &'static str> {
  if grammar_str.len() == 0 {
    return Err("Your list is empty.");
  }

  // Useful?
  // http://rustbyexample.com/std/hash.html
  let mut sentence_tree: Solver = HashMap::new();

  let grammar_lines = grammar_str.split("\n");
  for grammar_line in grammar_lines {
    let line: Vec<&str> = grammar_line.split("::=").collect();

    if line.len() < 2 {
      continue;
    }

    if grammar_contains(line[0], &sentence_tree) {
      return Err("You have the same nonterminal defined more than once.");
    }

    let grammar_match: Vec<String> = line[1].split("|").map(|s| s.trim().to_string()).collect();
    sentence_tree.insert(line[0].trim().to_string(), grammar_match);
  }

  return Ok(sentence_tree);
}



#[cfg(test)]
mod test {
  use std::collections::HashMap;
  use builder::Solver;
  use builder;

  // TODO: Add more tests

  #[test]
  fn basic() {
    // Build the grammar data structure or err
    let solver = builder::build_solver("n ::= apple | grape | banana | pear | peach");

    let mut temp: Solver = HashMap::new();
    temp.insert("n".to_string(),
                vec!("apple".to_string(),
                     "grape".to_string(),
                     "banana".to_string(),
                     "pear".to_string(),
                     "peach".to_string()));

    assert_eq!(solver, Ok(temp));

    let solver = builder::build_solver("");
    assert_eq!(solver, Err("Your list is empty."));
  }
}
