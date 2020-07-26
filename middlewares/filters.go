/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package middlewares

import (
	"errors"
	"github.com/MikaelLazarev/filebox-server/errorhandler"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func HasRole(role string) gin.HandlerFunc {

	return func(c *gin.Context) {
		rolesRaw, ok := c.Get("roles")

		if !ok {
			errorhandler.ResponseWithAPIError(c, errorhandler.HttpForbiddenRequestError())
			return
		}

		roles, ok := rolesRaw.(string)
		if !ok {
			errorhandler.ResponseWithAPIError(c, errorhandler.HttpForbiddenRequestError())
			return
		}

		if strings.Contains(roles, role) {
			c.Next()
		}

	}

}

func BlockAPICalls(c *gin.Context) {

	url := c.Request.URL.Path

	if strings.HasPrefix(url, "/api") || strings.HasPrefix(url, "/auth") {
		log.Println("Request was blocked by Block API Middleware cause it comes to static!")
		c.AbortWithStatusJSON(http.StatusForbidden, errorhandler.ForbiddenError(errors.New("request was blocked by Block API Middleware cause it comes to static")))
		return
	}

	c.Next()
}

func GrantAdminForTests(c *gin.Context) {
	c.Set("roles", []string{"admin"})
}
