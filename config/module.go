/*
 * FileBox server 0.0.1
 * Copyright (c) 2020. Mikhail Lazarev
 */

package config

import (
	"go.uber.org/fx"
)

var Module = fx.Option(
	fx.Provide(NewConfig))
