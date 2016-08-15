extern crate rand;
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


pub fn get_symbols<'a>(sentence_tree: &Solver) -> &'a str {
  return "Hello";
}


pub fn generate<'a>(target: &'a str, number: i32, sentence_tree: &Solver) -> Result<Vec<&'a str>, &'static str> {
  return Ok(vec!(target));
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
  use std::result::Result;
  use builder::Solver;
  use builder;

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
