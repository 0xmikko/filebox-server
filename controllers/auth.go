/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct {
	service core.UsersServiceI
}

func RegisterAuthController(g *gin.Engine, is core.UsersServiceI) {

	controller := authController{
		service: is,
	}
	r := g.Group("/auth/")
	//r.GET("/google/login/", controller.GoogleRedirect)
	r.POST("/login/apple/done/", controller.AppleLoginDone)
	r.POST("/token/refresh/", controller.RefreshToken)

}

// POST: /auth/token/refresh
func (u *authController) RefreshToken(c *gin.Context) {

	var tokenReq core.RefreshTokenReq

	if err := c.BindJSON(&tokenReq); err != nil {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(err))
		return
	}

	tokenPair, err := u.service.RefreshToken(tokenReq.Token)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}

func (u *authController) AppleLoginDone(c *gin.Context) {

	var appleCodeReq core.AppleCodeReq

	if err := c.BindJSON(&appleCodeReq); err != nil {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(err))
		return
	}

	tokenPair, err := u.service.LoginWithApple(&appleCodeReq)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, tokenPair)
}
