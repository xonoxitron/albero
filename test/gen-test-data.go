package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Defines a basic wallet to wallet tx structure
type tx struct {
	from string
	to   string
	eth  string
}

// Generates a random but valid ETH address
func generateAddress() string {
	const charset = "abcdef0123456789"
	result := "0x"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 40; i++ {
		result += string(charset[seededRand.Intn(len(charset))])
	}
	return result
}

// Generates mock transactions
func generateTransactions(numRecords int) []tx {
	transactions := make([]tx, numRecords)
	for i := 0; i < numRecords; i++ {
		from := generateAddress()
		to := generateAddress()
		eth := fmt.Sprintf("%.12f", rand.Float64())
		transactions[i] = tx{from, to, eth}
	}
	return transactions
}

// Writes the output to file
func writeTransactionsToFile(transactions []tx, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, t := range transactions {
		line := fmt.Sprintf("From: %s, To: %s, ETH: %s\n", t.from, t.to, t.eth)
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	input := os.Args[1]
	// Checks the input argument
	numRecords, err := strconv.Atoi(input)
	if err != nil {
		// Default in case of invalid input argument provided
		numRecords = 100000
	}
	fmt.Printf("Generating %d random transaction records...\n", numRecords)
	transactions := generateTransactions(numRecords)
	filename := fmt.Sprintf("transactions_%d.txt", numRecords)
	if writeTransactionsToFile(transactions, filename) != nil {
		fmt.Printf("Error writing transactions to file: %v\n", err)
	} else {
		fmt.Printf("Generated %d random transaction records and wrote them to %s\n", numRecords, filename)
	}

}
