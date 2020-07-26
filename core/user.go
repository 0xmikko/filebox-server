/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

import p "github.com/MikaelLazarev/filebox-server/payload"

type (
	User struct {
		ID    string
		Email string
		Name  string
		Score int
		Roles string
	}

	UsersRepositoryI interface {
		BaseRepositoryI
		FindByEmail(dest *User, email string) error
	}

	UsersServiceI interface {
		Retrieve(id string) (*p.UserRes, error)
		RefreshToken(refreshToken string) (*p.TokenPair, error)
	}
)
