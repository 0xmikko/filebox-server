/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package controllers

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Invoke(
		RegisterAuthController,
		RegisterBoxController),
	fx.Provide(NewServer))
