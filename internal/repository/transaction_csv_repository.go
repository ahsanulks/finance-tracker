package repository

import (
	"context"
	"encoding/csv"
	"errors"
	"financetracker/internal/entity"
	"os"
	"slices"
	"strconv"
	"sync"
	"time"
)

const (
	csvTransactionDateFormat = "2006/01/02" //YYYY/MM/DD
	recordChanSize           = 100
	transactionChanSize      = 100
	numWorkers               = 4
)

type TransactionCsvRepository struct {
	filePath string
}

func NewTransactionCsvRepository(filePath string) *TransactionCsvRepository {
	return &TransactionCsvRepository{
		filePath: filePath,
	}
}

/*
CSV template must follow the specified format, and no fields can be empty:
date,amount,content
- date: YYYY/MM/DD
- amount: integer
- content: string
*/
func (tcr *TransactionCsvRepository) FetchByPeriodDesc(
	ctx context.Context,
	transactionPeriod entity.TransactionPeriod,
) ([]*entity.Transaction, error) {
	file, err := os.Open(tcr.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// skip header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Channel to send records for processing
	recordCh := make(chan []string, recordChanSize)
	// Channel to collect transactions
	transactionCh := make(chan *entity.Transaction, transactionChanSize)
	// WaitGroup to track worker completion
	wg := new(sync.WaitGroup)

	// Start worker processing
	for range numWorkers {
		wg.Add(1)
		go tcr.processRecords(wg, recordCh, transactionCh, transactionPeriod)
	}

	go tcr.readCsvRecords(reader, recordCh)

	// Close transaction channel when all workers are done
	go func() {
		wg.Wait()
		close(transactionCh)
	}()

	// Collect transactions
	var transactions []*entity.Transaction
	for trx := range transactionCh {
		transactions = append(transactions, trx)
	}

	if len(transactions) == 0 {
		return nil, errors.New("no transaction found")
	}

	tcr.sortTransactionsDesc(transactions)

	return transactions, nil
}

func (tcr *TransactionCsvRepository) sortTransactionsDesc(transactions []*entity.Transaction) {
	slices.SortFunc(transactions, func(left, right *entity.Transaction) int {
		if left.Date().After(right.Date()) {
			return -1 // Descending order
		} else if left.Date().Before(right.Date()) {
			return 1
		}
		return 0
	})
}

// process individual records and add it to transactionCh
func (tcr *TransactionCsvRepository) processRecords(wg *sync.WaitGroup, recordCh <-chan []string, transactionCh chan<- *entity.Transaction, transactionPeriod entity.TransactionPeriod) {
	defer wg.Done()
	for record := range recordCh {
		trxDate, _ := time.ParseInLocation(csvTransactionDateFormat, record[0], time.Local)
		if transactionPeriod.IsSamePeriod(trxDate) {
			amount, _ := strconv.ParseInt(record[1], 10, 64)
			transactionCh <- entity.NewTransaction(trxDate, amount, record[2])
		}
	}
}

// Read CSV records and send them to recordCh
func (tcr *TransactionCsvRepository) readCsvRecords(reader *csv.Reader, recordCh chan<- []string) {
	defer close(recordCh)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		recordCh <- record
	}
}
