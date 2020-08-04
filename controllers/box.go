/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"encoding/json"
	"errors"
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/MikaelLazarev/filebox-server/core"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/MikaelLazarev/filebox-server/middlewares"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type BoxController struct {
	service core.BoxServiceI
	tempDir string
}

// BoxController: /api/boxes/
func RegisterBoxController(config *config.Config, g *gin.Engine, ls core.BoxServiceI) {

	controller := BoxController{
		service: ls,
		tempDir: config.TemporaryDir,
	}

	r := g.Group("/api/boxes/", middlewares.JWTAuthMiddleware())
	r.GET("/", controller.ListByCoord)
	r.GET("/i/:id/", withPrincipalAndId(controller.Retrieve))
	r.GET("/d/:id/", withId(controller.Download))
	r.POST("/", withFile(controller.Upload))

}

// GET: /api/boxes/
// Returns array of boxes around user by his/her coordinate
func (bc *BoxController) ListByCoord(c *gin.Context) {

	var req core.BoxListRequest
	if err := c.BindQuery(&req); err != nil {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get lng & lat")))
		return
	}

	log.Println(req)

	result, err := bc.service.FindNearAndTopBoxes(req)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET: /api/boxes/i/:id/
// Return Box info for IPFS hash
func (bc *BoxController) Retrieve(c *gin.Context, userId, id string) {
	log.Println("USER ID:", userId)
	result, err := bc.service.Retrieve(id)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GET: /api/boxes/d/:id/
// Return file for download
func (bc *BoxController) Download(c *gin.Context, id string) {
	tmpFileName, filename, err := bc.service.Download(id)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	defer os.Remove(tmpFileName)
	c.FileAttachment(tmpFileName, filename)
}

// POST: /api/boxes/
// Returns 201 if successfully created
func (bc *BoxController) Upload(c *gin.Context, filename, tmpFilename string) {

	var dto core.BoxCreateRequest

	// Getting box data from request
	boxJson := c.PostForm("box")
	if err := json.Unmarshal([]byte(boxJson), &dto); err != nil {
		errorhandler.ResponseWithAPIError(c, errorhandler.HttpBadRequestError(errors.New("Cant get box")))
		return
	}

	// Creating Box with file contents
	result, err := bc.service.Create(dto, tmpFilename, filename)
	if err != nil {
		errorhandler.ResponseWithAPIError(c, err)
		return
	}

	// Return 201 of succeeded with newBox parameters
	c.JSON(http.StatusCreated, result)
}
