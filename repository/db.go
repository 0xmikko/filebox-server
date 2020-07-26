/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"context"
	"github.com/MikaelLazarev/filebox-server/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// Connects to MongoDB using config credentials
func ConnectDB(config *config.Config) *mongo.Database {
	// Set up a timer for connect to 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.DatabaseUrl))
	if err != nil {
		panic(err)
	}

	// Defer DB close
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	// Return database
	return client.Database(config.DatabaseName)
}
