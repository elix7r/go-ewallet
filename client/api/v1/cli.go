package v1

import ewallet "github.com/go-ewallet/server/api/v1/proto/gen"

type CLI interface {
	Send(fromWalletUUID string, toWalletUUID string, amount float32) bool
	GetLast() ([]*ewallet.TransactionResponse, error)
}
