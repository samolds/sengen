all: GrammarMain.cpp
	@g++ -o run *.cpp

debug: GrammarMain.cpp
	@g++ -g *.cpp

clean:
	@rm run

do:
	@make clean
	@make all
	./run
