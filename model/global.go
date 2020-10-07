package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type DBModel struct {
	Template   interface{}
	session    *mongo.Client
	Collection *mongo.Collection
	Ctx        context.Context
	DBName     string
	ColName    string
}

func (m *DBModel) Init(client *mongo.Client, ctx context.Context) error {
	if len(m.DBName) == 0 || len(m.ColName) == 0 {
		return errors.New("Require valid DB name and collection name.")
	}

	m.session = client
	m.Ctx = ctx
	m.Collection = client.Database(m.DBName).Collection(m.ColName)
	return nil
}
