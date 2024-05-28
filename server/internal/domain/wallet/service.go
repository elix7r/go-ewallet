package wallet

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/leesper/couchdb-golang"
	couchdb2 "github.com/go-ewallet/server/pkg/client/couchdb"
	"github.com/go-ewallet/server/pkg/logging"
	"github.com/go-ewallet/server/pkg/util"
	"net/url"
)

const dbName = "wallets"

type Service struct {
	db     *couchdb.Database
	logger *logging.Logger
}

func NewService(db *couchdb2.CouchDB, logger *logging.Logger) Service {
	return Service{db: db.GetDatabase(dbName), logger: logger}
}

func (s *Service) GetAll() []*Wallet {
	var res []*Wallet

	query, err := s.db.Query(nil, `_id != ""`, nil, 100, nil, nil)
	if err != nil {
		return nil
	}

	for _, e := range query {
		w := &Wallet{}

		if err := util.Map2Struct(e, w); err != nil {
			s.logger.Fatalf("err: %v", err)
		}

		res = append(res, w)
	}
	return res
}

func (s *Service) GetByUUID(uuid string) (*Wallet, error) {
	query, err := s.db.Get(uuid, nil)
	if err != nil {
		return nil, err
	}

	w := &Wallet{UUID: uuid}

	if err = util.Map2Struct(query, w); err != nil {
		s.logger.Fatalf("err: %v", err)
	}
	return w, nil
}

func (s *Service) Create() (*Wallet, error) {
	w := &Wallet{
		UUID:    uuid.New().String(),
		Balance: 0,
	}

	res, err := util.Struct2Map(w)
	if err != nil {
		return nil, err
	}

	_, _, err = s.db.Save(res, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v. Err: %v", w, err)
	}
	return w, nil
}

func (s *Service) Update(w *Wallet) (*Wallet, error) {
	oldW, _ := s.GetByUUID(w.UUID)

	if oldW != nil {
		err := s.db.Delete(w.UUID)
		if err != nil {
			return nil, fmt.Errorf("failed to delete wallet with uuid: %s. Err: %v", w.UUID, err)
		}
	}

	res, err := util.Struct2Map(w)
	if err != nil {
		return nil, err
	}

	if err = s.db.Set(w.UUID, res); err != nil {
		return nil, fmt.Errorf("failed to update wallet: %v. Err: %v", w, err)
	}
	return w, nil
}

func (s *Service) DeleteByUUID(uuid string) (string, error) {

	err := s.db.Delete(uuid)
	if err != nil {
		return uuid, fmt.Errorf("failed to delete wallet with uuid: %s. Err: %v", uuid, err)
	}
	return uuid, nil
}
