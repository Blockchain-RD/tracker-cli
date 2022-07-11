package types

import "time"

type Transaction struct {
	Id    string
	Date  time.Time
	Coin  string
	Value float64
}

const (
	MESSAGE_TRANSACTION_KEEPER_GET_ALL_TRANSACTIONS string = "transactionKeeperGetAllTransactions"
	MESSAGE_TRANSACTION_KEEPER_GET_TRANSACTION      string = "transactionKeeperGetTransaction"
	MESSAGE_TRANSACTION_KEEPER_ADD_TRANSACTION      string = "transactionKeeperAddTransactions"
	MESSAGE_TRANSACTION_KEEPER_UPDATE_TRANSACTION   string = "transactionKeeperUpdateTransaction"
	MESSAGE_TRANSACTION_KEEPER_REMOVE_TRANSACTION   string = "transactionKeeperRemoveTransactions"
)
