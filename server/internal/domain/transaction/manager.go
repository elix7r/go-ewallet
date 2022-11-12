package transaction

import (
	"errors"
	"fmt"
	"github.com/firehead666/infotecs-go-test-task/server/internal/domain/wallet"
)

type ManagerCommand interface {
	execute() error
}

type TransferManagerCommand struct {
	Transaction *Transaction
}

func (t *TransferManagerCommand) execute(ws *wallet.Service) (*wallet.Wallet, *wallet.Wallet, error) {
	fromW, err := ws.GetByUUID(t.Transaction.From)
	if err != nil {
		return nil, nil, fmt.Errorf("can't find from fromWallet with uuid: %s", t.Transaction.From)
	}

	toW, err := ws.GetByUUID(t.Transaction.To)
	if err != nil {
		return nil, nil, fmt.Errorf("can't find toWallet with uuid: %s", t.Transaction.To)
	}

	if fromW.Balance < t.Transaction.Amount {
		return nil, nil, errors.New("can't transfer to another wallet. The sender's wallet does not have that much money on the balance")
	}

	fromW.Balance -= t.Transaction.Amount
	toW.Balance += t.Transaction.Amount

	return fromW, toW, nil
}
