package couchdb

import (
	"github.com/leesper/couchdb-golang"
	"github.com/titor999/infotecs-go-ewallet/server/internal/config"
	"github.com/titor999/infotecs-go-ewallet/server/pkg/logging"
)

type CouchDB struct {
	C      *config.StorageConfig
	Srv    *couchdb.Server
	DB     *couchdb.Database
	Url    string
	Logger *logging.Logger
}

func (db *CouchDB) Init(c *config.StorageConfig, logger *logging.Logger) {
	var err error

	db.C = c
	db.Logger = logger

	db.Url = "http://" + db.C.Username + ":" + db.C.Password + "@" + db.C.Host + ":" + db.C.Port

	db.Srv, err = couchdb.NewServer(db.Url)
	if err != nil {
		db.Logger.Printf("failed connection to db: %v\n", err)
	}
}

func (db *CouchDB) GetDatabase(dbName string) *couchdb.Database {
	if !db.Srv.Contains(dbName) {
		db.Logger.Debugf("the couchdb server: %s db isn't exists. creating a new database... ", dbName)
		dbs, err := db.Srv.Create(dbName)
		if err != nil {
			db.Logger.Fatalf("the cdb Init error: %v", err)
		}
		return dbs
	}

	db.Logger.Debugf("the couchdb server: %s db is exists", dbName)
	dbs, err := couchdb.NewDatabase(db.Url + "/" + dbName)
	if err != nil {
		db.Logger.Fatalf("the couchDB Init error: %v", err)
	}

	return dbs
}
