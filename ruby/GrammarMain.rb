require_relative 'GrammarSolver'


class String
  def is_i?
    /\A[-+]?\d+\z/ === self
  end
end


class GrammarMain
  def initialize
    puts "Welcome to the cse143 random sentence generator.\n"
    print "What is the name of the grammar file? "
    filename = gets.chomp()
    grammar = Array.new
    grammar_file = File.new(filename, "r")
    if !grammar_file
      puts "Unable to open file!"
    else
      IO.foreach(filename){ |line|
        if line.strip().length() > 0
          grammar.push(line.strip())
        end
      }
    end
    solver = GrammarSolver.new(grammar)
    show_results(solver)
  end

  def show_results(solver)
    done = false
    while !done do
      puts "\nAvailable symbols to generate are:"
      puts solver.get_symbols()
      print "What do you want generated (return to quit)? "
      target = gets.chomp()
      if target.length() == 0
        done = true
      elsif !solver.grammar_contains(target)
        puts "Illegal Symbol"
      else
        print "How many do you want me to generate? "
        value = gets.chomp()
        if !value.is_i?
          puts "that's not an integer"
        else
          number = value.to_i
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


GrammarMain.new
