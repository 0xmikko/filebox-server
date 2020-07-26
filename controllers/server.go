/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"context"
	"github.com/MikaelLazarev/filebox-server/config"
	"go.uber.org/fx"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewServer(config *config.Config) *gin.Engine {

	engine := gin.Default()

	if config.Env != "development" {
		gin.SetMode(gin.ReleaseMode)
	}
	// Redirects all request to https
	//engine.Use(middlewares.HTTPSRedirect())

	// CORS setup
	engine.Use(cors.New(cors.Config{
		//AllowAllOrigins:  true,
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:3000",
			"http://localhost:3001",
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return engine
}

func StartServer(lc fx.Lifecycle, config *config.Config, g *gin.Engine) {
	// Starting server
	addr := ":" + config.Port
	if config.Env != "PROD" {
		addr = "localhost" + addr
	}

	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			log.Print("Starting HTTP server...")
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go func() {

				err := g.Run(addr)
				if err != nil {
					log.Fatal(err)
				}
			}()
			return nil
		},
	})
}
