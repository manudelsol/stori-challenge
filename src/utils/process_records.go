package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ProcessRecords(records [][]string) (map[string]interface{}, float64, int64, error) {
	totalCredits := 0.0
	totalDebits := 0.0
	creditCount := 0
	debitCount := 0
	monthlyTransactions := make(map[int64]int64)

	for _, record := range records {
		month := strings.Split(record[1], "/")[0]
		transaction := record[2]
		intMonth, err := strconv.ParseInt(month, 10, 64)
		if err != nil {
			return nil, 0, 0, err
		}

		if strings.HasPrefix(transaction, "+") {
			amount, err := strconv.ParseFloat(strings.TrimPrefix(transaction, "+"), 64)
			if err != nil {
				return nil, 0, 0, err
			}
			totalCredits += amount
			creditCount++
		} else if strings.HasPrefix(transaction, "-") {
			amount, err := strconv.ParseFloat(strings.TrimPrefix(transaction, "-"), 64)
			if err != nil {
				return nil, 0, 0, err
			}
			totalDebits += amount
			debitCount++
		}

		monthlyTransactions[intMonth]++
	}

	totalBalance := totalCredits - totalDebits

	averageCredit := 0.0
	averageDebit := 0.0
	if creditCount > 0 {
		averageCredit = totalCredits / float64(creditCount)
	}
	if debitCount > 0 {
		averageDebit = totalDebits / float64(debitCount)
	}

	monthlyTransactionsHtmlString, err := generateHtmlElementsFromTransactions(monthlyTransactions)
	if err != nil {
		return nil, 0, 0, err
	}

	totalBalance = toFixed(totalBalance, 2)
	averageCredit = toFixed(averageCredit, 2)
	averageDebit = toFixed(averageDebit, 2)

	infoMap := map[string]interface{}{
		"total_balance":         totalBalance,
		"average_credit":        averageCredit,
		"average_debit":         averageDebit,
		"transactions_by_month": monthlyTransactionsHtmlString,
	}

	return infoMap, totalBalance, int64(creditCount + debitCount), nil
}

func generateHtmlElementsFromTransactions(monthlyTransactions map[int64]int64) (*string, error) {
	var monthNames = map[int64]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	var htmlContent string
	for month, totalTransactions := range monthlyTransactions {
		monthName, exists := monthNames[month]
		if !exists {
			return nil, fmt.Errorf("month out of range")
		}
		htmlContent += fmt.Sprintf(`<p>Total transactions in %s: %d</p>`, monthName, totalTransactions)
	}

	return &htmlContent, nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
