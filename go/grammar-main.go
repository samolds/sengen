package main

import (
	"bufio"
	"fmt"
	"os"
)

func showResults(solver map[string][]string) error {
	for true {
		fmt.Println("\nAvailable symbols to generate are:")
		fmt.Println(getSymbols(solver))
		var target string
		fmt.Print("What do you want generated (return to quit)? ")
		_, err := fmt.Fscanf(os.Stdin, "%s", &target)
		if err != nil {
			return err
		}

		if len(target) == 0 {
			break
		} else if !grammarContains(target, solver) {
			fmt.Fprintln(os.Stderr, "Illegal symbol")
			continue
		}

		var number int
		fmt.Print("How many do you want me to generate? ")
		n, err := fmt.Fscanf(os.Stdin, "%d", &number)
		if err != nil || n != 1 {
			fmt.Fprintln(os.Stderr, "that's not an integer")
			continue
		}

		if number < 0 {
			fmt.Fprintln(os.Stderr, "no negatives allowed")
			continue
		}

		answers, err := generate(target, number, solver)
		if err != nil {
			return err
		}

		for _, answer := range answers {
			fmt.Println(answer)
		}
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Println("Welcome to the cse143 random sentence generator.\n")

	var filename string
	fmt.Print("What is the name of the grammar file? ")
	_, err := fmt.Scanf("%s", &filename)
	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var grammar []string
	input := bufio.NewScanner(file)

	for input.Scan() {
		grammar = append(grammar, input.Text())
	}

	solver, err := buildSolver(grammar)
	if err != nil {
		return err
	}

	showResults(solver)

	return nil
}
