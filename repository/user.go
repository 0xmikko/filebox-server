/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package repository

import (
	"context"
	"github.com/MikaelLazarev/filebox-server/core"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	BaseRepository
}

func NewUsersRepository(db *mongo.Database) core.UsersRepositoryI {

	return &userRepository{
		BaseRepository{
			Col: db.Collection("users"),
		},
	}
}

func (repo *userRepository) FindByEmail(user *core.User ,email string) error {
	return repo.Col.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)

}
