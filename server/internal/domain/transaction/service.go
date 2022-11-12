package transaction

import (
	"fmt"
	"github.com/firehead666/infotecs-go-test-task/server/internal/domain/wallet"
	couchdb2 "github.com/firehead666/infotecs-go-test-task/server/pkg/client/couchdb"
	"github.com/firehead666/infotecs-go-test-task/server/pkg/logging"
	"github.com/firehead666/infotecs-go-test-task/server/pkg/util"
	"github.com/google/uuid"
	"github.com/leesper/couchdb-golang"
	"time"
)

const dbName = "transactions"

type Service struct {
	ws     *wallet.Service
	db     *couchdb.Database
	logger *logging.Logger
}

func NewService(db *couchdb2.CouchDB, logger *logging.Logger, ws *wallet.Service) Service {
	return Service{
		db:     db.GetDatabase(dbName),
		logger: logger,
		ws:     ws,
	}
}

func (s *Service) GetAll() ([]*Transaction, error) {
	var res []*Transaction

	query, err := s.db.Query(nil, `_id != ""`, nil, 100, nil, nil)
	if err != nil {
		return nil, err
	}

	for _, e := range query {
		t := &Transaction{}

		if err := util.Map2Struct(e, t); err != nil {
			s.logger.Fatalf("err: %v", err)
		}

		res = append(res, t)
	}

	return res, nil
}

func (s *Service) GetByIssued() ([]*Transaction, error) {
	var res []*Transaction

	query, err := s.db.Query(nil, `is_issued == false`, nil, 100, nil, nil)
	if err != nil {
		return nil, err
	}

	for _, e := range query {
		t := &Transaction{}

		if err := util.Map2Struct(e, t); err != nil {
			s.logger.Fatalf("err: %v", err)
		}

		res = append(res, t)
	}

	for _, el := range res {
		if err := s.UpdateIsIssuedByUUID(el.UUID); err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (s *Service) GetByUUID(uuid string) (*Transaction, error) {
	query, err := s.db.Get(uuid, nil)
	if err != nil {
		return nil, err
	}

	t := &Transaction{UUID: uuid}

	if err = util.Map2Struct(query, t); err != nil {
		s.logger.Fatalf("err: %v", err)
	}
	return t, nil
}

func (s *Service) Create(from string, to string, amount float64, tType Type) (*Transaction, error) {
	t := &Transaction{
		UUID:     uuid.New().String(),
		From:     from,
		To:       to,
		Type:     tType,
		Amount:   amount,
		Date:     time.Now().Format(time.UnixDate),
		IsIssued: false,
	}

	fromW, toW, err := s.execute(t)
	if err != nil {
		return nil, err
	}

	res, err := util.Struct2Map(t)
	if err != nil {
		return nil, err
	}

	_, err = s.ws.Update(fromW)
	if err != nil {
		return nil, err
	}

	_, err = s.ws.Update(toW)
	if err != nil {
		return nil, err
	}

	_, _, err = s.db.Save(res, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v. Err: %v", t, err)
	}

	return t, nil
}

func (s *Service) UpdateIsIssuedByUUID(uuid string) error {
	query, err := s.db.Get(uuid, nil)
	if err != nil {
		return err
	}

	if query != nil {
		if err := s.db.Delete(uuid); err != nil {
			return fmt.Errorf("failed to delete transaction with uuid: %s. Err: %v", uuid, err)
		}
	}

	t := &Transaction{}

	if err = util.Map2Struct(query, t); err != nil {
		s.logger.Fatalf("err: %v", err)
	}

	t.IsIssued = true

	res, err := util.Struct2Map(t)
	if err != nil {
		return err
	}
	if err = s.db.Set(uuid, res); err != nil {
		return fmt.Errorf("failed to update is_issued transaction field: %v. Err: %v", t, err)
	}

	return nil
}

func (s *Service) DeleteByUUID(uuid string) (string, error) {
	err := s.db.Delete(uuid)
	if err != nil {
		return uuid, fmt.Errorf("failed to delete transaction with uuid: %s. Err: %v", uuid, err)
	}

	return uuid, nil
}

func (s *Service) execute(t *Transaction) (*wallet.Wallet, *wallet.Wallet, error) {
	switch t.Type {
	case Transfer:
		tm := &TransferManagerCommand{Transaction: t}
		fromW, toW, err := tm.execute(s.ws)
		if err != nil {
			return nil, nil, err
		}
		return fromW, toW, nil
	}

	return nil, nil, nil
}
