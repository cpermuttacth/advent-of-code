package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner, err := getInputScanner()
	if err != nil {
		die("getting scanner", err)
	}

	fuel := 0
	for scanner.Scan() {
		module, err := strconv.Atoi(scanner.Text())
		if err != nil {
			die("converting to int", err)
		}

		fuel += module/3 - 2
	}
	if err := scanner.Err(); err != nil {
		die("reading input", err)
	}

	fmt.Println(fuel)
}

func getInputScanner() (*bufio.Scanner, error) {
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "input")
	}
	input, err := os.Open(os.Args[1])
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(input), nil
}

func die(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
	os.Exit(1)
}