package app

import (
	api "github.com/go-ewallet/server/api/v1/proto/gen"
	"github.com/go-ewallet/server/internal/config"
	"github.com/go-ewallet/server/internal/domain/server"
	"github.com/go-ewallet/server/internal/domain/ssl"
	"github.com/go-ewallet/server/internal/domain/transaction"
	"github.com/go-ewallet/server/internal/domain/wallet"
	"github.com/go-ewallet/server/pkg/client/couchdb"
	"github.com/go-ewallet/server/pkg/logging"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	cfg    *config.Config
	logger *logging.Logger
}

func NewApp(cfg *config.Config, logger *logging.Logger) (App, error) {
	return App{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (a *App) Run() {
	logger := logging.GetLogger(a.cfg.AppConfig.LogLevel)
	logger.Println("running application...")

	lis, err := net.Listen(a.cfg.Listen.Type, ":"+a.cfg.Listen.Port)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	logger.Println("initializing database...")
	db := &couchdb.CouchDB{}
	db.Init(&a.cfg.Storage, &logger)

	logger.Debug("initializing services...")
	ws := wallet.NewService(db, &logger)
	ts := transaction.NewService(db, &logger, &ws)

	logger.Debug("Load TLS Certificate...")
	tlsCredentials, err := ssl.LoadTLSCredentials()
	if err != nil {
		logger.Fatalf("cannot load TLS credentials: %v", err)
	}

	logger.Info("initializing grpc services...")
	s := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	api.RegisterEWalletServer(s, server.NewEWalletGRPCServer(&ws, &ts))

	logger.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	logger.Println("exiting with code 0")
}
