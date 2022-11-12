package app

import (
	api "github.com/firehead666/infotecs-go-test-task/server/api/v1/proto/gen"
	"github.com/firehead666/infotecs-go-test-task/server/internal/config"
	"github.com/firehead666/infotecs-go-test-task/server/internal/domain/server"
	"github.com/firehead666/infotecs-go-test-task/server/internal/domain/transaction"
	"github.com/firehead666/infotecs-go-test-task/server/internal/domain/wallet"
	"github.com/firehead666/infotecs-go-test-task/server/pkg/client/couchdb"
	"github.com/firehead666/infotecs-go-test-task/server/pkg/logging"
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

	logger.Info("initializing grpc services...")
	s := grpc.NewServer()
	api.RegisterEWalletServer(s, server.NewEWalletGRPCServer(&ws, &ts))

	logger.Println("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	logger.Println("exiting with code 0")
}
