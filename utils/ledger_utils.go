package utils

import (
	"errors"

	"github.com/Luismorlan/btc_in_go/model"
)

func CreateUtxoFromInput(input *model.Input) model.UTXO {
	return model.UTXO{
		PrevTxHash: input.PrevTxHash,
		Index:      input.Index,
	}
}

// Handle transaction:
// 1. Validate transaction.
// 2. Claim every input.
// 3. Store every output.
// Return true if currently handles the transactions, false if the transaction is invalid.
// Note: ledger will be changed afterwards, please make a deep copy before passing in.
func HandleTransaction(tx *model.Transaction, l *model.Ledger) error {
	// First validate the transaction.
	err := IsValidTransaction(tx, l)
	if err != nil {
		return err
	}

	ProcessInputsAndOutputs(tx, l)

	return nil
}

func ProcessInputsAndOutputs(tx *model.Transaction, l *model.Ledger) {
	// Claim every input
	for i := 0; i < len(tx.Inputs); i++ {
		input := tx.Inputs[i]
		utxo := CreateUtxoFromInput(input)
		delete(l.L, model.GetUtxoLite(&utxo))
	}

	// Store every output
	for i := 0; i < len(tx.Outputs); i++ {
		output := tx.Outputs[i]
		utxo := model.UTXO{
			PrevTxHash: tx.Hash,
			Index:      int64(i),
		}
		l.L[model.GetUtxoLite(&utxo)] = output
	}
}

// Handle a bunch of transactions.
// Note that ledger will be changed directly, when passing ledger to this function, be sure to pass a deep copy.
// When error, this function returns all transactions that causes error.
// MUTABLE:
// * l
func HandleTransactions(txs []*model.Transaction, l *model.Ledger) ([]*model.Transaction, error) {
	errTxs := []*model.Transaction{}
	for i := 0; i < len(txs); i++ {
		tx := txs[i]
		err := HandleTransaction(tx, l)
		if err != nil {
			errTxs = append(errTxs, tx)
		}
	}
	if len(errTxs) != 0 {
		return errTxs, errors.New("one or many transactions are invalid.")
	}
	return errTxs, nil
}

// Deep (well..not deep enough) copy a ledger.
func GetLedgerDeepCopy(l *model.Ledger) *model.Ledger {
	res := model.NewLedger()
	for k, v := range l.L {
		res.L[k] = v
	}
	return res
}
