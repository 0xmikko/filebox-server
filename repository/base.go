/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	Col *mongo.Collection
}

func (basicRepo *BaseRepository) FindOne(result interface{}, id string) error {
	return basicRepo.Col.FindOne(context.Background(), bson.M{"id": id}).Decode(result)
}

func (basicRepo *BaseRepository) FindAll(result interface{}) error {
	cursor, err := basicRepo.Col.Find(context.Background(), bson.M{})
	if err!=nil {
		return err
	}

	return cursor.All(context.Background(), result)
}

func (basicRepo *BaseRepository) Create(item interface{}) error {
	_, err := basicRepo.Col.InsertOne(context.Background(), item)
	return err
}

