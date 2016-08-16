use std::io::prelude::*;
use std::io;
//use std::io::Read;
use std::fs::File;
use std::error::Error;

use std::result::Result;
use builder::Solver;

mod builder;


// TODO: Comment and clean
// TODO: Don't just use `String`s everywhere. use `&'a str`
// TODO: Make sure i'm using references and stuff correctly


// https://github.com/rust-lang/rust/issues/23818
// Call after every `print!()`
fn flush_stdout() {
  io::stdout().flush().ok().expect("Could not flush stdout");
}


fn main() {
  match run() {
    // Useful?
    // http://rustbyexample.com/error/result/result_map.html
    Ok(_) => println!("Thanks for playing!"),
    Err(err) => println!("error: {}", err),
  }
}


fn run<'a>() -> Result<&'a str, String> {
  // Print out "Welcome"
  println!("Welcome to the cse143 random sentence generator (in Rust!).");
  print!("What is the name of the grammar file? ");
  flush_stdout();

  // Get filename from stdin and remove newline
  let mut filename = String::new();
  io::stdin().read_line(&mut filename).expect("failed to read input");
  filename.pop(); // Remove trailing newline TODO: Do this better? https://github.com/rust-lang/rust/issues/11404
  println!("Trying to open \"{}\"", filename);

  // Open the file and create a file object
  let mut file = match File::open(filename.to_string()) {
    Ok(s) => s,
    Err(e) => return Err(e.description().to_string()),
  };

  // Read contents to `grammar` string
  let mut grammar = String::new();
  match file.read_to_string(&mut grammar) {
    Ok(s) => s,
    Err(e) => return Err(e.description().to_string()),
  };

  // Build the grammar data structure or err
  let solver = match builder::build_solver(&grammar) {
    Ok(s) => s,
    Err(e) => return Err(e.to_string()),
  };

  // Launch main loop to get user's sentence request and generate until no more
  // requests
  return show_results(&solver);
}


fn show_results<'a>(solver: &Solver) -> Result<&'a str, String> {
  loop {
    println!("\nAvailable symbols to generate are:");
    println!("{}", builder::get_symbols(&solver));

    print!("What do you want generated (return to quit)? ");
    flush_stdout();

    let mut target = String::new();
    io::stdin().read_line(&mut target).expect("failed to read line");
    target.pop(); // Remove trailing newline TODO: Do this better? https://github.com/rust-lang/rust/issues/11404

    if target.len() == 0 {
      break;
    } else if !builder::grammar_contains(&target, &solver) {
      println!("Illegal symbol");
      continue;
    }

    print!("How many do you want me to generate? ");
    flush_stdout();

    let mut number_text = String::new();
    io::stdin().read_line(&mut number_text).expect("failed to read line");

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

    let answers: Vec<String> = match builder::generate(&target, number, &solver) {
      Ok(vals) => vals,
      Err(err) => return Err(err.to_string()),
    };

    for answer in answers {
      println!("{}", answer);
    }
  }

  return Ok("success");
}
