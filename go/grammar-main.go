package main

import (
  "bufio"
  "fmt"
  "os"
)


func check(e error) {
  if e != nil {
    panic(e)
  }
}


func showResults(solver map[string][]string) {
  done := false
  for !done {
    fmt.Println("\nAvailable symbols to generate are:")
    fmt.Println(getSymbols(solver))
    var target string
    fmt.Print("What do you want generated (return to quit)? ")
    fmt.Scanf("%s", &target)
    if len(target) == 0 {
      done = true;
    } else if !grammarContains(target, solver) {
      fmt.Println("Illegal symbol")
    } else {
      var number int
      fmt.Print("How many do you want me to generate? ")
      n, err := fmt.Scanf("%d", &number)
      if err != nil || n != 1 {
        fmt.Println("that's not an integer")
      } else {
        if number < 0 {
          fmt.Println("no negatives allowed")
        } else {
          answers := generate(target, number, solver)
          for _, answer := range answers {
            fmt.Println(answer)
          }
        }
      }
    }
  }
}


func main() {
  fmt.Println("Welcome to the cse143 random sentence generator.\n")

  var filename string
  fmt.Print("What is the name of the grammar file? ")
  fmt.Scanf("%s", &filename)

  file, err := os.Open(filename)
  check(err)
  var grammar []string
  input := bufio.NewScanner(file)
  for input.Scan() {
    grammar = append(grammar, input.Text())
  }
  file.Close()

  solver := buildSolver(grammar)
  showResults(solver)
}
