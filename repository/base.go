/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"context"
	"github.com/MikaelLazarev/filebox-server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type BaseRepository struct {
	Col *mongo.Collection
}

func (basicRepo *BaseRepository) FindOne(result interface{}, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return basicRepo.Col.FindOne(context.Background(), bson.M{"_id": objId}).Decode(result)
}

func (basicRepo *BaseRepository) FindAll(result interface{}) error {
	cursor, err := basicRepo.Col.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	return cursor.All(context.Background(), result)
}

func (basicRepo *BaseRepository) Create(item core.BaseModelI) error {
	result, err := basicRepo.Col.InsertOne(context.Background(), item)
	item.SetID(result.InsertedID.(primitive.ObjectID))
	return err
}

func (basicRepo *BaseRepository) Save(item core.BaseModelI) error {
	id := item.GetID()
	result, error := basicRepo.Col.UpdateOne(context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": item},
	)
	log.Println(result)
	return error
}
