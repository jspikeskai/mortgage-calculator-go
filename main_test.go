package main

import (
	"math"
	"testing"
)

func TestCalculateMortgage(t *testing.T) {
	tests := []struct {
		name            string
		principal       float64
		downPayment     float64
		interestRate    float64
		period          float64
		expectedMonthly float64
	}{
		{
			name:            "Standard",
			principal:       1_000_000,
			downPayment:     200_000,
			interestRate:    6.128,
			period:          30,
			expectedMonthly: 4_862.44,
		},

		{
			name:            "Low Interest",
			principal:       300_000,
			downPayment:     75_000,
			interestRate:    5.322,
			period:          15,
			expectedMonthly: 1_817.25,
		},

		{
			name:            "Trump 50 Year",
			principal:       625_000,
			downPayment:     0,
			interestRate:    8.232,
			period:          50,
			expectedMonthly: 4_359.61,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := MortgageResult{
				Principal:    test.principal,
				DownPayment:  test.downPayment,
				InterestRate: test.interestRate,
				Period:       test.period,
			}

			m.InterestRate = (m.InterestRate / 100.0) / 12.0
			m.Period *= 12

			m.CalculateMortgage()

			if !withinTolerance(m.MortgageAmount, test.expectedMonthly, 0.01) {
				t.Errorf("Expected monthly payment %.2f, but got %.2f", test.expectedMonthly, m.MortgageAmount)
			}
		})
	}
}

func withinTolerance(x, y, tolerance float64) bool {
	return math.Abs(x-y) <= tolerance
}
