/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package services

import (
	"errors"
	"fmt"
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/devfeel/mapper"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type userService struct {
	repository core.UsersRepositoryI
	jwtSecret  []byte
}

func NewUserService(config *config.Config, repo core.UsersRepositoryI) core.UsersServiceI {
	return &userService{
		repository: repo,
		jwtSecret:  []byte(config.AuthJWTSecretKey),
	}
}

// Return User by email
func (s *userService) Retrieve(email string) (*core.UserRes, error) {
	var user core.User
	if err := s.repository.FindByEmail(&user, email); err != nil {
		return nil, errorhandler.DBError(err, "User not found")
	}

	// Map User to UserDTO
	var userDTO core.UserRes
	mapper.Mapper(user, &userDTO)

	return &userDTO, nil
}

// Generates token pair for particular user
func (s *userService) generateTokenPair(user *core.User) (*core.TokenPair, error) {
	// CreateWithEmail the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["roles"] = user.Roles
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		errorhandler.ReportError(err)
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["user_id"] = user.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	rtString, err := refreshToken.SignedString(s.jwtSecret)

	if err != nil {
		errorhandler.ReportError(err)
		return nil, err
	}

	return &core.TokenPair{
		Access:  tokenString,
		Refresh: rtString,
	}, err
}

func (s *userService) RefreshToken(refreshToken string) (*core.TokenPair, error) {
	// Parse takes the token string and a function for looking up the key.
	// The latter is especially useful if you use multiple keys for your application.
	// The standard is to use 'kid' in the head of the token to identify
	// which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, errorhandler.HttpForbiddenRequestError()
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		user_id, ok := claims["user_id"].(string)
		if !ok {
			return nil, errorhandler.HttpForbiddenRequestError()
		}

		var user core.User

		if err := s.repository.FindOne(&user, user_id); err != nil {
			return nil, err
		}
		return s.generateTokenPair(&user)

	}
	return nil, errorhandler.ForbiddenError(errors.New("Refresh token problem"))
}
