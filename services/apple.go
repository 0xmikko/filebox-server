/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package services

import (
	"context"
	"errors"
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/Timothylock/go-signin-with-apple/apple"
	"log"
)

type AppleLoginService struct {
	ClientID string
	secret   string
}

func NewAppleLoginService(config *config.Config) (*AppleLoginService, error) {

	secret, err := apple.GenerateClientSecret(config.AuthAppleSigningKey,
		config.AuthAppleTeamID,
		config.AuthAppleClientID,
		config.AuthAppleKeyID)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &AppleLoginService{
		ClientID: config.AuthAppleClientID,
		secret:   secret,
	}, nil
}

func (s *AppleLoginService) AuthWithCode(code string) (*core.AppleAuthResponse, error) {
	vReq := apple.AppValidationTokenRequest{
		ClientID:     s.ClientID,
		ClientSecret: s.secret,
		Code:         code,
	}

	var resp apple.ValidationResponse

	client := apple.New()
	// Do the verification
	if err := client.VerifyAppToken(context.Background(), vReq, &resp); err != nil {
		log.Println(err)
		return nil, err
	}
	claim, err := apple.GetClaims(resp.IDToken)
	if err != nil {
		return nil, err
	}

	email, ok := (*claim)["email"].(string)
	if !ok {
		return nil, errors.New("Cant get email")
	}
	emailVerified, ok := (*claim)["email_verified"].(string)
	if !ok {
		return nil, errors.New("Cant get email_verified")
	}
	isPrivateEmail, ok := (*claim)["is_private_email"].(string)
	if !ok {
		return nil, errors.New("Cant get email_is_private")
	}

	return &core.AppleAuthResponse{
		Email:           email,
		EmailIsVerified: emailVerified == "true",
		IsPrivateEmail:  isPrivateEmail == "true",
	}, nil
}
