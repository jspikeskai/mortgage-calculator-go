package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const QUIT_KEYWORD = "quit"
const WELCOME_MESSAGE = "Welcome to Mortgage Calculator!"
const INFO_MESSAGE = "\nTo exit the program type '" + QUIT_KEYWORD + "'\n"
const QUIT_MESSAGE = "Thanks for using Mortgage Calculator!"
const INVALID_MESSAGE = "Invalid input: "

type Mortgage struct {
	MinPrincipal int
	MaxPrincipal int

	MinDownPayment int

	MinInterestRate int
	MaxInterestRate int

	MinPeriod int
	MaxPeriod int
}

func main() {
	isQuitting := false
	scanner := bufio.NewScanner(os.Stdin)

	for !isQuitting {
		//config := Mortgage{
		//	MinPrincipal: 1_000,
		//	MaxPrincipal: 1_000_000,
		//
		//	MinDownPayment: 0,
		//
		//	MinInterestRate: 5,
		//	MaxInterestRate: 30,
		//
		//	MinPeriod: 5,
		//	MaxPeriod: 30,
		//}

		fmt.Println(INFO_MESSAGE)

		for !isQuitting {
			principal := readInput("Principal: ", scanner)
			fmt.Println(principal)
		}
	}
}

func readInput(prompt string, scanner *bufio.Scanner) int {
	fmt.Print(prompt)

	for {
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "quit" {
			fmt.Println(QUIT_MESSAGE)
			os.Exit(0)
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Printf("%s, Must enter a number value", INVALID_MESSAGE)
			continue
		}

		return value
	}
}
