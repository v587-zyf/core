package db

import (
	"context"
	"github.com/qiniu/qmgo"
)

type MongoClient struct {
	client *qmgo.Client

	db *qmgo.Database
}

type MongoOption func(cli *MongoClient)

func NewMongoClient(uri string, dbName string) (*MongoClient, error) {

	qc, err := qmgo.NewClient(context.Background(), &qmgo.Config{Uri: uri})
	if err != nil {
		return nil, err
	}

	var db *qmgo.Database
	if dbName != "" {
		db = qc.Database(dbName)
	}

	client := &MongoClient{
		client: qc,
		db:     db,
	}

	return client, nil
}

func (c *MongoClient) GetClient() *qmgo.Client {
	return c.client
}

func (c *MongoClient) GetDB() *qmgo.Database {
	return c.db
}

func (c *MongoClient) DB(dbName string) *qmgo.Database {
	return c.client.Database(dbName)
}

func (c *MongoClient) Collection(colName string) *qmgo.Collection {
	return c.db.Collection(colName)
}
