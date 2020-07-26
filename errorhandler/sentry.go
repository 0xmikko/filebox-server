/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package errorhandler

import (
	"errors"
	"fmt"
	"github.com/MikaelLazarev/filebox-server/config"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"log"
)

var envIsProd bool

func NewErrorReporter(config *config.Config, router *gin.Engine) {

	if config.Env != "development" {

		if err := sentry.Init(sentry.ClientOptions{
			Dsn: config.SentryDSN,
		}); err != nil {
			log.Fatalf("Sentry initialization failed: %v\n", err)
		}

		router.Use(sentrygin.New(sentrygin.Options{}))
		log.Println("Sentry service was started")
		envIsProd = true
	}
}

func ReportError(err error) {

	pae, ok := err.(*ApiError)
	if ok {
		err = errors.New(fmt.Sprintf("[%s]: %s\nError: %s\n", pae.Module, pae.Message, pae.Error()))
	}

	ae, ok := err.(ApiError)
	if ok {
		err = errors.New(fmt.Sprintf("[%s]: %s\nError: %s\n", ae.Module, ae.Message, ae.Error()))
	}

	if envIsProd {
		sentry.CaptureException(err)
		return
	}
	log.Fatalln(err)
}

func ReportMessage(msg string) {
	if envIsProd {
		sentry.CaptureMessage(msg)
	} else {
		log.Println(msg)
	}
}
