package app

import (
	"context"
	"encoding/json"
	"flag"
	"github.com/firehead666/infotecs-go-test-task/client/internal/config"
	"github.com/firehead666/infotecs-go-test-task/client/internal/domain"
	api "github.com/firehead666/infotecs-go-test-task/server/api/v1/proto/gen"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

type App struct {
	cfg *config.Config
}

func NewApp(cfg *config.Config) (App, error) {
	return App{
		cfg: cfg,
	}, nil
}

func (a *App) Run() {
	sendFlag := flag.Bool("send", false, "use send method in CLI interface")
	getLastFlag := flag.Bool("getlast", false, "use getlast method in CLI interface")
	flag.Parse()

	conn, err := grpc.Dial(":"+a.cfg.GRPC.ServerPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err: %v", err)
		return
	}

	c := api.NewEWalletClient(conn)

	cliService := domain.NewCLIService(context.Background(), c)

	if *sendFlag {
		if flag.NArg() != 3 {
			log.Fatalf("with the `-send` flag you need to provide three args: " +
				"\n\t(1) from wallet UUID (string)" +
				"\n\t(2) to wallet UUID (string)" +
				"\n\t(3) amount float value." +
				"\nEx. go run cmd/client/main.go -send ad3941b2-1382-11ed-861d-0242ac120002 b1ce72c4-1382-11ed-861d-0242ac120002 100")
			return
		}

		amount, err := strconv.ParseFloat(flag.Arg(2), 32)
		if err != nil {
			log.Fatalf("can't convert inputed amount(`%s`) to float value", flag.Arg(2))
			return
		}
		send, err := cliService.Send(flag.Arg(0), flag.Arg(1), float32(amount))
		if err != nil {
			log.Fatalf("Err: %v", err)
			return
		}
		log.Printf("Transaction successful: %t", send)
		return
	}

	if *getLastFlag {
		if flag.NArg() != 0 {
			log.Fatalf("with the `-getlast` flag you don't need to provide any args" +
				"\nEx. go run cmd/client/main.go -getlast")
			return
		}

		res, err := cliService.GetLast()
		if err != nil {
			log.Fatal(err)
		}
		if res == nil {
			log.Println("not issued transactions no found")
			return
		}

		prettyJson, err := json.MarshalIndent(res, "", "    ")
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println(string(prettyJson))
	}
}
