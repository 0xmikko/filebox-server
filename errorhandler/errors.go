/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package errorhandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiError struct {
	Module  string
	Message string
	Code    int
	Err     error
}

func (ae ApiError) Error() string {
	if ae.Err != nil {
		return ae.Err.Error()
	}
	return "Empty error provided"
}

func ResponseWithAPIError(c *gin.Context, e error) {

	LogError(e)

	code := http.StatusInternalServerError
	message := e.Error()

	if ae, ok := e.(ApiError); ok {
		code = ae.Code
		message = ae.Message
	}

	c.AbortWithStatusJSON(code, gin.H{"message": message})
}

func LogError(err error) {
	ae, ok := err.(ApiError)
	if ok {
		log.Printf("[%s]: %s\nError: %s\n", ae.Module, ae.Message, ae.Error())
		return
	}
	log.Println(err)
}
