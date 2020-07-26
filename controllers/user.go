/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"errors"
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userController struct {
	service core.UsersServiceI
}

func RegisterUserController(g *gin.Engine, is core.UsersServiceI) {

	controller := userController{
		service: is,
	}
	g.GET("/user/", controller.Retrieve)

}

func (u *userController) GoolgeRedirect(c *gin.Context) {

}

func (t *userController) Retrieve(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get ID")))
		return
	}

	result, err := t.service.Retrieve(id)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
