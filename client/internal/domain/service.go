package domain

import (
	"context"
	"fmt"
	api "github.com/go-ewallet/server/api/v1/proto/gen"
)

type CLIService struct {
	ctx context.Context
	c   api.EWalletClient
}

func NewCLIService(ctx context.Context, c api.EWalletClient) CLIService {
	return CLIService{ctx: ctx, c: c}
}

func (c *CLIService) Send(fromWalletUUID string, toWalletUUID string, amount float32) (bool, error) {
	res, err := c.c.Send(c.ctx, &api.SendRequest{
		Trq: &api.TransactionRequest{
			To:     toWalletUUID,
			From:   fromWalletUUID,
			Amount: amount,
		},
	})
	if err != nil {
		// the variable from the structure is not returned here
		//because the generated file for grpc returns nil when an error occurs
		return false, err
	}

	return res.IsSuccessful, nil
}

func (c *CLIService) GetLast() ([]*api.TransactionResponse, error) {
	res, err := c.c.GetLast(c.ctx, &api.GetLastRequest{})
	if err != nil {
		return nil, fmt.Errorf("some error while getting last transcations. Err: %v", err)
	}

	return res.GetT(), nil
}
