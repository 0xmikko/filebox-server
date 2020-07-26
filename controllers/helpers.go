/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"errors"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/gin-gonic/gin"
)

func withId(handler func(c *gin.Context, id string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ID, ok := c.Params.Get("id")
		if !ok {
			errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get ID")))
			return
		}

		handler(c, ID)
	}
}

func withPrincipal(handler func(c *gin.Context, userId string)) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, ok := c.MustGet("userId").(string)
		if !ok {
			errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get ID")))
			return
		}

		handler(c, userID)
	}
}

func withPrincipalAndId(handler func(c *gin.Context, userId string, id string)) gin.HandlerFunc {

	return withPrincipal(func(c *gin.Context, userID string) {

		id, ok := c.Params.Get("id")
		if !ok {
			errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get ID")))
			return
		}

		handler(c, userID, id)
	})
}
