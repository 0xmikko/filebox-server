/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPSRedirect(t *testing.T) {

	SSLRedirect := true

	w := httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	c, r := gin.CreateTestContext(w)

	r.GET("/test", HTTPSRedirect(), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	c.Request, _ = http.NewRequest(http.MethodGet, "/test", nil)
	c.Request.Header.Set("X-Forwarded-Proto", "http")

	r.ServeHTTP(w, c.Request)
	assert.Equal(t, http.StatusPermanentRedirect, w.Code)
	assert.Equal(t, "https://"+SSLRedirect+"/test", w.Header().Get("Location"))
}
