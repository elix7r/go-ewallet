package server

import (
	"context"
	api "github.com/go-ewallet/server/api/v1/proto/gen"
	"github.com/go-ewallet/server/internal/domain/transaction"
	"github.com/go-ewallet/server/internal/domain/wallet"
)

type EWalletGRPCServer struct {
	api.UnimplementedEWalletServer
	ws *wallet.Service
	ts *transaction.Service
}

func NewEWalletGRPCServer(ws *wallet.Service, ts *transaction.Service) api.EWalletServer {
	return &EWalletGRPCServer{
		ws: ws,
		ts: ts,
	}
}

func (s *EWalletGRPCServer) Send(_ context.Context, req *api.SendRequest) (*api.SendResponse, error) {
	_, err := s.ts.Create(req.Trq.From, req.Trq.To, float64(req.Trq.Amount), transaction.Type(req.Trq.Type))
	if err != nil {
		return &api.SendResponse{IsSuccessful: false}, err
	}

	return &api.SendResponse{IsSuccessful: true}, nil
}

func (s *EWalletGRPCServer) GetLast(_ context.Context, _ *api.GetLastRequest) (*api.GetLastResponse, error) {
	var res []*api.TransactionResponse

	transactions, err := s.ts.GetByIssued()
	if err != nil {
		return nil, err
	}

	for _, s := range transactions {
		res = append(res, &api.TransactionResponse{
			To:     s.To,
			Amount: float32(s.Amount),
			Date:   s.Date,
		})
	}

	return &api.GetLastResponse{T: res}, nil
}
