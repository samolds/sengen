#include <iostream>
#include <fstream>
#include <string>
#include "GrammarSolver.h"

using namespace std;


#define MAX_ARRAY_SIZE 100


int readFile(string fileName, string * grammars) {
  bool cont = true;
  int grammarSize, i = 0;
  string line;
  string current;
  ifstream in(fileName);

  while (cont) {
    getline(in, line);
    if (!in.fail()) {
      grammars[grammarSize++] = line;
    } else {
      cont = false;
    }
  }

  return grammarSize;
}

void showResults(GrammarSolver & solver) {
  int i, number = 0;
  string input, target;
  bool done = false;
  while (!done) {
    cout << "\nAvailable symbols to generate are:" << endl;
    cout << solver.getSymbols() << endl;
    cout << "What do you want generated ('q' to quit)? ";
    cin >> target;
    if (target == "q") {
      done = true;
    } else if (!solver.grammarContains(target)) {
      cout << "Illegal symbol" << endl;
    } else {
      cout << "How many do you want me to generate? ";
      cin >> input;
      number = atoi(input.c_str());
      if (i <= 0) {
        cout << "Please enter a valid integer greater than 0." << endl;
      } else {
        string * answers = new string[number];
        solver.generate(target, answers, number);
        /*
        for (i = 0; i < number; i++) {
          cout << answers[i] << endl;
        }
        */
        delete [] answers;
      }
    }
  }
}


int main(int argc, char ** argv) {
  int grammarSize = 0;
  string * grammars = new string[MAX_ARRAY_SIZE];
  string gramFile;

  cout << "Welcome to the cse143 randome sentence generator.\n" << endl;
  cout << "What is the name of the grammar file? ";
  cin >> gramFile;

  grammarSize = readFile(gramFile, grammars);
  GrammarSolver grammar = GrammarSolver(grammars, grammarSize);
  delete [] grammars;

  showResults(grammar);

  return 0;
}
