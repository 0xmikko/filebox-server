/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */
package middlewares

import (
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var signingKey []byte

func InitJWTAuthMiddleware(config *config.Config) {
	signingKey = []byte(config.AuthJWTSecretKey)
	log.Println(signingKey)
}

func JWTAuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		if err := JWTAuth(c); err != nil {
			log.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func JWTAuth(c *gin.Context) error {

	//if c.GetString("projectID") != "" {
	//	return nil
	//}

	token := c.Request.Header.Get("Authorization")
	// Check if toke in correct format
	// ie Bearer: xx03xllasx
	b := "Bearer "
	if !strings.Contains(token, b) {
		return errorhandler.InvalidAuthorisationTokenError()
	}
	t := strings.Split(token, b)
	if len(t) < 2 {
		return errorhandler.InvalidAuthorisationTokenError()
	}
	// Validate token
	valid, err := ValidateToken(t[1])
	if err != nil {
		return errorhandler.InvalidAuthorisationTokenError()
	}

	// set userId Variable
	userID, ok := valid.Claims.(jwt.MapClaims)["user_id"].(string)
	if !ok {
		return errorhandler.InvalidAuthorisationTokenError()
	}

	roles, ok := valid.Claims.(jwt.MapClaims)["roles"].(string)
	if !ok {
		return errorhandler.InvalidAuthorisationTokenError()
	}

	// Set UserID and Role to context
	c.Set("userId", userID)
	c.Set("roles", roles)

	return nil
}

func ValidateToken(t string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	return token, err
}
