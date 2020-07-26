/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package middlewares

import (
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

var SSLRedirect string

func InitHttpRedirect(config *config.Config) {
	SSLRedirect = config.SSLRedirect
}

func HTTPSRedirect() gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.Request.Header.Get("X-Forwarded-Proto") == "http" {
			url := c.Request.URL
			url.Scheme = "https"
			url.Host = SSLRedirect

			c.Redirect(http.StatusPermanentRedirect, url.String())
			return
		}

		c.Next()
	}
}
