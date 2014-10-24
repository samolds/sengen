class GrammarMain
  def initialize
    puts "Welcome to the cse143 random sentence generator.\n"
    puts "What is the name of the grammar file? "
    filename = gets.chomp()
    # grammar = []
    # for line in grammar_file:
    #   if (len(line.strip()) > 0):
    #     grammar.append(line.strip())
    # solver = GrammarSolver(grammar)
    # self.show_results(solver)
  end

  def show_results(solver)
    done = false
    while !done
      puts "\nAvailable symbols to generate are:"
      puts solver.get_symbols()
      puts "What do you want generated (return to quit)? "
      target = gets.chomp()
      if target.length() == 0
        done = true
      else if !solver.grammar_contains(target)
        puts "Illegal Symbol"
      else
        puts "How many do you want me to generate? "
        value = gets.chomp()
        if !is_number?(value)
          puts "that's not an integer"
        else
          number = value.to_int()
          if number < 0
            puts "no negatives allowed"
          else
            answers = solver.generate(target, number)
            for answer in answers
              puts answer
            end
          end
        end
      end
    end
  end
end



GrammarMain()
