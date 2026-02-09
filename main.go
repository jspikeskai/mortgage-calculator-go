package main

import "C"
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

const WELCOME_MESSAGE = "Welcome to Mortgage Calculator!"
const INFO_MESSAGE = "To exit the program type "
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

func (c MortgageConfig) IsValidData(data float64, min float64, max float64) bool {
	return data >= min && data <= max
}

type MortgageResult struct {
	Principal    float64
	DownPayment  float64
	InterestRate float64
	Period       float64

	MortgageAmount float64
	TotalPayment   float64
	TotalInterest  float64
}

func (m *MortgageResult) CalculateMortgage() {
	m.Principal -= m.DownPayment
	partialFormula := math.Pow(1.0+m.InterestRate, m.Period)
	m.MortgageAmount = m.Principal * (m.InterestRate * partialFormula) / (partialFormula - 1.0)
	m.TotalPayment = m.MortgageAmount * m.Period
	m.TotalInterest = m.TotalPayment - m.Principal
}

var config = MortgageConfig{
	MinPrincipal: 1_000,
	MaxPrincipal: 1_000_000,

	MinDownPayment: 0,
	MaxDownPayment: -1,

	MinInterestRate: 5,
	MaxInterestRate: 30,

	MinPeriod: 5,
	MaxPeriod: 50,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	printer := message.NewPrinter(language.AmericanEnglish)

	for {
		mortgage := MortgageResult{
			Principal:    -1,
			DownPayment:  -1,
			InterestRate: -1,
			Period:       -1,
		}

		fmt.Println(Color(BLUE, WELCOME_MESSAGE))
		fmt.Printf("%s'%s'\n\n", INFO_MESSAGE, Color(RED, "quit"))

		isValid := false
		for !isValid {
			mortgage.Principal = readInput("Principal: ", scanner)
			isValid = config.IsValidData(mortgage.Principal, config.MinPrincipal, config.MaxPrincipal)

			if !isValid {
				printInvalidDataMessage(printer, config.MinPrincipal, config.MaxPrincipal)
			}
		}

		// If Principal is less than MAX_DOWN_OFFSET, we set the MaxDownPayment to 0
		config.MaxDownPayment = math.Max(mortgage.Principal-MAX_DOWN_OFFSET, 0)

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
		mortgage.InterestRate = (mortgage.InterestRate / 100.0) / 12.0

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

		mortgage.CalculateMortgage()

		mortgageMessage := Color(BLUE, "Monthly Mortgage: ")
		mortgageAmount := currency.Symbol(currency.USD.Amount(mortgage.MortgageAmount))

		totalPaymentMessage := Color(BLUE, "Total Payment: ")
		totalPaymentAmount := currency.Symbol(currency.USD.Amount(mortgage.TotalPayment))

		totalInterestMessage := Color(BLUE, "Total Interest: ")
		totalInterestAmount := currency.Symbol(currency.USD.Amount(mortgage.TotalInterest))

		printer.Printf("%s%v\n", mortgageMessage, Color(GREEN, printer.Sprintf("%v", mortgageAmount)))
		printer.Printf("%s%v\n", totalPaymentMessage, Color(GREEN, printer.Sprintf("%v", totalPaymentAmount)))
		printer.Printf("%s%v\n\n", totalInterestMessage, Color(GREEN, printer.Sprintf("%v", totalInterestAmount)))
	}
}

func readInput(prompt string, scanner *bufio.Scanner) float64 {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		input = strings.ReplaceAll(input, "$", "")
		input = strings.ReplaceAll(input, ",", "")
		input = strings.ReplaceAll(input, " ", "")

		if strings.ToLower(input) == "quit" {
			fmt.Println(Color(CYAN, QUIT_MESSAGE))
			os.Exit(0)
		}

		value, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Printf("%sMust enter a number value\n", Color(RED, INVALID_MESSAGE))
			continue
		}

		return value
	}
}

func printInvalidDataMessage(printer *message.Printer, min float64, max float64) {
	printer.Printf("%sNumber must be between %v and %v\n",
		Color(RED, INVALID_MESSAGE),
		number.Decimal(min, number.MaxFractionDigits(2)),
		number.Decimal(max, number.MaxFractionDigits(2)),
	)
}
