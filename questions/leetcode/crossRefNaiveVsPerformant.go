package questions

import "fmt"

type Transaction struct {
	ID        string
	AccountID string
	Amount    float64
}

// Lookup

// Naive Approach
func FindFlaggedTransactions(transactions []Transaction, flaggedAccounts []string) []Transaction {

	transactionMap := make(map[string]Transaction)

	for _, transaction := range transactions {
		transactionMap[transaction.AccountID] = transaction
	}

	var flaggedTransactions []Transaction

	for _, transaction := range transactions {
		for _, accountId := range flaggedAccounts {
			if transaction.AccountID == accountId {
				flaggedTransactions = append(flaggedTransactions, transaction)
			}
		}
	}

	fmt.Printf("results: %v\n", flaggedTransactions)
	return nil
}

func TestFlagged() {
	transactions := []Transaction{
		{"TXN001", "ACC123", 250.50},
		{"TXN002", "ACC456", 89.99},
		{"TXN003", "ACC789", 1500.00},
		{"TXN004", "ACC123", 45.25},
		{"TXN005", "ACC999", 325.75},
		{"TXN006", "ACC456", 12.50},
		{"TXN007", "ACC888", 750.00},
	}

	// Flagged accounts (suspicious activity detected)
	flaggedAccounts := []string{"ACC123", "ACC888", "ACC999"}

	FindFlaggedTransactions(transactions, flaggedAccounts)
}

/*
*
*
*

8
*
*
8
8
88
8
*/
func FindFlaggedTransactionsReal(transactions []Transaction, flaggedAccounts []string) []Transaction {
	// look up for flagged acocunt
	unsafeTransactions := make(map[string]bool)

	// create unsafe transactions map
	for _, aId := range flaggedAccounts {
		unsafeTransactions[aId] = true
	}

	var flaggedTransactions []Transaction
	for _, transaction := range transactions {
		_, exists := unsafeTransactions[transaction.AccountID]

		// store it if it exists
		if exists {
			flaggedTransactions = append(flaggedTransactions, transaction)
		}
	}

	fmt.Printf("results: %v\n", flaggedTransactions)
	return nil
}
