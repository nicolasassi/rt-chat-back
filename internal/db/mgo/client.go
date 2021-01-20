package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mgo struct {
	uri string
}

func NewClient(ctx context.Context, opts ...Option) (*mongo.Client, error) {
	mgo := new(mgo)
	for _, opt := range opts {
		opt(mgo)
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mgo.uri))
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	return client, err
}
