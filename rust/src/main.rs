use std::io::prelude::*;
use std::io;
//use std::io::Read;
use std::fs::File;

use std::result::Result;
use builder::Solver;

mod builder;


fn main() {
  match run() {
    // Useful?
    // http://rustbyexample.com/error/result/result_map.html
    Ok(_) => println!("Thanks for playing!"),
    Err(err) => println!("{}", err),
  }
}


fn run<'a>() -> Result<&'a str, &'static str> {
  // Print out "Welcome"
  println!("Welcome to the cse143 random sentence generator (in Rust!).");
  print!("What is the name of the grammar file? ");
  io::stdout().flush().ok().expect("Could not flush stdout"); // https://github.com/rust-lang/rust/issues/23818

  // Get filename from stdin and remove newline
  let mut filename = String::new();
  io::stdin().read_line(&mut filename).expect("failed to read input");
  filename.pop(); // Remove trailing newline TODO: Do this better?
  println!("Trying to open \"{}\"", filename);

  // Open the file and create a file object
  let mut file = match File::open(filename.to_string()) {
    Ok(s) => s,
    Err(_) => return Err("Error opening file."), // TODO: Return error string
  };

  // Read contents to `grammar` string
  let mut grammar = String::new();
  match file.read_to_string(&mut grammar) {
    Ok(s) => s,
    Err(_) => return Err("Error reading file."), // TODO: Return error string
  };

  // Build the grammar data structure or err
  // Useful?
  // http://stackoverflow.com/questions/30801031/read-a-file-and-get-an-array-of-strings
  let solver = match builder::build_solver(&grammar) { // TODO: builder::build_solver
    Ok(s) => s, // std::collections::HashMap<String Vec<String>>
    Err(err) => return Err(&err),
  };

  // Launch main loop to get user's sentence request and generate until no more
  // requests
  return show_results(&solver);
}


fn show_results<'a>(solver: &Solver) -> Result<&'a str, &'static str> {
  loop {
    println!("\nAvailable symbols to generate are:");
    println!("{}", builder::get_symbols(&solver));

    print!("What do you want generated (return to quit)? ");

    let mut target = String::new();
    io::stdin().read_line(&mut target)
      .expect("failed to read line");

    if target.len() == 0 {
      break;
    } else if !builder::grammar_contains(&target, &solver) {
      println!("Illegal symbol");
      continue;
    }

    println!("How many do you want me to generate? ");
    let mut number_text = String::new();
    io::stdin().read_line(&mut number_text)
      .expect("failed to read line");

    let number: i32 = match number_text.trim().parse() {
      Ok(n) => n,
      Err(..) => {
        println!("that's not an integer");
        continue;
      }
    };

    if number < 0 {
      println!("no negatives allowed");
      continue;
    }

    let answers: Vec<&str> = match builder::generate(&target, number, &solver) {
      Ok(vals) => vals,
      Err(err) => return Err(err),
    };

    for answer in answers {
      println!("{}", answer);
    }
  }

  return Ok("success");
}
