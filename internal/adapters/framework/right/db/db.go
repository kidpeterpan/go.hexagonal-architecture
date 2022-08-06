package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Adapter struct {
	DBClient *mongo.Client
}

func NewAdapter() *Adapter {
	// connect to db
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// test db connection
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	return &Adapter{
		DBClient: client,
	}
}

func (a Adapter) AddToHistory(answer int32, operation string) error {
	collection := a.DBClient.Database("hexagonal").Collection("arith_history")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	arithHis := struct {
		Answer    int32     `bson:"answer"`
		Operation string    `bson:"operation"`
		CreatedAt time.Time `bson:"created_at"`
	}{
		Answer:    answer,
		Operation: operation,
		CreatedAt: time.Now(),
	}
	_, err := collection.InsertOne(ctx, arithHis)
	if err != nil {
		return err
	}
	return nil
}

func (a Adapter) CloseDbConnection() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.DBClient.Disconnect(ctx); err != nil {
		return err
	}
	return nil
}
