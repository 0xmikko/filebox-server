/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/MikaelLazarev/filebox-server/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BoxController struct {
	service core.BoxServiceI
}

func RegisterBoxController(g *gin.Engine, ls core.BoxServiceI) {

	controller := BoxController{
		service: ls,
	}

	r := g.Group("/api/boxes/", middlewares.JWTAuthMiddleware())
	r.GET("/", withId(controller.ListByCoord))
	r.POST("/:id/", controller.Upload)

}

// GET: /api/Boxs/:id/
func (lc *BoxController) ListByCoord(c *gin.Context, id string) {

	result, err := lc.service.Retrieve(id)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}


// GET: /api/Boxs/:id/
func (lc *BoxController) Upload(c *gin.Context) {

	result, err := lc.service.Create()
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
