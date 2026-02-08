package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const QUIT_KEYWORD = "quit"
const WELCOME_MESSAGE = "Welcome to Mortgage Calculator!"
const INFO_MESSAGE = "To exit the program type '" + QUIT_KEYWORD
const QUIT_MESSAGE = "Thanks for using Mortgage Calculator!"
const INVALID_MESSAGE = "Invalid input: "

// MAX_DOWN_OFFSET MaxDownPayment = (Principal - MAX_DOWN_OFFSET)
const MAX_DOWN_OFFSET = 50_000

type MortgageConfig struct {
	MinPrincipal float64
	MaxPrincipal float64

	MinDownPayment float64
	MaxDownPayment float64

	MinInterestRate float64
	MaxInterestRate float64

	MinPeriod float64
	MaxPeriod float64
}

func (m MortgageConfig) IsValidData(data float64, min float64, max float64) bool {
	return data >= min && data <= max
}

type MortgageResult struct {
	Principal    float64
	DownPayment  float64
	InterestRate float64
	Period       float64
}

var config = MortgageConfig{
	MinPrincipal: 1_000,
	MaxPrincipal: 1_000_000,

	MinDownPayment: 0,
	MaxDownPayment: -1,

	MinInterestRate: 5,
	MaxInterestRate: 30,

	MinPeriod: 5,
	MaxPeriod: 30,
}

func main() {
	isQuitting := false
	scanner := bufio.NewScanner(os.Stdin)
	printer := message.NewPrinter(language.English)

	for !isQuitting {
		mortgage := MortgageResult{
			Principal:    -1,
			DownPayment:  -1,
			InterestRate: -1,
			Period:       -1,
		}

		fmt.Println(WELCOME_MESSAGE)
		fmt.Println(INFO_MESSAGE)

		for !isQuitting {
			isValid := false
			for !isValid {
				mortgage.Principal = readInput("Principal: ", scanner)
				isValid = config.IsValidData(mortgage.Principal, config.MinPrincipal, config.MaxPrincipal)

				if !isValid {
					printInvalidDataMessage(printer, config.MinPrincipal, config.MaxPrincipal)
				}
			}

			// Todo: Clamp this value so it cant go below 0
			config.MaxDownPayment = mortgage.Principal - MAX_DOWN_OFFSET

			isValid = false
			for !isValid {
				mortgage.DownPayment = readInput("Down Payment: ", scanner)
				isValid = config.IsValidData(mortgage.DownPayment, config.MinDownPayment, config.MaxDownPayment)

				if !isValid {
					printInvalidDataMessage(printer, config.MinDownPayment, config.MaxDownPayment)
				}
			}

			isValid = false
			for !isValid {
				mortgage.InterestRate = readInput("Interest Rate: ", scanner)
				isValid = config.IsValidData(mortgage.InterestRate, config.MinInterestRate, config.MaxInterestRate)

				if !isValid {
					printInvalidDataMessage(printer, config.MinInterestRate, config.MaxInterestRate)
				}
			}

			// Convert annual interest rate to monthly
			mortgage.InterestRate /= 100 / 12

			isValid = false
			for !isValid {
				mortgage.Period = readInput("Period: ", scanner)
				isValid = config.IsValidData(mortgage.Period, config.MinPeriod, config.MaxPeriod)

				if !isValid {
					printInvalidDataMessage(printer, config.MinPeriod, config.MaxPeriod)
				}
			}

			// Convert years to months
			mortgage.Period *= 12
		}
	}
}

func readInput(prompt string, scanner *bufio.Scanner) float64 {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		input = strings.ReplaceAll(input, ",", "")

		if strings.ToLower(input) == "quit" {
			fmt.Println(QUIT_MESSAGE)
			os.Exit(0)
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("%sMust enter a number value\n", INVALID_MESSAGE)
			continue
		}

		return value
	}
}

func printInvalidDataMessage(printer *message.Printer, min float64, max float64) {
	printer.Printf("%sNumber must be between %v and %v\n",
		INVALID_MESSAGE,
		number.Decimal(min, number.MaxFractionDigits(2)),
		number.Decimal(max, number.MaxFractionDigits(2)),
	)
}
