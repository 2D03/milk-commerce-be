package model

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DBModel struct {
	Template   interface{}
	session    *mongo.Client
	Collection *mongo.Collection
	Ctx        context.Context
	DBName     string
	ColName    string
}

type APIResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	ErrorCode string      `json:"errorCode,omitempty"`
	Total     int64       `json:"total,omitempty"`
}

type StatusEnum struct {
	Ok           string
	Error        string
	Invalid      string
	NotFound     string
	Forbidden    string
	Existed      string
	Unauthorized string
}

var APIStatus = &StatusEnum{
	Ok:           "OK",
	Error:        "ERROR",
	Invalid:      "INVALID",
	NotFound:     "NOT_FOUND",
	Forbidden:    "FORBIDDEN",
	Existed:      "EXISTED",
	Unauthorized: "UNAUTHORIZED",
}

func ConvertToBson(item interface{}) (bson.M, error) {
	obj, err := bson.Marshal(item)
	if err != nil {
		return nil, err
	}

	myBson := bson.M{}
	_ = bson.Unmarshal(obj, &myBson)
	return myBson, nil
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

func (m *DBModel) ContextWithTimeout(duration time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	return ctx
}
