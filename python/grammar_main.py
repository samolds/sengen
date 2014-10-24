from grammar_solver import GrammarSolver

class GrammarMain:
    def __init__(self):
        print "Welcome to the cse143 random sentence generator.\n"
        filename = raw_input("What is the name of the grammar file? ")
        grammar_file = open(filename.strip(), 'r')
        grammar = []
        for line in grammar_file:
            if (len(line.strip()) > 0):
                grammar.append(line.strip())

        solver = GrammarSolver(grammar)
        self.show_results(solver)

    def show_results(self, solver):
        done = False
        while not done:
          print "\nAvailable symbols to generate are:"
          print solver.get_symbols()
          target = raw_input("What do you want generated (return to quit)? ")
          if len(target) == 0:
              done = True
          elif not solver.grammar_contains(target):
              print "Illegal Symbol"
          else:
              value = raw_input("How many do you want me to generate? ")
              try:
                  number = int(value)
                  if number < 0:
                      print "no negatives allowed"
                  else:
                      answers = solver.generate(target, number)
                      for answer in answers:
                          print answer
              except ValueError:
                  print("that's not an integer")


GrammarMain()
