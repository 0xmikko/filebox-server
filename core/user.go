/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package core

type (
	User struct {
		BaseModel
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
		Retrieve(id string) (*UserRes, error)
		LoginWithApple(req *AppleCodeReq) (*TokenPair, error)
		RefreshToken(refreshToken string) (*TokenPair, error)
	}
)
