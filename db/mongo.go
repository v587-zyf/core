package db

import (
	"context"
	"sync"

	"github.com/qiniu/qmgo"
)

var defaultClient *MongoClient
var defalutMutex sync.Mutex

func Init(cfg *Config) error {

	defalutMutex.Lock()
	defer defalutMutex.Unlock()

	cli, err := NewMongoClient(cfg.Uri, cfg.DB)
	if err != nil {
		return err
	}

	// close old client
	if defaultClient != nil {
		defaultClient.client.Close(context.Background())
		defaultClient = nil
	}

	defaultClient = cli

	return nil
}

func Client() *MongoClient {
	return defaultClient
}

func DB(dbName string) *qmgo.Database {
	return defaultClient.client.Database(dbName)
}

func Collection(colName string) *qmgo.Collection {
	return defaultClient.db.Collection(colName)
}
